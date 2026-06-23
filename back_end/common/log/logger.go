package logger

import (
	"alto_server/conf"
	"fmt"
	"io"
	"log"

	// "log/syslog"
	"github.com/RackSec/srslog"

	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// 定义全局日志记录器变量
var (
	ConsoleLogger *logrus.Logger
	ErrorLogger   *logrus.Logger
	SystemLogger  *logrus.Logger
	OssLogger     *logrus.Logger
	OntLogger     *logrus.Logger
	OltLogger     *logrus.Logger
	AaaLogger     *logrus.Logger
	ChainLogger   *logrus.Logger
	AlarmLogger   *logrus.Logger

	NorthboundLogger *logrus.Logger
)

// 日志文件配置常量
const (
	MaxFileSize = 50        // 单个日志文件的最大大小（MB）
	MaxBackups  = 5         // 最大备份文件数量
	MaxAge      = 7         // 日志文件保留的最大天数
	logPrefix   = "./logs/" // 日志文件目录前缀
)

type SyslogRemoteLevel int

const (
	SYSLOG_REMOTE_ALERT    SyslogRemoteLevel = iota + 1 // 日志级别：ALERT
	SYSLOG_REMOTE_CRITICAL                              // 日志级别：CRITICAL
	SYSLOT_REMOTE_ERROR                                 // 日志级别：ERROR
	SYSLOG_REMOTE_INFO                                  // 日志级别：INFO
)

func (s SyslogRemoteLevel) String() string {
	switch s {
	case SYSLOG_REMOTE_ALERT:
		return "[ALERT]"
	case SYSLOG_REMOTE_CRITICAL:
		return "[CRITICAL]"
	case SYSLOT_REMOTE_ERROR:
		return "[ERROR]"
	case SYSLOG_REMOTE_INFO:
		return "[INFO]"
	default:
		return "[UNKNOWN]"
	}
}

// 日志文件路径映射
var logFiles = map[string]string{
	"system":     logPrefix + "syslog.log",
	"error":      logPrefix + "error.log",
	"oss":        logPrefix + "oss.log",
	"ont":        logPrefix + "ont.log",
	"olt":        logPrefix + "olt.log",
	"aaa":        logPrefix + "aaa.log",
	"chain":      logPrefix + "chain.log",
	"alarm":      logPrefix + "alarm.log",
	"northbound": logPrefix + "nbi.log",
}

// 初始化所有日志记录器
func InitLogger() {
	createLogDir() // 创建日志目录
	//ConsoleLogger = initConsoleLogger(logrus.DebugLevel)
	ErrorLogger = initFileLogger(logFiles["error"], logrus.DebugLevel)
	SystemLogger = initFileLogger(logFiles["system"], logrus.DebugLevel)
	NorthboundLogger = initFileLogger(logFiles["nbi"], logrus.DebugLevel)
	OssLogger = initFileLogger(logFiles["oss"], logrus.DebugLevel)
	OntLogger = initFileLogger(logFiles["ont"], logrus.DebugLevel)
	OltLogger = initFileLogger(logFiles["olt"], logrus.DebugLevel)
	AaaLogger = initFileLogger(logFiles["aaa"], logrus.DebugLevel)
	ChainLogger = initFileLogger(logFiles["chain"], logrus.DebugLevel)
	AlarmLogger = initFileLogger(logFiles["alarm"], logrus.DebugLevel)
}

// 创建日志目录
func createLogDir() {
	if err := os.MkdirAll(logPrefix, 0777); err != nil {
		fmt.Printf("Create log files failed: %v\n", err)
	}
}

// 添加日志前缀常量
const (
	SystemPrefix     = "[SYSTEM]"
	ErrorPrefix      = "[ERROR]"
	OssPrefix        = "[OSS]"
	OntPrefix        = "[ONT]"
	OltPrefix        = "[OLT]"
	AaaPrefix        = "[AAA]"
	ChainPrefix      = "[CHAIN]"
	AlarmPrefix      = "[ALARM]"
	NorthboundPrefix = "[NBI]"
)

type CustomFormatter struct {
	Prefix string // 添加前缀字段
}

// 自定义日志格式
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000000")
	level := strings.ToUpper(entry.Level.String())
	file := ""
	if entry.HasCaller() {
		// 获取当前工作目录
		projectRoot, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		relativePath, err := filepath.Rel(projectRoot, entry.Caller.File)
		if err != nil {
			relativePath = entry.Caller.File
		}
		file = fmt.Sprintf("%s:%d", relativePath, entry.Caller.Line)
	}
	// 添加前缀到消息中
	prefix := f.Prefix
	if prefix == "" {
		prefix = "[CONSOLE]"
	}
	message := fmt.Sprintf("%s %s %s %s - %s\n", timestamp, prefix, level, file, entry.Message)
	return []byte(message), nil
}

