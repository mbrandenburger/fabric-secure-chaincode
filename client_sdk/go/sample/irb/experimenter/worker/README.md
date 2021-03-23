Build worker
```bash
GOOS=linux go build -o ./app .
```

```bash
eval $(minikube docker-env)
docker build -t fpc/experimenter-worker .
```

docker run -p 5000:5000 -it fpc/experimenter-worker