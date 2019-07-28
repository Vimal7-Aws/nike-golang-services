init:
	echo "Init"

build-docker-image:
	docker build -t go-platform-service -f Dockerfile .

push:
	docker tag go-platform-service gcr.io/golden-tenure-231009/go-platform-service:1.0
	docker push gcr.io/golden-tenure-231009/go-platform-service:1.0

all: init  build-docker-image push
