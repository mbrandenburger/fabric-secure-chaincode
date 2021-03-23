# Setup

## Before 

## Run out-of-cluster

```bash
go run .
```

## Run in-cluster

TODO

Build the experimenter client
```bash
GOOS=linux go build -o ./app .
```

```bash
eval $(minikube docker-env)
docker build -t fpc/experimenter-client .
```

Allow the 
```bash
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=hyperledger:default
```

Create experimenter service
```bash
kubectl apply -f org2/org2-experimenter-deployment.yaml
```

Make service available at local host; use port as returned by this command
```bash
minikube service experimenter-client-service --url --namespace=hyperledger
```

Restart pod
```bash
kubectl -n hyperledger rollout restart deployment experimenter-client
```