const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

// 获取当前工作目录（项目根目录）
const projectRoot = process.cwd();
// 构建build文件夹的绝对路径
const buildDir = path.join(projectRoot, 'build');
const buildDirName = path.basename(buildDir);

// 检查build文件夹是否存在
if (!fs.existsSync(buildDir)) {
  console.error(`错误: 找不到build文件夹，路径: ${buildDir}`);
  console.error('请确认是否已执行构建命令，如npm run build');
  process.exit(1);
}

// 确定目标目录
const isWindows = process.platform === 'win32';
const targetDir = isWindows 
  ? path.join(projectRoot, 'dist', 'win-unpacked', 'resources') 
  : path.join(projectRoot, 'dist', 'linux-unpacked', 'resources');

// 确保目标目录存在
if (!fs.existsSync(targetDir)) {
  fs.mkdirSync(targetDir, { recursive: true });
  console.log(`已创建目标目录: ${targetDir}`);
}

// 执行复制命令 - 使用系统原生命令确保兼容性
try {
  console.log(`正在将build文件夹从 ${buildDir} 复制到 ${targetDir}...`);
  
  let copyCommand;
  if (isWindows) {
    // Windows使用xcopy命令，/E复制所有子目录和文件，/I假设目标是文件夹
    copyCommand = `xcopy "${buildDir}" "${path.join(targetDir, buildDirName)}" /E /I`;
  } else {
    // Linux/macOS使用cp命令，-R递归复制目录
    copyCommand = `cp -R "${buildDir}" "${targetDir}"`;
  }
  
  // 执行系统原生复制命令
  execSync(copyCommand, {
    stdio: 'inherit',
    shell: true
  });
  
  console.log('Build文件夹复制成功');
} catch (err) {
  console.error('复制build文件夹失败:', err.message);
  process.exit(1);
}
    