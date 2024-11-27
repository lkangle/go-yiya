package compress

import (
	"io"
	"log"
	"net/http"
	"yiya-v2/backend/utils"

	"github.com/tidwall/gjson"
)

type binInfo struct {
	Version     string
	DownloadUrl string
}

func getLatestBin() *binInfo {
	tgfilename := utils.GetPngquantFilename()
	if tgfilename == "" {
		panic("不支持的平台")
	}

	resp, err := http.Get("https://api.github.com/repos/lkangle/pngquant-c/releases/latest")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	r := gjson.ParseBytes(body)
	version := r.Get("name").String()
	if version == "" {
		return nil
	}

	info := binInfo{}

	assets := r.Get("assets").Array()
	for _, asset := range assets {
		filename := asset.Get("name").String()
		downloadUrl := asset.Get("browser_download_url").String()
		if tgfilename == filename && downloadUrl != "" {
			info.Version = version
			info.DownloadUrl = downloadUrl
			return &info
		}
	}
	return nil
}

func LoadPngquant() {
	info := getLatestBin()
	if info == nil {
		log.Println("png压缩命令行文件不存在..")
		return
	}
	oldVersion := utils.GetBinVersion()

	if oldVersion == "" || utils.CompareVersions(oldVersion, info.Version) < 0 {
		filepath := utils.GetPngquantFullPath()
		err := utils.DownloadFile(info.DownloadUrl, filepath)
		if err != nil {
			log.Println("下载失败: ", err.Error())
		} else {
			log.Println("下载成功: ", filepath)
		}

		utils.SetBinVersion(info.Version)
	}
}
