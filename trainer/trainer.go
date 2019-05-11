package trainer

import (
	"fmt"

	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
)

const (
	defaultSamples = 10
	escapeKey      = 27
)

// FaceTrainer is an abstraction around the necessary OpenCV and MachineBox
// resources in order to produce a model of a face or facial-expression
type FaceTrainer struct {
	captureDevice  *gocv.VideoCapture
	window         *gocv.Window
	baseImgMatrix  gocv.Mat
	faceClassifier gocv.CascadeClassifier
	faceboxClient  *facebox.Client
	samples        int
}

// Settings contains necessary values to initialize a FaceTrainer
type Settings struct {
	// CaptureDeviceID (default is 0)
	CaptureDeviceID int

	// FaceboxAddress is the URL of the MachineBox to use - host:port,
	// (default is http://localhost:8080)
	FaceboxAddress string

	// FaceAlgoPath is the path of the classifier algorithm XML file to use
	// (no default, this is a required argument)
	FaceAlgoPath string

	// WindowTitle is the title to be displayed on the window
	// (default is "Face Trainer")
	WindowTitle string
}

// NewFaceTrainer is the FaceTrainer constructor
func NewFaceTrainer(s *Settings) (*FaceTrainer, error) {
	device, err := gocv.VideoCaptureDevice(s.CaptureDeviceID)
	if err != nil {
		return nil, fmt.Errorf("error opening capture device: %s", err)
	}
	if s.FaceboxAddress == "" {
		s.FaceboxAddress = "http://localhost:8080"
	}
	if s.WindowTitle == "" {
		s.WindowTitle = "Face Trainer"
	}
	if s.FaceAlgoPath == "" {
		return nil, fmt.Errorf("no face recognition algorithm provided")
	}
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(s.FaceAlgoPath)
	return &FaceTrainer{
		captureDevice:  device,
		window:         gocv.NewWindow(s.WindowTitle),
		baseImgMatrix:  gocv.NewMat(),
		faceClassifier: classifier,
		faceboxClient:  facebox.New(s.FaceboxAddress),
		samples:        defaultSamples,
	}, nil
}

// Close closes all closers within a FaceTrainer, should be used in a defer
// statement immediately after checking the error from NewFaceTrainer
func (t *FaceTrainer) Close() error {
	errs := []error{}
	if err := t.faceClassifier.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close classifier: %s", err))
	}
	if err := t.window.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close window: %s", err))
	}
	if err := t.captureDevice.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close capture device: %s", err))
	}
	if len(errs) != 0 {
		errsStr := ""
		for _, err := range errs {
			errsStr = fmt.Sprintf("%s | %s", errsStr, err)
		}
		return fmt.Errorf("%s: %s", "Close() encountered some errors", errsStr)
	}
	return nil
}

// SampleFace samples a face t.samples times, and saves each sample to a file
func (t *FaceTrainer) SampleFace() error {
	for {
		if ok := t.captureDevice.Read(&t.baseImgMatrix); !ok || t.baseImgMatrix.Empty() {
			continue
		}

		rects := t.faceClassifier.DetectMultiScale(t.baseImgMatrix)
		for _, r := range rects {
			imgFace := t.baseImgMatrix.Region(r)
			imgName := fmt.Sprintf("%d.jpg", t.samples)
			gocv.IMWrite(imgName, imgFace)
			imgFace.Close()
			if t.samples--; t.samples == 0 {
				return nil
			}
		}

		t.window.IMShow(t.baseImgMatrix)
		if t.window.WaitKey(100) == escapeKey {
			return nil
		}
	}
}
