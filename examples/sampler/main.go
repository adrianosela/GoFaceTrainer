package main

import (
	"github.com/adrianosela/GoFaceTrainer/sampler"
	"log"
)

func main() {
	s, err := sampler.NewFaceSampler(&sampler.Settings{
		CaptureDeviceID: 0,
		FaceboxAddress:  "http://localhost:8080",
		FaceAlgoPath:    "../../face_algos/haarcascade_frontalface_default.xml",
		WindowTitle:     "Face Sampler",
		NSamples:        25,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.Run()
}
