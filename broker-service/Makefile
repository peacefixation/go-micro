DOCKER_IMAGE_NAME=peacefixation/broker-service
DOCKER_IMAGE_VERSION=1.0.1

# generate protocol buffers source
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs/logs.proto

# build docker image and tag it
docker-tag:
	docker build -f broker-service.dockerfile -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} .

# requires "docker login" first
# enter username and password for hub.docker.com
docker-push:
	docker push ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}
