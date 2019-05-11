package trainer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/machinebox/sdk-go/facebox"
)

const (
	defaultFaceboxAddr = "http://localhost:8080"
	defaultSrcFacesDir = "."
	defaultPersonName  = "person"
)

// FaceTrainer is an abstraction around the necessary Machinebox
// resources in order to train a model of a face or facial-expression
type FaceTrainer struct {
	faceboxClient *facebox.Client
	facesDir      string
	personName    string
}

// Settings contains necessary values to initialize a FaceTrainer
type Settings struct {
	// FaceboxAddress is the URL of the MachineBox to use - host:port,
	// (default is http://localhost:8080)
	FaceboxAddress string

	// FacesSrcDir is the directory where .jpg samples will be sources from
	// (default is the current directory ".")
	FacesSrcDir string

	// PersonName is the name of the person who'se face is being sampled
	// (default is "person")
	PersonName string
}

// NewFaceTrainer is the FaceTrainer constructor
func NewFaceTrainer(settings Settings) (*FaceTrainer, error) {
	if settings.FaceboxAddress == "" {
		settings.FaceboxAddress = defaultFaceboxAddr
	}
	if settings.FacesSrcDir == "" {
		settings.FacesSrcDir = defaultSrcFacesDir
	}
	if settings.PersonName == "" {
		settings.PersonName = defaultPersonName
	}
	return &FaceTrainer{
		faceboxClient: facebox.New(settings.FaceboxAddress),
		facesDir:      settings.FacesSrcDir,
		personName:    settings.PersonName,
	}, nil
}

// Run trains a machinebox/facebox model with all .jpg files in its FacesSrcDir
func (t *FaceTrainer) Run() error {
	files, err := ioutil.ReadDir(t.facesDir)
	if err != nil {
		return fmt.Errorf("could not read faces source directory: %s", err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".jpg") {
			fd, err := os.Open(fmt.Sprintf("%s/%s", t.facesDir, f.Name()))
			if err != nil {
				log.Printf("could not open img file %s: %s", f.Name(), err)
				continue // move on to next img if we can't open one
			}
			id := fmt.Sprintf("%s-%s", t.personName, strings.TrimPrefix(f.Name(), ".jpg"))
			if err := t.faceboxClient.Teach(fd, id, t.personName); err != nil {
				log.Printf("could not teach facebox using file %s: %s", f.Name(), err)
				continue
				/*
					move on to next img if we can't use one. Errors here happen because
					some objects may be sampled as faces. Facebox will reject images without
					a face with a 400 Bad Request error
				*/
			}
			log.Printf("Trained facebox with file: %s", f.Name())
		}
	}
	return nil
}
