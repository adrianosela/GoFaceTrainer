package main

import (
	"log"
	"os"
	"strconv"

	"github.com/adrianosela/GoFaceTrainer/detector"
	"github.com/adrianosela/GoFaceTrainer/sampler"
	"github.com/adrianosela/GoFaceTrainer/trainer"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("usage: go run main.go [sample/train/run] ...args")
	}

	switch os.Args[1] {
	case "sample":
		var set sampler.Settings
		var err error

		if len(os.Args) != 6 {
			log.Fatal("usage: go run main.go sample <device_id> <face_algo_path> <nsamples> <save_dir>")
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

		if len(os.Args) != 5 {
			log.Fatal("usage: go run main.go train <facebox_url> <faces_src_dir> <person_name>")
		}

		set.FaceboxAddress = os.Args[2]
		set.FacesSrcDir = os.Args[3]
		set.PersonName = os.Args[4]

		t, err := trainer.NewFaceTrainer(set)
		if err != nil {
			log.Fatal(err)
		}
		t.Run()
		return
	case "run":
		var set detector.Settings
		var err error

		if len(os.Args) != 5 {
			log.Fatal("usage: go run main.go run <device_id> <facebox_url> <face_algo_path>")
		}

		if set.CaptureDeviceID, err = strconv.Atoi(os.Args[2]); err != nil {
			log.Fatalf("capture device not an integer: %s", err)
		}

		set.FaceboxAddress = os.Args[3]
		set.FaceAlgoPath = os.Args[4]

		d, err := detector.NewFaceDetector(set)
		if err != nil {
			log.Fatal(err)
		}
		defer d.Close()
		d.Run()
		return
	default:
		log.Fatal("usage: go run main.go [sample/train/run] ...args")
	}
}
