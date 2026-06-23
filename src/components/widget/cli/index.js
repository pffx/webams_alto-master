import React from 'react';
import CLIWindow from '@nokia-csf-uxr/ccfk/CLIWindow';
import { Terminal } from 'xterm';
import { WebglAddon } from 'xterm-addon-webgl';
import { FitAddon } from "xterm-addon-fit";
import { AttachAddon } from 'xterm-addon-attach'
import { getContext } from '@nokia-csf-uxr/ccfk/common';
const DUMMY_TEXT = "Socket connected failed!";
const enter = window.navigator.platform === 'win32' ? '\r\n' : '\n';

const CLIWindowWidget = (props,terminalRef) => {
  //const terminalRef = React.useRef();
  const containerRef = React.useRef();
  const theme = getContext(({ theme }) => theme.CLIWINDOW);
  const fitAddon = new FitAddon();
  // const attachAddon = new AttachAddon(props.socket);
  React.useLayoutEffect(() => {
    terminalRef.current = new Terminal({
      fontFamily: `'Fira Code', monospace`,
      fontSize: 12,
      lineHeight: 1.0,
      convertEol: true,
      cursorBlink: true,
    });
    //console.log("containerRef  = ",containerRef.current);
    terminalRef.current.open(containerRef.current);
    terminalRef.current.loadAddon(new WebglAddon());
    terminalRef.current.loadAddon(fitAddon);
    // console.log("CLIWindowWidget    props.socket  : ",props.socket)
    if(props.socket){
      terminalRef.current.loadAddon(new AttachAddon(props.socket));
      // props.socket.onmessage = function (evt) {
      //   console.log("CLIWindowWidget   evt = ",evt)
      //   var messages = evt.data.split('\n');
      //   terminalRef.current.write(enter+messages);
      //   // terminalRef.current.write(enter);
      // };
    }else{
      //Write text inside the terminal
      terminalRef.current.write(DUMMY_TEXT);
    }
    
    // terminalRef.current.onKey(key => {
    //   const char = key.domEvent.key;
    //   console.log("CLIWindowWidget   key = ",key)
    //   // if (char === "Enter") {
    //   //   console.log(terminalRef.current)
    //   //   terminalRef.current.prompt();
    //   // } else
    //    if (char === "Backspace") {
    //     terminalRef.current.write("\b \b");
    //   } else {
    //     terminalRef.current.write(char);
    //   }
    // });
    terminalRef.current.onData(key => {  // 粘贴的情况
      if(key.length > 1) terminalRef.current.write(key)
    })
    fitAddon.fit();

  }, []);
  React.useLayoutEffect(() => {
    //Styling
    terminalRef && terminalRef.current && terminalRef.current.setOption("theme", {
      background: theme.background,
      foreground: theme.text,
      cursor: theme.cursor,
      selection: theme.selectedBackground,
    });
  }, [theme]);
  React.useEffect(() => {
    const resize = () => {
      // triggers a reize in xterm.js
      fitAddon.fit();
    }
    window.addEventListener('resize', resize);
    return () => {
      window.removeEventListener('resize', resize);
    }
  }, [])
  
  return (
    <CLIWindow {...props}  ref={containerRef} />
  );
};
export default React.forwardRef(CLIWindowWidget);