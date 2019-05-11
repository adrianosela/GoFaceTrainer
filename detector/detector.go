package detector

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
)

const (
	defaultWindowTitle = "Face Detector"
	defaultFaceboxAddr = "http://localhost:8080"

	escapeKey = 27
)

// FaceDetector uses OpenCV and Machinebox/Facebox to detect and compare
// faces and/or facial expressions
type FaceDetector struct {
	captureDevice  *gocv.VideoCapture
	window         *gocv.Window
	baseImgMatrix  gocv.Mat
	faceboxClient  *facebox.Client
	faceClassifier gocv.CascadeClassifier
}

// Settings contains necessary values to initialize a FaceDetector
type Settings struct {
	// CaptureDeviceID (default is 0)
	CaptureDeviceID int

	// FaceAlgoPath is the path of the classifier algorithm XML file to use
	// (no default, this is a required argument)
	FaceAlgoPath string

	// WindowTitle is the title to be displayed on the window
	// (default is "Face Trainer")
	WindowTitle string

	// FaceboxAddress is the URL of the MachineBox to use - host:port,
	// (default is http://localhost:8080)
	FaceboxAddress string
}

// NewFaceDetector is the FaceDetector constructor
func NewFaceDetector(settings Settings) (*FaceDetector, error) {
	device, err := gocv.VideoCaptureDevice(settings.CaptureDeviceID)
	if err != nil {
		return nil, fmt.Errorf("error opening capture device: %s", err)
	}
	if settings.FaceboxAddress == "" {
		settings.FaceboxAddress = defaultFaceboxAddr
	}
	if settings.WindowTitle == "" {
		settings.WindowTitle = defaultWindowTitle
	}
	if settings.FaceAlgoPath == "" {
		return nil, fmt.Errorf("no face recognition algorithm provided")
	}
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(settings.FaceAlgoPath)
	return &FaceDetector{
		faceboxClient:  facebox.New(settings.FaceboxAddress),
		captureDevice:  device,
		window:         gocv.NewWindow(settings.WindowTitle),
		baseImgMatrix:  gocv.NewMat(),
		faceClassifier: classifier,
	}, nil
}

// Close closes all closers within a FaceDetector, should be used in a defer
// statement immediately after checking the error from NewFaceDetector
func (d *FaceDetector) Close() error {
	errs := []error{}
	if err := d.faceClassifier.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close classifier: %s", err))
	}
	if err := d.window.Close(); err != nil {
		errs = append(errs, fmt.Errorf("could not close window: %s", err))
	}
	if err := d.captureDevice.Close(); err != nil {
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
func (d *FaceDetector) Run() error {
	for {
		if ok := d.captureDevice.Read(&d.baseImgMatrix); !ok || d.baseImgMatrix.Empty() {
			log.Print("cannot read webcam")
			continue
		}

		rects := d.faceClassifier.DetectMultiScale(d.baseImgMatrix)
		for _, r := range rects {
			faceImg := d.baseImgMatrix.Region(r)
			buf, err := gocv.IMEncode(".jpg", faceImg)
			faceImg.Close()

			if err != nil {
				log.Printf("could not encode face image to matrix: %s", err)
				continue // move on to next face
			}

			faces, err := d.faceboxClient.Check(bytes.NewReader(buf))
			if err != nil {
				log.Printf("unable to recognize face: %v", err)
			}

			var caption = "UNKNOWN"
			if len(faces) > 0 {
				caption = faces[0].Name
			}

			size := gocv.GetTextSize(caption, gocv.FontHersheyPlain, 3, 2)
			point := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&d.baseImgMatrix, caption, point, gocv.FontHersheyPlain, 3, color.RGBA{0, 0, 255, 0}, 2)
			gocv.Rectangle(&d.baseImgMatrix, r, color.RGBA{0, 0, 255, 0}, 3)
		}

		d.window.IMShow(d.baseImgMatrix)
		if d.window.WaitKey(5) == 27 {
			return nil
		}
	}
}
