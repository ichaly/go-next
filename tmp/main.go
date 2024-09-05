package main

import (
	"fmt"
	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
	"image/color"
	"log"
)

//export PKG_CONFIG_PATH="/opt/homebrew/Cellar/dlib/19.24.6/lib/pkgconfig:$PKG_CONFIG_PATH"
//export PKG_CONFIG_PATH="/opt/homebrew/Cellar/jpeg-turbo/3.0.3/lib/pkgconfig:$PKG_CONFIG_PATH"
//
//export C_INCLUDE_PATH="/opt/homebrew/Cellar/dlib/19.24.6/include:$C_INCLUDE_PATH"
//export CPLUS_INCLUDE_PATH="/opt/homebrew/Cellar/dlib/19.24.6/include:$CPLUS_INCLUDE_PATH"
//export C_INCLUDE_PATH="/opt/homebrew/Cellar/jpeg-turbo/3.0.3/include:$C_INCLUDE_PATH"
//export CPLUS_INCLUDE_PATH="/opt/homebrew/Cellar/jpeg-turbo/3.0.3/include:$CPLUS_INCLUDE_PATH"
//
//export DYLD_LIBRARY_PATH="/usr/local/lib:/opt/homebrew/Cellar/dlib/19.24.6/lib:$DYLD_LIBRARY_PATH"
//export LIBRARY_PATH="/opt/homebrew/Cellar/dlib/19.24.6/lib:$LIBRARY_PATH"
//export DYLD_LIBRARY_PATH="/opt/homebrew/Cellar/jpeg-turbo/3.0.3/lib:$DYLD_LIBRARY_PATH"
//export LIBRARY_PATH="/opt/homebrew/Cellar/jpeg-turbo/3.0.3/lib:$LIBRARY_PATH"

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	idolsDir  = dataDir + "/idols"
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

// 图片中的人名
var labels = []string{
	"萧敬腾",
	"周杰伦",
	"Unknow",
	"王力宏",
	"陶喆",
	"林俊杰",
}

func main() {
	// initializes the identifier
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	// 调用该方法，传入路径。返回面部数量和任何错误
	faces, err := rec.RecognizeFile(idolsDir + "/heyin.jpeg")
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	// 打印人脸数量
	fmt.Println("图片人脸数量: ", len(faces))

	var samples []face.Descriptor
	var peoples []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// 每张脸唯一 id
		peoples = append(peoples, int32(i))
	}

	// 传入样例到识别器
	rec.SetSamples(samples, peoples)

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load("testdata/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	//// load video file
	//video, err := gocv.VideoCaptureFile("testdata/test_video.mp4")
	//if err != nil {
	//	println("Cannot open video file")
	//	return
	//}
	//defer video.Close()

	// open webcam
	video, err := gocv.OpenVideoCapture(0)
	if err != nil {
		println("Cannot webcam")
		return
	}
	defer video.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	for {
		if ok := video.Read(&img); !ok || img.Empty() {
			fmt.Println("无法读取图像帧")
			break
		}

		//dist := face.SquaredEuclideanDistance(f.Descriptor, descriptor)
		//fmt.Println("欧拉距离 = ", dist)

		//gocv.CvtColor(img, &img, gocv.ColorBGRToGray)

		//detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 1)
			// 在图片上画人脸框
			//pt := image.Pt(r.Rectangle.Min.X, r.Rectangle.Min.Y-20)
			//gocv.PutText(&background, "jay", pt, gocv.FontHersheyPlain, 2, c, 2)

			//// 计算特征值之间的欧拉距离
			//dist := face.SquaredEuclideanDistance(f.Descriptor, descriptor)
			//fmt.Println("欧拉距离 = ", dist)

			buff, err := gocv.IMEncode(".jpg", img.Region(r))
			if err != nil {
				fmt.Println("encoding to jpg err:%v", err)
				break
			}
			RecognizePeopleFromMemory(rec, buff.GetBytes())
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)

		// 等待用户按下ESC键退出
		if window.WaitKey(1) == 27 {
			return
		}
	}
}

func RecognizePeopleFromMemory(rec *face.Recognizer, img []byte) {
	people, err := rec.RecognizeSingle(img)
	if err != nil {
		log.Println("无法识别: %v", err)
		return
	}
	if people == nil {
		log.Println("图片上不是一张脸")
		return
	}
	peopleID := rec.ClassifyThreshold(people.Descriptor, 0.2)
	if peopleID < 0 {
		log.Println("无法区分")
		return
	}
	fmt.Println(peopleID)
	fmt.Println(labels[peopleID])
}
