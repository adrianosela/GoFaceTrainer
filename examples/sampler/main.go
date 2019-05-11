package main

import (
	"log"

	"github.com/adrianosela/GoFaceTrainer/sampler"
)

func main() {
	s, err := sampler.NewFaceSampler(sampler.Settings{
		CaptureDeviceID: 0,
		FaceAlgoPath:    "../../face_algos/haarcascade_frontalface_default.xml",
		WindowTitle:     "Face Sampler",
		NSamples:        25,
		SaveSamplesDir:  ".",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.Run()
}
