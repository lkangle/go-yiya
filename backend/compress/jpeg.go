package compress

import (
	"image/jpeg"
	"os"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"
)

type jpegLessener struct {
}

func (j *jpegLessener) Compress(input, output string, quality int) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	outFile, err := utils.CreateFile(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	q := j.mapQuality(quality)

	options := &jpeg.Options{Quality: q}
	err = jpeg.Encode(outFile, img, options)
	return err
}

func (j *jpegLessener) mapQuality(q int) int {
	if q == types.QuaNormal {
		return 72
	}
	return 75
}
