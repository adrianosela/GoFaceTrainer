## GoFaceTrainer - Facial Recognition Model in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/adrianosela/gofacetrainer)](https://goreportcard.com/report/github.com/adrianosela/gofacetrainer)
[![Documentation](https://godoc.org/github.com/adrianosela/GoFaceTrainer?status.svg)](https://godoc.org/github.com/adrianosela/GoFaceTrainer)
[![GitHub issues](https://img.shields.io/github/issues/adrianosela/GoFaceTrainer.svg)](https://github.com/adrianosela/GoFaceTrainer/issues)
[![license](https://img.shields.io/github/license/adrianosela/gofacetrainer.svg)](https://github.com/adrianosela/GoFaceTrainer/blob/master/LICENSE)

### Step-By-Step Tutorial: 

* **STEP 0 - Prerequisites:**
	* Download OpenCV for your OS: on mac ```brew install opencv3```
	* Download Docker for your OS
	* Download the Go Programming Language (Golang)
	* Clone this repository:
```git clone https://github.com/adrianosela/GoFaceTrainer```

* **STEP 1 - Getting a Machine Box API key:** 
	* Head over to ```https://machinebox.io/account``` and create a Veritone account; follow the account creation procedure (confirmation email, web app onboarding, etc)
	* On the web application, click on "Machine Box" on the left side dropdown
	* Scroll down to "Your key" and click on "Reveal your key"
	* Copy and save this API key somewhere safe; you will use it to authenticate against a local Machine Box API which we will deploy later on in this tutorial. I exported it on my command line: ```export MB_KEY=[the key]```

![](./tutorial_assets/step1.png)

* **STEP 2 - Running Machine Box Locally:**
	* Pull the [machinebox/facebox](https://hub.docker.com/r/machinebox/facebox) Docker image: ```docker pull machinebox/facebox```
	* Run the Docker image, mapping port 8080 on the container to your desired port on your machine (I have chosen port 8080 for my machine port as well), passing your Machine Box API key from step 1 as the ```MB_KEY``` environment variable (example command below). Note that you can override the container's serving port (default 8080) with the ```MB_PORT``` environment variable
	* On your browser, head over to http://localhost:[mapped-port] and verify you can see the Machine Box console/UI

```
$ docker run -p 8080:8080 -e "MB_KEY=$MB_KEY" machinebox/facebox
[INFO]     starting...

	Welcome to Facebox by Machine Box
	(facebox 3d8ecd86)

	Visit the console to see what this box can do:
	http://localhost:8080

	If you have any questions or feedback, get in touch:
	https://machinebox.io/contact

	Report bugs and issues:
	https://github.com/machinebox/issues

	Tell us what you build on Twitter @machineboxio

[INFO]     box ready
```

![](./tutorial_assets/step2.png)

* **STEP 3 - Sample Your Face:**
	* Run the Go program in this directory with arguments as below. Note that you must create the output directory before running the command if it does not exist

```
go run main.go sample [CAMERA_ID] [FACE_ALGO] [N_SAMPLES] [OUTPUT_DIR]
```

Note before the example line below I ran ```mkdir out``` to create the output directory I'm using

> Example:```go run main.go sample 0 face_algos/haarcascade_frontalface_default.xml 10 out
```

* **STEP 4 - Train Your Facebox Instance With Your Face Samples:**
	* Run the Go program in this directory with arguments as follows:

```
go run main.go train [FACEBOX_URL] [FACES_SRC_DIR] [PERSON_NAME/EXPRESSION_NAME]
``` 

> Example: ```go run main.go train http://localhost:8080 out Adriano```

* **STEP 5 - Repeat steps 3 and 4 for each person (or facial expression) you want the model to recognize**

* **STEP 6 - Test Your Model:**
	* Run the Go program in this directory with arguments as follows:

```
go run main.go run [CAMERA_ID] [FACEBOX_URL] [FACE_ALGO]
```

> Example:```go run main.go run 0 http://localhost:8080 face_algos/haarcascade_frontalface_default.xml```

---

You are done! hopefully your model is accurate enough. If not, you can always try adding more images by repeating steps 3 and 4

##### I hope you have enjoyed, issues and PRs are welcome!

### Credit given where credit is due:

This project was inspired by [packagemain's video](https://www.youtube.com/watch?v=rbZeZNVA-Q4): "Face Detection in Go using OpenCV and MachineBox". I highly recommend watching his videos and subscribing if you are a Go enthusiast!
