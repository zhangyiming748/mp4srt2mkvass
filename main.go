package main

import (
	"mp4srt2mkvass/merge"
)

func main() {
	folderPath := "/data" // 指定文件夹路径
	extension := ".mp4"   // 指定扩展名

	merge.MkvWithAss(folderPath, extension)
}
