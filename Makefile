build:
	CGO_ENABLED=0 GOOS=linux go build -o main -a main.go
image: build
	docker build -t xiaowang1234567/test-controller:v1 .
push: image
	docker push xiaowang1234567/test-controller:v1
