package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"log/slog"
	"mp4srt2mkvass/constant"
	"mp4srt2mkvass/merge"
	"os"
	"strings"
)

func init() {
	setLog()
}
func main() {
	if r := os.Getenv("root"); r == "" {
		slog.Info("没有设置root变量,使用默认")
	} else {
		constant.SetRoot(r)
	}
	extension := ".mp4" // 指定扩展名

	merge.MkvWithAss(constant.GetRoot(), extension)
}
func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "mylog.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
