# cinema : the lightweight video editor for golang 🤠

warning: still in beta, so some things might not be perfect

## Overview

cinema is a super simple video editor that supports video io, video trimming, and resizing. it is dependent on ffmpeg, an advanced command-line tool used for handling video, audio, and other multimedia files and streams. 

## Install
You must have FFMPEG installed on your machine! Make sure `ffmpeg` and `ffprobe` are available from the command line on your machine.
```
go get github.com/jtguibas/cinema
```

## Example Usage

```golang
func main() { // cinema/test/test.go

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
```

## Documentation
| Function | Description |
| --- | --- |
| cinema.MakeVideo(filepath string) | creates Video object based on filepath, will load relevant metadata using ffprobe |
| v.Render(output string) | render the video based on your configuration based on the output string |
| v.Trim(start float64, end float64) | trim video based in seconds |
| v.SetStart(start float64) | set the start of the video in seconds |
| v.SetEnd(end float64) | set the end of the video in seconds |
| v.SetSize(width int64, height int64) | set the size of the video in piexels |
| v.SetWidth(width int64) | set the width of the video in pixels |
| v.SetHeight(height int64) | set the height of the video in pixels |
| v.GetFilepath() string| get the filepath for the current video object |
| v.GetStart() float64 | get the start in seconds for the current video object|
| v.GetEnd() float64 | get the end in seconds for the current video object |
| v.GetWidth() int64 | get the width for the current video object in pixels |
| v.GetHeight() int64  | get the height for the current video object in pixels  |
| v.GetDuration() float64  | get the duration of the current video in seconds |
