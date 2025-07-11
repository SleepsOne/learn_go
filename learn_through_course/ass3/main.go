package main

import (
	"fmt"
	"io"
	"os"
)

var count int

type ReaderCustom struct {
	data []byte
	pos  int // this may changed so we have to pass the pointer into receiver function
}

func (rc *ReaderCustom) Read(p []byte) (n int, err error) {
	if rc.pos >= len(rc.data) {
		return 0, io.EOF
	}

	n = copy(p, rc.data[rc.pos:])
	rc.pos += n
	count++
	fmt.Println(count)

	return n, nil

}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Provide filename to read from")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer file.Close() // đảm bảo file được đóng sau khi dùng xong

	// Copy dữ liệu từ file (io.Reader) sang os.Stdout (io.Writer)
	n, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error copying file contents:", err)
		os.Exit(1)
	}

	fmt.Printf("\nSize of file: %d bytes\n", n)

	/*
		Sol 2 : read all the file by os.Readfile (*not recommended)
		bs, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		rc := &ReaderCustom{data: bs, pos: 0}

		_, err = io.Copy(os.Stdout, rc)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)

		}
	*/
}
