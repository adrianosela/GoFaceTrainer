package main

import (
	"log"

	"github.com/adrianosela/GoFaceTrainer/trainer"
)

func main() {
	t, err := trainer.NewFaceTrainer(trainer.Settings{
		FaceboxAddress: "http://localhost:8080",
		FacesSrcDir:    ".",
	})
	if err != nil {
		log.Fatal(err)
	}
	t.Run()
}
