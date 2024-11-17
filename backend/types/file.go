package types

type ImageFileInfo struct {
	Id          string `json:"id"`
	GroupId     string `json:"group_id"` // 同一批拖拽进来的文件 才会有组 且为同一组
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	TempPath    string `json:"temp_path"` // 临时原文件的位置 在覆盖型的压缩时用于恢复
	OriginSize  int64  `json:"originSize"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"` // image/png  image/jpeg 之类
	Message     string `json:"message"`     // 存一些额外的信息 如错误
}

type ImageInfoList []ImageFileInfo
