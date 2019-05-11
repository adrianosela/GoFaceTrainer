sample:
	mkdir out
	go run cmd/cmd.go sample 0 face_algos/haarcascade_frontalface_default.xml 10 ./out

train:
	go run cmd/cmd.go train http://localhost:8080 ./out YOURNAME

run:
	go run cmd.go run 0 http://localhost:8080 face_algos/haarcascade_frontalface_default.xml

clean:
	rm -rf out
