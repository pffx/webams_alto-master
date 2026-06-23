package ws_command

import (
	logger "alto_server/common/log"
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

var (
	addr    = flag.String("addr", "127.0.0.1:5500", "http service address")
	cmdPath string
)

func pumpStdin(ws *websocket.Conn, w io.Writer) {
	defer ws.Close()
	var command []byte
	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := ws.ReadMessage()
		fmt.Printf("pumpStdin  message:  %+v\r\n", message)
		if err != nil {
			break
		}
		command = append(command, message...)
		fmt.Printf("pumpStdin  command:  %+v\r\n", command)
		if message[len(message)-1] == 13 {
			fmt.Printf("pumpStdin  Write:  \r\n")
			command = append(command, '\n')
			if _, err := w.Write(command); err != nil {
				fmt.Printf("pumpStdin Write   err:  %+v\r\n", err)
				break
			} else {
				command = []byte("")
			}
		}

	}
}

// func pumpStdin(ws *websocket.Conn, w io.Writer) {
// 	defer ws.Close()
// 	ws.SetReadLimit(maxMessageSize)
// 	ws.SetReadDeadline(time.Now().Add(pongWait))
// 	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
// 	for {
// 		_, message, err := ws.ReadMessage()
// 		if err != nil {
// 			break
// 		}
// 		message = append(message, '\n')
// 		if _, err := w.Write(message); err != nil {
// 			break
// 		}
// 	}
// }

func pumpStdout(ws *websocket.Conn, r io.Reader, done chan struct{}) {
	defer func() {
	}()
	s := bufio.NewScanner(r)
	for s.Scan() {
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		fmt.Printf("pumpStdout  websocket.TextMessage:  %+v\r\n", websocket.TextMessage)
		if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil {
			ws.Close()
			break
		}
	}
	if s.Err() != nil {
		logger.SystemLogger.Debug("scan:", s.Err())
	}
	close(done)

	ws.SetWriteDeadline(time.Now().Add(writeWait))
	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(closeGracePeriod)
	ws.Close()
}

func ping(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				logger.SystemLogger.Debug("ping:", err)
			}
		case <-done:
			return
		}
	}
}

func internalError(ws *websocket.Conn, msg string, err error) {
	logger.SystemLogger.Debug(msg, err)
	fmt.Printf("Web Socket  internalError   error:  %+v\r\n", err)
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}

// 过滤url，可以对请求来源做一些处理
func checkOrigin(r *http.Request) bool {
	return true
}

// 初始化一个websockt升级程序
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func ServeWs(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.SystemLogger.Debug("upgrade:", err)
		return
	}

	defer ws.Close()
	// cmdPath, err := exec.LookPath("powershell.exe")
	cmdPath, err := exec.LookPath("cmd.exe")
	if err != nil {
		fmt.Printf(" ServeWs   powershell  err:  %+v\r\n", err)
		cmdPath, err = exec.LookPath("bash")
		if err != nil {
			fmt.Printf(" ServeWs   bash  err:  %+v\r\n", err)
			return
		}
	}
	fmt.Printf(" ServeWs     cmdPath:  %+v\r\n", cmdPath)

	outr, outw, err := os.Pipe()
	if err != nil {
		internalError(ws, "stdout:", err)
		return
	}
	defer outr.Close()
	defer outw.Close()

	inr, inw, err := os.Pipe()
	if err != nil {
		internalError(ws, "stdin:", err)
		return
	}
	defer inr.Close()
	defer inw.Close()

	proc, err := os.StartProcess(cmdPath, flag.Args(), &os.ProcAttr{
		Files: []*os.File{inr, outw, outw},
	})

	if err != nil {
		fmt.Printf("Start cmd process   failed! error:  %+v\r\n", err)
		internalError(ws, "start:", err)
		return
	}
	logger.SystemLogger.Debug("Start cmd process   success! ")
	fmt.Printf("Start cmd process   success!: \r\n")

	inr.Close()
	outw.Close()

	stdoutDone := make(chan struct{})
	go pumpStdout(ws, outr, stdoutDone)
	go ping(ws, stdoutDone)

	pumpStdin(ws, inw)

	// Some commands will exit when stdin is closed.
	inw.Close()

	// Other commands need a bonk on the head.
	if err := proc.Signal(os.Interrupt); err != nil {
		logger.SystemLogger.Debug("inter:", err)
	}

	select {
	case <-stdoutDone:
	case <-time.After(time.Second):
		// A bigger bonk on the head.
		if err := proc.Signal(os.Kill); err != nil {
			logger.SystemLogger.Debug("term:", err)
		}
		<-stdoutDone
	}

	if _, err := proc.Wait(); err != nil {
		logger.SystemLogger.Debug("wait:", err)
	}
}
