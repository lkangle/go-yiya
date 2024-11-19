package compress

import (
	"os"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"

	"github.com/chai2010/webp"
)

type webpLessener struct {
}

func (j *webpLessener) mapQuality(q int) int {
	if q == types.QuaNormal {
		return 70
	}
	return 75
}

func (j *webpLessener) Compress(input, output string, quality int) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := webp.Decode(file)
	if err != nil {
		return err
	}

	outFile, err := utils.CreateFile(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	q := j.mapQuality(quality)

	opt := &webp.Options{
		Lossless: false,
		Quality:  float32(q),
		Exact:    false,
	}
	return webp.Encode(outFile, img, opt)
}
