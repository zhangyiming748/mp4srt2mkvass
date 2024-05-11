package main

import (
	"log/slog"
	"mp4srt2mkvass/constant"
	"mp4srt2mkvass/merge"
	"os"
)

func main() {
	if r := os.Getenv("root"); r == "" {
		slog.Info("没有设置root变量,使用默认")
	} else {
		constant.SetRoot(r)
	}
	extension := ".mp4" // 指定扩展名

	merge.MkvWithAss(constant.GetRoot(), extension)
}
