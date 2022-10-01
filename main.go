package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var namefile = "Win 10 LTSC by Aiko.iso"

// check file exist if file Win 10 LTSC by Aiko.iso exist then print file exist and exit
func fileExists(filename string) {
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File exist")
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

var (
	version  = "0.0.1"
	codename = "Auto Download"
	intro    = "Tool by Aiko"
)

var (
	printVersion = flag.Bool("version", false, "show version")
)

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}

func main() {
	flag.Parse()
	showVersion()
	if *printVersion {
		return
	}
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
