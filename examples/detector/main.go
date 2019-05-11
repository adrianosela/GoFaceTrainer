package main

import (
	"log"

	"github.com/adrianosela/GoFaceTrainer/detector"
)

func main() {
	d, err := detector.NewFaceDetector(detector.Settings{
		CaptureDeviceID: 0,
		FaceAlgoPath:    "../../face_algos/haarcascade_frontalface_default.xml",
		WindowTitle:     "Face Detector",
		FaceboxAddress:  "http://localhost:8080",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	d.Run()
}
