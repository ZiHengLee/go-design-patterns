package 组合模式

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	folder := Folder{name: "文件夹A"}
	file1 := &File{name: "文件1", text: "文件1的内容"}
	file2 := &File{name: "文件2", text: "文件2的内容Hello World!"}
	fmt.Println("文件1字节数:", file1.countBytes())
	fmt.Println("文件2字节数:", file2.countBytes())
	folder.add(file1)
	folder.add(file2)
	fmt.Println("文件夹A字节数:", folder.countBytes())
}
