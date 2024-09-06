package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"gocv.io/x/gocv"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type _FrameExtractSuite struct {
	suite.Suite
}

func TestFrameExtract(t *testing.T) {
	suite.Run(t, new(_FrameExtractSuite))
}

func (my *_FrameExtractSuite) SetupSuite() {

}

func (my *_FrameExtractSuite) TestMetadata() {
	defer TrackTime()()
	url := "testdata/test_video.mp4"
	url = "http://file.te0.cn/exam.mp4"
	url = "testdata/exam.mp4"
	name := strings.TrimSuffix(filepath.Base(url), filepath.Ext(url))
	video, err := gocv.VideoCaptureFile(url)
	my.Require().Nil(err)

	frames := video.Get(gocv.VideoCaptureFrameCount)
	fps := video.Get(gocv.VideoCaptureFPS)
	duration := frames / fps

	for i := 0.0; i < math.Ceil(duration); i++ {
		img := gocv.NewMat()
		video.Set(gocv.VideoCapturePosFrames, i*fps)
		video.Read(&img)

		base := fmt.Sprintf("out/%s", name)
		_ = os.MkdirAll(base, os.ModePerm)
		src := fmt.Sprintf("%s/frame_%f.jpg", base, i)
		gocv.IMWrite(src, img)
	}

	my.T().Log(frames, fps, duration)
}
