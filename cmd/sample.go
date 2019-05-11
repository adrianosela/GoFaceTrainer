package main

import (
	"log"
	"os"
	"strconv"

	"github.com/adrianosela/GoFaceTrainer/sampler"
)

func main() {
	var set sampler.Settings
	var err error

	if len(os.Args) != 5 {
		log.Fatalf("usage: go run sample.go <device_id> <face_algo_path> <nsamples> <save_dir>")
	}

	if set.CaptureDeviceID, err = strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("capture device not an integer: %s", err)
	}
	set.FaceAlgoPath = os.Args[2]
	if set.NSamples, err = strconv.Atoi(os.Args[3]); err != nil {
		log.Fatalf("nsamples not an integer: %s", err)
	}
	set.SaveSamplesDir = os.Args[4]

	s, err := sampler.NewFaceSampler(set)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.Run()
}
