package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func dowfile() {
	// tạo file mới
	file, err := os.Create("Win 10 LTSC by Aiko.iso")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// tải file
	resp, err := http.Get("https://go.microsoft.com/fwlink/p/?LinkID=2195404&clcid=0x409&culture=en-us&country=US")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// đọc file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded a file")
}

func main() {
	go dowfile()
}
