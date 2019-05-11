package sampler

import (
	"fmt"

	"gocv.io/x/gocv"
)

const (
	defaultWindowTitle    = "Face Sampler"
	defaultSamples        = 10
	defaultSaveSamplesDir = "."

	escapeKey = 27
)

// FaceSampler is an abstraction around the necessary OpenCV
// resources in order to sample a model of a face or facial-expression
type FaceSampler struct {
	captureDevice  *gocv.VideoCapture
	window         *gocv.Window
	baseImgMatrix  gocv.Mat
	faceClassifier gocv.CascadeClassifier
	samples        int
	samplesSaveDir string
}

// Settings contains necessary values to initialize a FaceSampler
type Settings struct {
	// CaptureDeviceID (default is 0)
	CaptureDeviceID int

	// FaceAlgoPath is the path of the classifier algorithm XML file to use
	// (no default, this is a required argument)
	FaceAlgoPath string

	// WindowTitle is the title to be displayed on the window
	// (default is "Face Trainer")
	WindowTitle string

	// NSamples is how many samples to take in one run of the sampler
	// (default is 10)
	NSamples int

	// SaveSamplesDir is the directory where .jpg samples will be saved to
	// (default is the current directory ".")
	SaveSamplesDir string
}

// NewFaceSampler is the FaceSampler constructor
func NewFaceSampler(settings Settings) (*FaceSampler, error) {
	device, err := gocv.VideoCaptureDevice(settings.CaptureDeviceID)
	if err != nil {
		return nil, fmt.Errorf("error opening capture device: %s", err)
	}
	if settings.WindowTitle == "" {
		settings.WindowTitle = defaultWindowTitle
	}
	if settings.FaceAlgoPath == "" {
		return nil, fmt.Errorf("no face recognition algorithm provided")
	}
	if settings.NSamples == 0 {
		settings.NSamples = defaultSamples
	}
	if settings.SaveSamplesDir == "" {
		settings.SaveSamplesDir = defaultSaveSamplesDir
	}
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(settings.FaceAlgoPath)
	return &FaceSampler{
		captureDevice:  device,
		window:         gocv.NewWindow(settings.WindowTitle),
		baseImgMatrix:  gocv.NewMat(),
		faceClassifier: classifier,
		samples:        settings.NSamples,
		samplesSaveDir: settings.SaveSamplesDir,
	}, nil
}

// Close closes all closers within a FaceSampler, should be used in a defer
// statement immediately after checking the error from NewFaceSampler
func (s *FaceSampler) Close() error {
	errs := []error{}
	if err := s.faceClassifier.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close classifier: %s", err))
	}
	if err := s.window.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close window: %s", err))
	}
	if err := s.captureDevice.Close(); err != nil {
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

// Run samples a face NSamples times, and saves each sample to a file
func (s *FaceSampler) Run() error {
	for {
		if ok := s.captureDevice.Read(&s.baseImgMatrix); !ok || s.baseImgMatrix.Empty() {
			continue
		}

		rects := s.faceClassifier.DetectMultiScale(s.baseImgMatrix)
		for _, r := range rects {
			imgFace := s.baseImgMatrix.Region(r)
			imgName := fmt.Sprintf("%s/%d.jpg", s.samplesSaveDir, s.samples)
			gocv.IMWrite(imgName, imgFace)
			imgFace.Close()
			if s.samples--; s.samples == 0 {
				return nil
			}
		}

		s.window.IMShow(s.baseImgMatrix)
		if s.window.WaitKey(10) == escapeKey {
			return nil
		}
	}
}
