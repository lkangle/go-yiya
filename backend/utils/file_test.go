package utils

import (
	"fmt"
	"testing"
)

func TestMime(t *testing.T) {
	fmt.Printf("vaild, jpg %t, png %t, webp %t, svg %t/n",
		IsValidImage("bigtitle_1@2x.jpg"),
		IsValidImage("bigtitle_1@2x.png"),
		IsValidImage("bigtitle_1@2x.webp"),
		IsValidImage("bigtitle_1@2x.svg"))

}

func TestParsePath(t *testing.T) {
	list := ParseDropPaths([]string{"C:/Users/likangle/Desktop/房间赛资格任务"})

	for _, i := range list {
		fmt.Println(i.Filename, i.Path)
	}
}

func TestCopyToTemp(t *testing.T) {
	f := "/Users/likangle/Desktop/MyProjects/go-yiya/test.png"

	o := CopyToTemp(f)
	fmt.Println(o)
}

func TestParseDropPaths(t *testing.T) {
	infos := ParseDropPaths([]string{
		"C:\\Users\\likangle\\Desktop\\房间赛资格任务",
	})

	for _, inf := range infos {
		out := GetOutputWithSuffix(inf, "")
		fmt.Println(inf.Path, out)
	}
}
