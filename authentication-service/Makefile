
# build docker image and tag it
docker-tag:
	docker build -f authentication-service.dockerfile -t peacefixation/authentication-service:1.0.0 .

# requires "docker login" first
# enter username and password for hub.docker.com
docker-push:
	docker push peacefixation/authentication-service:1.0.0

