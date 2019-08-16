package cinema

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/tidwall/gjson"
)

type Video struct {
	filepath string
	width    int64
	height   int64
	start    float64
	end      float64
	duration float64
}

func MakeVideo(filepath string) Video {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filepath)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Error: FFPROBE did not work. Check the filename and thatt ffmpeg is correctly installed.", filepath)
		log.Fatal(err)
	}

	json := string(out)
	width := gjson.Get(json, "streams.0.width").Int()
	height := gjson.Get(json, "streams.0.height").Int()
	duration := gjson.Get(json, "format.duration").Float()

	return Video{filepath: filepath, width: width, height: height, start: 0, end: duration, duration: duration}
}

func (v *Video) Render(output string) {

	cmd := exec.Command("ffmpeg",
		"-y",
		"-i",
		v.filepath,
		"-vf",
		"scale="+strconv.Itoa(int(v.width))+":"+strconv.Itoa(int(v.height))+",setsar=1",
		"-ss",
		strconv.Itoa(int(v.start)),
		"-t",
		strconv.Itoa(int(v.end-v.start)),
		"-strict",
		"-2",
		output)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func (v *Video) Trim(start float64, end float64) {
	v.start = start
	v.end = end
}

// Setters
func (v *Video) SetStart(start float64) {
	v.start = start
}

func (v *Video) SetEnd(end float64) {
	v.end = end
}

func (v *Video) SetSize(width int64, height int64) {
	v.width = width
	v.height = height
}

func (v *Video) SetWidth(width int64) {
	v.width = width
}

func (v *Video) SetHeight(height int64) {
	v.height = height
}

// Getters
func (v *Video) GetFilepath() string {
	return v.filepath
}

func (v *Video) GetStart() float64 {
	return v.start
}

func (v *Video) GetEnd() float64 {
	return v.end
}

func (v *Video) GetWidth() int64 {
	return v.width
}

func (v *Video) GetHeight() int64 {
	return v.height
}

func (v *Video) GetDuration() float64 {
	return v.duration
}
