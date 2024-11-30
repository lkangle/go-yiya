package compress

import (
	"fmt"
	"testing"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"
)

func TestJp(t *testing.T) {
	p := "C:/Users/likangle/Downloads/卡2-3x.jpg"

	info := utils.GetBaseInfo(p)

	out, err := DoCompress(info, types.CompressOptions{
		Quality: 1,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	size := utils.GetSize(out)

	r := 1.0 - float32(size)/float32(info.Size)
	fmt.Println(out)
	fmt.Printf("out size: %d, origin size %d, dd rate %.2f\n", size, info.Size, r)
}

func TestLoadPngquant(t *testing.T) {
	LoadPngquant()
}
