# generate protocol buffers source
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs/logs.proto

# build docker image and tag it
docker-tag:
	docker build -f logger-service.dockerfile -t peacefixation/logger-service:1.0.0 .

# requires "docker login" first
# enter username and password for hub.docker.com
docker-push:
	docker push peacefixation/logger-service:1.0.0
