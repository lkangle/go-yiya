package compress

import "errors"

type pngLessener struct {
}

func (j *pngLessener) Compress(input, output string, quality int) error {
	return errors.New("暂不支持")
}
