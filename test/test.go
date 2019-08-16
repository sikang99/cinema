package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jtguibas/cinema"
)

func main() {

	// loading the test video
	fmt.Println("Downloading Test Video...")
	video_url := "https://media.w3.org/2010/05/sintel/trailer.mp4"
	if err := DownloadFile("test_input.mp4", video_url); err != nil {
		panic(err)
	}

	// initializing the test video as a cinema video object
	v := cinema.MakeVideo("test_input.mp4")

	// testing all setters
	v.Trim(0, 10)
	v.SetStart(0)
	v.SetEnd(10)
	v.SetSize(400, 400)
	v.Render("test_output.mp4")

	// testing all getters
	fmt.Println(v.GetFilepath())
	fmt.Println(v.GetStart())
	fmt.Println(v.GetEnd())
	fmt.Println(v.GetWidth())
	fmt.Println(v.GetHeight())
	fmt.Println(v.GetDuration())
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
