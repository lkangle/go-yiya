package types

type CompressResult struct {
	TempFile string `json:"temp_file"` // 可能存在的原文件复制的临时文件
	Path     string `json:"path"`      // 压缩后的文件路径
	Size     int    `json:"size"`
	Success  bool   `json:"success"` //是否压缩成功
	Message  string `json:"message"` // 压缩失败的异常信息
}

type ImageFileInfo struct {
	Id          string `json:"id"`
	GroupId     string `json:"group_id"` // 同一批拖拽进来的文件 才会有组 且为同一组
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"` // image/png  image/jpeg 之类
}

type ImageInfoList []ImageFileInfo
