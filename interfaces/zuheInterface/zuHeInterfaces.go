package main

import (
	"fmt"
	"io"
)

// 定义基本接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// 组合接口：ReadWriter = Reader + Writer
type ReadWriter interface {
	Reader
	Writer
}

// 组合接口：ReadWriteCloser = Reader + Writer + Closer
type ReadWriteCloser interface {
	ReadWriter
	Closer
}

// 实现具体类型
type File struct {
	name string
	data []byte
	pos  int
}

func (f *File) Read(p []byte) (n int, err error) {
	if f.pos >= len(f.data) {
		return 0, io.EOF
	}
	n = copy(p, f.data[f.pos:])
	f.pos += n
	return n, nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *File) Close() error {
	fmt.Printf("文件 %s 已关闭\n", f.name)
	return nil
}

func main() {
	// 创建File实例
	file := new(File)
	file.name = "file.txt"
	fmt.Printf("file 类型%T\n fila.name 是%s\n", file, file.name)
	// File自动满足ReadWriteCloser接口
	var rwc ReadWriteCloser = file

	// 使用组合接口的方法
	data := []byte("Hello, Go!")
	n, _ := rwc.Write(data)
	fmt.Printf("写入 %d 字节\n", n)

	buf := make([]byte, 10)
	n, _ = rwc.Read(buf)
	fmt.Printf("读取 %d 字节: %s\n", n, buf[:n])

	err := rwc.Close()
	if err != nil {
		return
	}

	// 单独使用ReadWriter接口
	var rw ReadWriter = file
	_, err = rw.Write([]byte("More data"))
	if err != nil {
		return
	}

	// 类型断言检查
	if closer, ok := rw.(Closer); ok {
		fmt.Println("rw 也实现了 Closer")
		closer.Close()
	} else {
		fmt.Println("rw 没有实现 Closer")
	}
}
