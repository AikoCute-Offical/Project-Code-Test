package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var namefile = "Win 10 LTSC by Aiko.iso"

// check file exist or not
func fileExists(filename string) {
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File", filename, "exists")
		fmt.Println("Press enter to exit")
		fmt.Scanln()
		os.Exit(0)
	}
}

// dowload file from link
func downloadFile(url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(namefile)
	if err != nil {
		return err
	}
	defer out.Close()

	// hiện thị đã tải được bao nhiêu % file
	var downloaded int64
	var percent int
	for {
		n, err := io.CopyN(out, resp.Body, 1024*1024)
		downloaded += n
		percent = int(float64(downloaded) / float64(resp.ContentLength) * 100)
		fmt.Printf("Downloaded %d%% \r", percent)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}

func main() {
	fileExists(namefile)
	// hiển thị đang tải và tên file (namefile)
	fmt.Println("Downloading", namefile, "...")
	url := "https://go.microsoft.com/fwlink/p/?LinkID=2195404&clcid=0x409&culture=en-us&country=US"
	err := downloadFile(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("File downloaded")
}
