// demoFileOperation project main.go
package main

import (
	"os"
)

func main() {
	//创建一个文件夹
	//os.Mkdir("test", 0777)
	//os.MkdirAll("C:\\filetest\\pic\\jpg", 0777)

	//删除文件或文件夹
	os.Remove("test\\test.txt")
	os.RemoveAll("C:\\filetest")

	//创建文件,写文件
	fout, _ := os.Create("test\\test.txt")
	for i := 0; i < 3; i++ {
		fout.WriteString("it's a test\r\n")
		fout.Write([]byte("this is a test\r\n"))
	}

	//将test中的内容复制到test1中
	f1, _ := os.Open("test\\test.txt")
	f2, _ := os.Create("test\\test1.txt")
	defer f1.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		f2.Write(buf[:n])
	}

}
