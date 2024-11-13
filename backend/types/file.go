package types

type ImageFileInfo struct {
	Id          string `json:"id"`
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	OriginSize  int64  `json:"originSize"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"` // image/png  image/jpeg 之类
}
