FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerServiceApp
MAIL_BINARY=mailerApp
LISTENER_BINARY=listenerApp
FRONT_BINARY=frontEndApp
STACK=goMicro

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_logger build_mail build_listener
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_front_linux: builds the broker binary as a linux executable
build_front_linux:
	@echo "Building front end linux binary..."
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_BINARY} ./cmd/web
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} .
	@echo "Done!"

docker-tag-caddy:
	docker build -f caddy.dockerfile -t peacefixation/gomicro-caddy:1.0.0 .

docker-push-caddy:
	docker push peacefixation/gomicro-caddy:1.0.0

docker-tag-caddy-prod:
	docker build -f caddy.production.dockerfile -t peacefixation/gomicro-caddy-production:1.0.1 .

docker-push-caddy-prod:
	docker push peacefixation/gomicro-caddy-production:1.0.1

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

# initialise the docker swarm
swarm-init:
	docker swarm init

swarm-join-token-worker:
	docker swarm join-token worker

swarm-join-token-manager:
	docker swarm join-token manager

# deploy the docker swarm
swarm-deploy:
	docker stack deploy -c swarm.yml ${STACK}

# stop the docker swarm
# we scale any mode: replicated containers to 0
# don't try to scale mode: global nodes (they don't scale)
swarm-stop:
	docker service scale ${STACK}_authentication-service=0
	docker service scale ${STACK}_broker-service=0
	docker service scale ${STACK}_listener-service=0
	docker service scale ${STACK}_logger-service=0
	docker service scale ${STACK}_mail-service=0
	docker service scale ${STACK}_postgres=0

swarm-rm:
	docker stack rm ${STACK}

swarm-leave:
	docker swarm leave --force

# update a running service
# docker service update --image peacefixation/logger-service:1.0.1 goMicro_logger-service

minikube-start:
	minikube start

minikube-stop:
	minikube stop

minikube-dashboard:
	minikube dashboard

kubectl-deploy:
	kubectl apply -f k8s

# run add "127.0.0.1 front-end.info broker-service.info" to /etc/hosts
# run "minikube tunnel" after the ingress is up
kubectl-deploy-ingress:
	kubectl apply -f ingress.yml

kubectl-undeploy:
	kubectl delete deployments broker-service mongo rabbitmq mailhog mail logger listener authentication front-end
	kubectl delete svc broker-service mongo rabbitmq mailhog mail logger listener authentication front-end

kubectl-pods:
	kubectl get pods

kubectl-deployments:
	kubectl get deployments

kubectl-services:
	kubectl get services

# kubectl logs <pod name>

# simulate connecting to an external postgres instance by spinning up postgres in a separate docker container
k8s-remote-postgres:
	docker-compose -f postgres.yml up -d
