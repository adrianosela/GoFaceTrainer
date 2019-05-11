package main

import (
	"log"
	"os"
	"strconv"

	"github.com/adrianosela/GoFaceTrainer/sampler"
	"github.com/adrianosela/GoFaceTrainer/trainer"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("usage: go run cmd.go [sample/train/run] ...args")
	}

	switch os.Args[1] {
	case "sample":
		var set sampler.Settings
		var err error

		if len(os.Args) != 6 {
			log.Fatal("usage: go run cmd.go sample <device_id> <face_algo_path> <nsamples> <save_dir>")
		}

		if set.CaptureDeviceID, err = strconv.Atoi(os.Args[2]); err != nil {
			log.Fatalf("capture device not an integer: %s", err)
		}
		set.FaceAlgoPath = os.Args[3]
		if set.NSamples, err = strconv.Atoi(os.Args[4]); err != nil {
			log.Fatalf("nsamples not an integer: %s", err)
		}
		set.SaveSamplesDir = os.Args[5]

		s, err := sampler.NewFaceSampler(set)
		if err != nil {
			log.Fatal(err)
		}
		defer s.Close()
		s.Run()
		return
	case "train":
		var set trainer.Settings
		var err error

		if len(os.Args) != 4 {
			log.Fatal("usage: go run cmd.go train <facebox_url> <faces_src_dir>")
		}

		set.FaceboxAddress = os.Args[2]
		set.FacesSrcDir = os.Args[3]

		t, err := trainer.NewFaceTrainer(set)
		if err != nil {
			log.Fatal(err)
		}
		t.Run()
		return
	case "run":
		// TODO
	default:
		log.Fatal("usage: go run cmd.go [sample/train/run] ...args")
	}
}