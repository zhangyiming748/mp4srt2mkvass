package merge

import (
	"fmt"
	"github.com/zhangyiming748/FastMediaInfo"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func MkvWithAss(dir, pattern string) {
	files, err := getFilesWithExtension(dir, pattern)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		srt := strings.Replace(file, pattern, ".srt", 1)
		//srt = strings.Join([]string{"\"", srt, "\""}, "")
		//file = strings.Join([]string{"\"", file, "\""}, "")
		slog.Info("加引号", slog.String("file", file), slog.String("srt", srt))
		if isExist(srt) {
			output := strings.Replace(file, pattern, "_with_subtitle.mkv", 1)
			p := FastMediaInfo.GetStandMediaInfo(file)
			width, _ := strconv.Atoi(p.Video.Width)
			height, _ := strconv.Atoi(p.Video.Height)
			slog.Info("获取到的分辨率", slog.String("文件路径", file), slog.Int("width", width), slog.Int("height", height))
			crf := FastMediaInfo.GetCRF("vp9", width, height)
			if crf == "" {
				crf = "31"
			}
			//cmd := exec.Command("ffmpeg", "-i", file, "-itsoffset", "1", "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
			cmd := exec.Command("ffmpeg", "-i", file, "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
			fmt.Printf("生成的命令: %s\n", cmd.String())
			combinedOutput, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("命令执行失败: %s\n", err.Error())
				continue
			} else {
				fmt.Printf("命令成功执行: %s\n", string(combinedOutput))
				os.Remove(file)
			}
		}
	}
}
func isExist(fp string) bool {
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		fmt.Printf("%s 对应的字幕文件不存在\n", fp)
		return false
	} else {
		fmt.Printf("%s 对应的字幕文件存在\n", fp)
		return true
	}
}
func getFilesWithExtension(folderPath string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
