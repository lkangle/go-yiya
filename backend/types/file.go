package types

type CompressResult struct {
	InputPath     string `json:"$input_path"`      // 输入文件的原始路径
	InputTempPath string `json:"$input_temp_path"` // 输入文件的临时保存路径 override时存在
	Path          string `json:"path"`             // 压缩后的文件路径
	Size          int64  `json:"size"`             // 压缩后的文件尺寸

	Code    int    `json:"code"`    // 压缩结果code
	Success bool   `json:"success"` // 是否压缩成功
	Message string `json:"message"` // 压缩失败的异常信息
}

type ImageFileInfo struct {
	Id          string `json:"id"`
	DirPath     string `json:"dir_path"` // 拖拽进来一组文件的共同父目录
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"` // image/png  image/jpeg 之类
}
