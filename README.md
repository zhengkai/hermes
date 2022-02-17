# Hermes

借助 [FFmpeg](https://www.ffmpeg.org/) 在命令行下预览图片/视频的工具。[演示地址](https://www.youtube.com/watch?v=g7OHCnZ9fbk)

![screenshot](/static/screenshot.webp)

## 安装

[最新版本下载](https://github.com/zhengkai/hermes/releases/tag/v1.0.2)，请选择对应平台。

本工具基于 [FFmpeg](https://www.ffmpeg.org/download.html)，请确保已安装、命令行下输 `ffmpeg` 可以看到一堆版本信息。

确保终端已支持 24 位色，输入 `echo $COLORTERM` 可以看到 `24bit` 或 `truecolor`

## 使用

    ./hermes [-size 60x40] [-frames 1] [-seek 00:01:23.456] 文件名

所有参数除了文件名均为可选项。

默认使用终端窗口最大尺寸，也可以用 `-size` 指定输出大小，注意每个字符可以显示两个像素，所以高度是实际行数的一半，如 `60x40` 实际为 20 个字符的高度。

也可以设置环境变量 `HERMES_DEFAULT_SIZE` 作为 `-size` 的默认值。

`-seek` 跳到指定的时间位置播放，格式为 `00:00:00.000`，如果时分秒超过 59 会自动进位，  
如 `-seek 90` 相当于 `-seek 01:30`。

`-frames` 只播放前 n 帧。

## TODO

* docker
* telnet server
* 对比前后帧，优化输出流量
* output as nodejs/shell script
* 目前是通过管道方式获取每帧的 BMP 文件并转成 ANSI color，而且一阻塞就会丢帧，如果有更底层的方式就可以做成有快进/后退/暂停的完善播放器了
