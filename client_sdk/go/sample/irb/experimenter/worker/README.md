# Graphene worker executing some data analytics provisioned with data with the help of FPC 

## TODOs

- Package this container with Graphene Shielded Containers
- Test attestation (see todos within `main.go`)
- Run with Graphene

## Build and run a worker (simulation mode - no Graphene)

Build worker webservice
```bash
GOOS=linux go build -o ./app .
```

Build the docker container (copying the webservice `app` and the python code from `/src`)
```bash
# do this only if you wanna run this code with k8s minikube
eval $(minikube docker-env)

docker build -t fpc/experimenter-worker .
```

### Graphene Shielded Containers

TODO follow https://graphene.readthedocs.io/en/latest/manpages/gsc.html

### Configuration

The webserivce is responsible to fetch the input data from redis. Therefore, you need
to specify the redis host, port, and password. When you start the container make sure
you set the following env vars to configure the redis client. Also make sure you can
reach the redis host from within the container.

```bash
REDIS_HOST="localhost"
REDIS_PORT=6379
REDIS_PASSWORD=""
```

Run locally, the webservice listens on port 5000.
```bash
docker run -p 5000:5000 -it fpc/experimenter-worker
```
