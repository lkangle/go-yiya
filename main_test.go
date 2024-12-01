package main

import (
	"fmt"
	"testing"
	"yiya-v2/backend/services"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"
)

func TestCompressService(t *testing.T) {
	f := services.FileService()

	info := utils.GetBaseInfo("/Users/likangle/Desktop/MyProjects/go-yiya/test.png")

	r := f.CompressImage(info, types.CompressOptions{
		Override: false,
	})

	fmt.Println(r.Success, r)
}

func TestCompressOverride(t *testing.T) {
	f := services.FileService()

	info := utils.GetBaseInfo("/Users/likangle/Desktop/MyProjects/go-yiya/test.png")

	r := f.CompressImage(info, types.CompressOptions{
		Override: true,
	})

	fmt.Println(r.Success, r)
}
