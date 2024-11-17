package utils

import (
	"fmt"
	"testing"
)

func TestMime(t *testing.T) {
	fmt.Printf("vaild, jpg %t, png %t, webp %t, svg %t/n",
		IsVaildImage("bigtitle_1@2x.jpg"),
		IsVaildImage("bigtitle_1@2x.png"),
		IsVaildImage("bigtitle_1@2x.webp"),
		IsVaildImage("bigtitle_1@2x.svg"))

}

func TestParsePath(t *testing.T) {
	list := ParseDropPaths([]string{"C:/Users/likangle/Desktop/房间赛资格任务"})

	for _, i := range list {
		fmt.Println(i.Filename, i.Path)
	}
}
