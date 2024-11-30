package compress

import (
	"bytes"
	"embed"
	"io"
	"log"
	"os"
	"yiya-v2/backend/utils"
)

//go:embed all:bin
var assets embed.FS

func LoadPngquant() {
	binBytes, err := assets.ReadFile("bin/" + utils.GetPngquantFilename())
	if err != nil {
		log.Println("read embed file fail: " + err.Error())
		return
	}
	filepath := utils.GetPngquantFullPath()

	_, err = os.Stat(filepath)
	if err == nil {
		log.Println("可执行文件已存在: " + filepath)
		return
	}

	file, err := utils.CreateFile(filepath)
	if err != nil {
		log.Println("create file fail: " + err.Error())
		return
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewBuffer(binBytes))

	if err != nil {
		log.Println("copy file fail: " + err.Error())
	} else {
		log.Println("copy file success: " + filepath)
	}
}
