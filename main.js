// 引入electron并创建一个Browserwindow
const {app, BrowserWindow} = require('electron')
const path = require('path')
// const url = require('url')
// const fs = require('fs')
var child = require('child_process').execFile;
const isDev = require('electron-is-dev')
var ginProc = null

// 保持window对象的全局引用,避免JavaScript对象被垃圾回收时,窗口被自动关闭.
let mainWindow
//判断命令行脚本的第二参数是否含--debug
// const debug = /--debug/.test(process.argv[2]);
// function makeSingleInstance () {
//     if (process.mas) return;
//     app.requestSingleInstanceLock();
//     app.on('second-instance', () => {
//         if (mainWindow) {
//             if (mainWindow.isMinimized()) mainWindow.restore()
//             mainWindow.focus()
//         }
//     })
// }

// 启动 gin server，通过子进程执行已经将go打包好的exe文件（打包阶段）
function startServerEXE() {
  let oldUrl = path.join(__dirname, '/AltoServer.exe')
  console.log("startServerEXE  oldUrl= ",oldUrl)
  let newUrl = path.join(process.cwd(),'/resources/AltoServer.exe')
  console.log("startServerEXE  newUrl= ",newUrl)
  const exe = isDev?oldUrl:newUrl
  // ginProc = require('child_process').execFile(exe)

  ginProc = child(exe, function(err, data) {
    if(err){
        console.error(err);
        return;
    }

    console.log(data.toString());
  });

  if (ginProc != null) {
      console.log('Gin server start success!')
  }
  // fs.stat(oldUrl, function (err) {
  //   if (err) {
  //     ginProc = require('child_process').execFile(newUrl)
  //     if (ginProc != null) {
  //         console.log('Gin server start success in production mode!')
  //     }
  //   } else {
  //     ginProc = require('child_process').execFile(oldUrl)
  //     if (ginProc != null) {
  //         console.log('Gin server start success in debug mode!')
  //     }
  //   }
  // })
  
}

// 停止 gin server 函数
function stopServer() {
  ginProc.kill()
  console.log('kill Gin server  success')
  // console.log('kill Gin server path = ',path.join(__dirname, '/conf'))
  // fs.rm(path.join(__dirname, '/conf/'),{ recursive: true},function(err) {
  //   if(err){
  //       console.error(err);
  //       return;
  //   }
  // })
  // fs.rm(path.join(__dirname, '/db/'),{ recursive: true},function(err) {
  //   if(err){
  //       console.error(err);
  //       return;
  //   }
  // })

  ginProc = null
}
// 初始化函数
function initApp() {
  
  createWindow();
  startServerEXE();
}
function createWindow () {
  const windowOptions = {
    width: 800,
    height: 600,
    //frame:false,
  };
  mainWindow = new BrowserWindow(windowOptions);

  /* 
   * 加载应用-----  electron-quick-start中默认的加载入口
    mainWindow.loadURL(url.format({
      pathname: path.join(__dirname, 'index.html'),
      protocol: 'file:',
      slashes: true
    }))
  */
  // 加载应用----适用于 react 项目
  //const urlLocation = isDev? 'http://localhost:3000' :`file://${__dirname}/build/index.html`
  const urlLocation = isDev? 'http://localhost:3000' :'http://localhost:5600'
  mainWindow.loadURL(urlLocation);
  // mainWindow.loadURL(path.join('file://', __dirname, '/build/index.html'));
  
  // 打开开发者工具，默认不打开
  if(isDev){
    mainWindow.webContents.openDevTools()
    //npm install --save-dev devtron --force
    //require('devtron').install();
  }

  // 关闭window时触发下列事件.
  mainWindow.on('closed', function () {
    mainWindow = null
  })
}
//makeSingleInstance();
app.commandLine.appendSwitch("--disable-http-cache");
// 当 Electron 完成初始化并准备创建浏览器窗口时调用此方法
app.on('ready', initApp)

// 所有窗口关闭时退出应用.
app.on('window-all-closed', function () {
  // macOS中除非用户按下 `Cmd + Q` 显式退出,否则应用与菜单栏始终处于活动状态.
  if (process.platform !== 'darwin') {
    app.quit()
  }
  stopServer()
})

app.on('activate', function () {
   // macOS中点击Dock图标时没有已打开的其余应用窗口时,则通常在应用中重建一个窗口
  if (mainWindow === null) {
    createWindow()
  }
})

// 你可以在这个脚本中续写或者使用require引入独立的js文件.