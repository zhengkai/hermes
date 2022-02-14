# Hermes

命令行下预览视频文件的工具

[演示地址](https://www.youtube.com/watch?v=g7OHCnZ9fbk)

## 安装

[最新版本下载](https://github.com/zhengkai/hermes/releases/tag/v1.0.1)，请选择对应平台

本工具基于 [FFmpeg](https://www.ffmpeg.org/download.html)，请确保已安装、命令行下输 `ffmpeg` 可以看到一堆版本信息。

确保终端已支持 24 位色，输入 `echo $COLORTERM` 可以看到 `24bit` 或 `truecolor`

## 使用

    ./hermes [-size 60x40] 文件名

默认使用终端窗口最大尺寸，也可以用 `-size` 指定输出大小，注意每个字符可以显示两个像素，所以高度是实际行数的一半，如 `60x40` 实际为 20 个字符的高度。

也可以设置环境变量 `HERMES_DEFAULT_SIZE` 作为 `-size` 时的默认值。

## TODO

* telnet server
* 对比前后帧，优化输出流量
