# Go Micro

Udemy: Working with microservices in Go

See the project/Makefile for targets to build the suite of docker containers and start/stop the front end. The different orchestration methods are detailed at the end of this document.

The front end is a basic go application that serves a HTML page. Use the UI to send test requests to various microservices (authentication, log, mail, rabbitmq, etc).

The broker service handles all requests and routes them to the appropriate service.

The authentication service handles requests to athenticate a user against details stored in Postgres database.

The mail service uses MailHog to simulate sending an email.

The listener service listens for messages in RabbitMQ.

The logging service accepts requests to add logs via JSON, RPC and gRPC and saves logs to MongoDB.

The Postgres database is populated manually at this point (i.e. use a Postgres client and import the SQL in the included users.sql)

## Dependencies

### Protocol Buffers

Go tools
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

Protocol Buffer Compiler
- https://grpc.io/docs/protoc-installation/
- apt install -y protobuf-compiler

## Tools

### Postgres

Beekeeper Studio
- https://www.beekeeperstudio.io/

### MongoDB

MongoDB Compass
- https://www.mongodb.com/try/download/compass

### Kubernetes

Minikube
- https://minikube.sigs.k8s.io/docs/start/

Kubectl
- https://kubernetes.io/docs/tasks/tools/

## Orchestration

### Docker Compose

Build the suite of docker containers:
- `make up_build`
Hit the front end in a browser
- `http://localhost:8084`

### Docker Swarm

The docker swarm includes a Caddy reverse proxy to forward requests to containers in the swarm.

There is a production version of the deployment descriptor (swarm.production.yml) for deploying to cloud services (i.e. Linode).

Part of this exercise was deploying the swarm to a Linode host. This was done, but the nodes are since deleted to avoid ongoing costs. See project/LINODE.md for provisioning and deployment instructions.

Initialise the swarm:
- `make swarm-init`
Deploy the swarm:
- `make swarm-deploy`
Hit the front end in a browser:
- `http://localhost`
Stop the swarm (scale replicated containers to 0, remove the stack)]
- `make swarm-stop`
- `make swarm-rm`

### Kubernetes

Install `minikube` on the host machine. Install `kubectl` on the host. The kubernetes deployment includes an `nginx` ingress to forward requests to the pods. We simulate a remote postgres service by running it in a separate docker container.

Add hostnames to /etc/hosts:
- `127.0.0.1 front-end.info broker-service.info`

Start minikube:
- `make minikube-start`
View minikube dashboard:
- `make minikube-dashboard`
Deploy pods:
- `make kubectl-deploy`
Deploy ingress:
- `make kubectl-deploy-ingress`
Deploy (simulated remote) postgres:
- `make k8s-remote-postgres`
Run minikube tunnel (so we can reach the ingress):
- `minikube tunnel`
Hit the front end in a browser:
- `http://front-end.info`
