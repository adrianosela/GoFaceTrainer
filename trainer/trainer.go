package trainer

import (
	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
)

// FaceTrainer is an abstraction around the necessary OpenCV and MachineBox
// resources in order to produce a model of a face or facial-expression
type FaceTrainer struct {
	facebox       *facebox.Client
	captureDevice *gocv.VideoCapture
	window        *gocv.Window
	baseImgMatrix gocv.Mat
	faceAlgo      string
}

// NewFaceTrainer is the FaceTrainer constructor
func NewFaceTrainer( /* TODO */ ) (*FaceTrainer, error) {
	return &FaceTrainer{
		// TODO
	}, nil
}
