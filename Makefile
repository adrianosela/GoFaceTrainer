sample:
	mkdir out
	go run main.go sample 0 face_algos/haarcascade_frontalface_default.xml 10 ./out

train:
	go run main.go train http://localhost:8080 ./out YOURNAME

run:
	go run main.go run 0 http://localhost:8080 face_algos/haarcascade_frontalface_default.xml

clean:
	rm -rf out
