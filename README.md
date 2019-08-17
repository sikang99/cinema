# cinema : a lightweight video editor for golang 🤠

## Overview [![GoDoc](https://godoc.org/github.com/jtguibas/cinema?status.svg)](https://godoc.org/github.com/jtguibas/cinema)

cinema is a super simple video editor that supports video io, video trimming, and resizing. it is dependent on ffmpeg, an advanced command-line tool used for handling video, audio, and other multimedia files and streams. 

## Install
You must have [FFMPEG](https://ffmpeg.org/download.html) installed on your machine! Make sure `ffmpeg` and `ffprobe` are available from the command line on your machine.
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
	v.Render("test_output.mov") // notice how it can convert formats with ease

	// testing all getters
	fmt.Println(v.GetFilepath())
	fmt.Println(v.GetStart())
	fmt.Println(v.GetEnd())
	fmt.Println(v.GetWidth())
	fmt.Println(v.GetHeight())
	fmt.Println(v.GetDuration())
}
```

### TODO

- [ ] add concatenation support
- [ ] improve godoc documentation
- [ ] add cropping support
- [ ] expand to audio