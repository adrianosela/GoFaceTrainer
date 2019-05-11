sample:
	mkdir out
	go run cmd/sample.go 0 face_algos/haarcascade_frontalface_default.xml 10 ./out

clean:
	rm -rf out