// 初始化文件日志记录器
func initFileLogger(fileName string, level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	// 设置日志输出到文件，使用lumberjack进行日志轮转
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fileName,    // 日志文件路径
		MaxSize:    MaxFileSize, // 单文件最大容量
		MaxBackups: MaxBackups,  // 最大备份文件数量
		MaxAge:     MaxAge,      // 文件保留天数
		LocalTime:  true,
		Compress:   true, // 是否压缩
	})
	logger.SetLevel(level)       // 设置日志级别
	logger.SetReportCaller(true) // 打印调用者信息

	// 根据文件名设置对应的前缀
	prefix := ""
	switch fileName {
	case logFiles["system"]:
		prefix = SystemPrefix
	case logFiles["error"]:
		prefix = ErrorPrefix
	case logFiles["oss"]:
		prefix = OssPrefix
	case logFiles["ont"]:
		prefix = OntPrefix
	case logFiles["olt"]:
		prefix = OltPrefix
	case logFiles["aaa"]:
		prefix = AaaPrefix
	case logFiles["chain"]:
		prefix = ChainPrefix
	case logFiles["alarm"]:
		prefix = AlarmPrefix
	case logFiles["nbi"]:
		prefix = NorthboundPrefix
	}

	logger.SetFormatter(&CustomFormatter{Prefix: prefix})
	//logger.AddHook(&ConsoleHook{ConsoleLogger})
	return logger
}

// ConsoleHook 用于将日志同时输出到控制台
type ConsoleHook struct {
	logger *logrus.Logger
}

// Levels 返回钩子支持的日志级别
func (hook *ConsoleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 将日志条目写入控制台
func (hook *ConsoleHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	hook.logger.Out.Write([]byte(line))
	return nil
}

func PanicHandler() {
	if r := recover(); r != nil {
		// Create or open the log file
		l, err := os.OpenFile("./logs/panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to open panic log file err :", err)
		}
		defer l.Close()

		// Create a multi-writer to log to both file and stdout
		multiWriter := io.MultiWriter(l, os.Stdout)
		_, file, line, _ := runtime.Caller(2)
		log.SetOutput(multiWriter)
		log.Printf("Panic occurred: %v\n", r)
		log.Printf("File: %s\n", file)
		log.Printf("Line: %d\n", line)
		log.Printf("Stack Trace: %s\n", string(debug.Stack()))
	}
}

func SendSyslog(level SyslogRemoteLevel, msg string) {
	syslogServer, err := conf.Cfg.Section("server").GetKey("SYSLOG_SVR")
	if err != nil {
		return
	}
	w, err := srslog.Dial("udp", syslogServer.Value(), srslog.LOG_INFO|srslog.LOG_USER, "kogen")
	if err != nil {
		log.Fatalf("Failed to connect to syslog: %v", err)
	}
	defer w.Close()

	switch level {
	case SYSLOG_REMOTE_ALERT:
		w.Alert(msg)
	case SYSLOG_REMOTE_CRITICAL:
		w.Crit(msg)
	case SYSLOT_REMOTE_ERROR:
		w.Err(msg)
	case SYSLOG_REMOTE_INFO:
		w.Info(msg)
	default:
		w.Info(msg)
	}
}
