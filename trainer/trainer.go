package trainer

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/machinebox/sdk-go/facebox"
)

const (
	defaultFaceboxAddr = "http://localhost:8080"
	defaultSrcFacesDir = "."
)

// FaceTrainer is an abstraction around the necessary Machinebox
// resources in order to train a model of a face or facial-expression
type FaceTrainer struct {
	faceboxClient *facebox.Client
	facesDir      string
}

// Settings contains necessary values to initialize a FaceTrainer
type Settings struct {
	// FaceboxAddress is the URL of the MachineBox to use - host:port,
	// (default is http://localhost:8080)
	FaceboxAddress string

	// FacesSrcDir is the directory where .jpg samples will be sources from
	// (default is the current directory ".")
	FacesSrcDir string
}

// NewFaceTrainer is the FaceTrainer constructor
func NewFaceTrainer(settings Settings) (*FaceTrainer, error) {
	if settings.FaceboxAddress == "" {
		settings.FaceboxAddress = defaultFaceboxAddr
	}
	if settings.FacesSrcDir == "" {
		settings.FacesSrcDir = defaultSrcFacesDir
	}
	return &FaceTrainer{
		faceboxClient: facebox.New(settings.FaceboxAddress),
		facesDir:      settings.FacesSrcDir,
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
			// TODO
			fmt.Println(f.Name())
		}
	}
	return nil
}
