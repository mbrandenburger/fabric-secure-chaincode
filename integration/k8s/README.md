# Fabric Private Chaincode goes K8s

Many steps of this tutorial can also be invoked by using `just`. See official [installation](https://github.com/casey/just#installation) documentation. For Mac install via `brew install just`.

## TODOS

- get this network running with minikube
  [x] sim mode
  [] hw
- run this on a k8s cluster
  [x] minikube
  [] real cluster

## Prepare FPC deployment components

Let's first build the FPC components.
```bash
cd $FPC_PATH
make build
```

Next create the FPC Enclave Registry docker image
```bash
cd $FPC_PATH/ercc
make docker
```

### Write your FPC chaincode

For this demo we use the FPC Hello World chaincode as introduces in `$FPC_PATH/examples`.
Once you have completed the tutorial you should have the enclave binary available in `$FPC_PATH/examples/helloworld/_build/lib`.

Now we package the enclave binary in a FPC Chaincode docker image using the following command:

```bash
# TODO SETUP and test with SGX HW mode
cd ${FPC_PATH}/examples/helloworld
make

export CC_NAME=helloworld
make -C ${FPC_PATH}/ecc DOCKER_IMAGE=fpc/fpccc DOCKER_ENCLAVE_SO_PATH=${FPC_PATH}/examples/helloworld/_build/lib all docker
```

### Writing your FPC Client App

In order to communicate with the FPC hello world chaincode we use an app written in Go using our FPC Client SDK.

#### Build the client app

```bash
cd $FPC_PATH/client_sdk/go/sample/helloworld/
make build docker
```

## Minikube on Mac

This tutorial is heavily inspired by the article [How to implement Hyperledger Fabric External Chaincodes within a Kubernetes cluster](https://medium.com/swlh/how-to-implement-hyperledger-fabric-external-chaincodes-within-a-kubernetes-cluster-fd01d7544523) by Pau Aragonès Sabaté. Thank you!

### Start minikube

First start minikube and mount FPC files.
```bash
cd $FPC_PATH
minikube start --mount --mount-string="$(pwd):/fpc"

# Let's start he k8s dashboard
minikube dashboard
```

#### Push FPC images to minikube registry

First let's go the demo folder

```bash
cd $FPC_PATH/integration/k8s/demo
```

In order to use our docker images for the Enclave Registry, the FPC chaincode, and the FPC Client app, we need to
make these images available inside the k8s environment.
With minikube we can easily do that by calling the following commands:

```bash
minikube cache add fpc/ercc:latest
minikube cache add fpc/fpccc:latest
minikube cache add fpc/fpcclient:latest
```

You can double check if the images are available. Please also carefully check if the image ID corresponds to the image you want to use.
```bash
minikube ssh
docker images | grep fpc
```

When you update docker images, you need to tell minikube to use the updated images.
```bash
minikube cache reload
```

### Prepare K8s network

Prepare crypto material for the network
```bash
./generate.sh
````

Prepare k8s environment:
```bash
kubectl create ns hyperledger
kubectl create configmap peer-config --from-file=core.yaml -n hyperledger
kubectl create configmap chaincode-config --from-env-file=packages/chaincode-config.properties -n hyperledger
```

Deploy the Fabric network
```bash
kubectl create -f orderer-service/
kubectl create -f org1/
kubectl create -f org2/
kubectl create -f org3/
```

Alternatively, you can also call `just generate up`.

### Setting up our channel

In the next steps we need three terminals, each dedicated for a particular organization, to enter the CLI container.

Show all pods
```bash
# list all pods
kubectl get pods -n hyperledger
```

Org1
```bash
# copy cli_org1 pod name from above and enter container
kubectl exec -it cli_org1_pod_name -n hyperledger -- bash

peer channel create -o orderer0:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls true --cafile $ORDERER_CA
peer channel join -b mychannel.block
```

Org2 and Org3
```bash
# copy cli_org1 pod name from above and enter container
kubectl exec -it cli_orgX_pod_name -n hyperledger -- bash

peer channel fetch 0 mychannel.block -c mychannel -o orderer0:7050 --tls --cafile $ORDERER_CA
peer channel join -b mychannel.block
peer channel list
```

Alternatively, you can use `just run cli-org1` to login into the cli instance of org1.

#### Set anchor peers

In order to allow the peers to discover peers of another organization we need to setup an anchor peer for each org.
Since we only have a single peer per org, we set each peer as an anchor peer for their org.

As this task is somewhat tedious we provide a little helper script.

Execute the following command for each org in corresponding CLI container.
```bash
./scripts/setAnchorPeer.sh
```

For more details on anchor peer setup and channel configuration updates please see the official Hyperledger Fabric
[documentation](https://hyperledger-fabric.readthedocs.io/en/latest/config_update.html). 

### Install FPC Enclave Registry and FPC ECC

Now it is time to install FPC components on our network.
In particular, we need to install the FPC Enclave Registry (ERCC) and the actual chaincode.
With FPC, there are two options available to run these components based on the external launcher feature of Fabric.
We can run these components as normal Chaincode without docker or as chaincode as a Server (External chaincode).

#### Normal Chaincode without docker
The first option requires our custom FPC peer images, which provide the necessary software package to run Intel SGX applications. 

#### Chaincode as a Server (External chaincode)
The second option allows us to use standard Fabric peer images but requires the deployer to setup and run the chaincodes as a server.
Thus, the responsibility to maintain the chaincodes is in the hand of the network operator.

In this tutorial we show how to setup our FPC components using the Chaincode as a Server method.
Note that the `core.yaml` we have already loaded in the configmap contains the necessary configuration details to use the external builder scripts provided in the FPC repository. 

#### Install Enclave Registry and FPC Chaincode

For each Org:
```bash
# copy cli_orgX pod name from above and enter container
kubectl exec -it cli_orgX_pod_name bash -n hyperledge

# for all peers of orgX we install ercc and fpccc
export ERCC=ercc-peer0-$ORG; export FPCCC=fpccc-peer0-$ORG

peer lifecycle chaincode install packages/$ERCC.tgz
peer lifecycle chaincode install packages/$FPCCC.tgz
peer lifecycle chaincode queryinstalled

export ERCC_PKG_ID=$(peer lifecycle chaincode queryinstalled | grep ercc | awk '{print $3}' | sed 's/.$//')
export FPCCC_PKG_ID=$(peer lifecycle chaincode queryinstalled | grep fpccc | awk '{print $3}' | sed 's/.$//')

echo "$ERCC=$ERCC_PKG_ID" >> packages/chaincode-config.properties
echo "$FPCCC=$FPCCC_PKG_ID" >> packages/chaincode-config.properties
```

Alternatively, you can also run `./scripts/installCC.sh` inside the cli_orgX_pod_name.

#### Start chaincode container

Now, we have installed our FPC Enclave Registry and the FPC Chaincode on each peer. Since we are using the chaincode
as s server deployment model, it is now time to start the chaincodes.

Run the following commands from your host terminal.

```bash
kubectl create configmap chaincode-config --from-env-file=packages/chaincode-config.properties -n hyperledger --dry-run=client -o yaml | kubectl apply -f -
kubectl create -f chaincode/ercc/
kubectl create -f chaincode/fpccc/
```

Alternatively, you can also call `just chaincode`.

#### Approve

Next, we complete chaincode installation process by approving and committing the chaincode definitions.

For each Org:
```bash
# copy cli_orgX pod name from above and enter container
kubectl exec -it cli_orgX_pod_name bash -n hyperledge

# make sure you still have ERCC_PKG_ID, FPCCC_PKG_ID, and FPC_MRENCLAVE defined in your env (see above)

# approve enclave registry
peer lifecycle chaincode approveformyorg --channelID mychannel --name ercc --version 1.0 --package-id $ERCC_PKG_ID --sequence 1 -o orderer0:7050 --tls --cafile $ORDERER_CA

# approve FPC chaincode
peer lifecycle chaincode approveformyorg --channelID mychannel --name fpccc --version $FPC_MRENCLAVE --package-id $FPCCC_PKG_ID --sequence 1 -o orderer0:7050 --tls --cafile $ORDERER_CA
```

Alternatively, you can also run `./scripts/approve.sh` inside the cli_orgX_pod_name.

#### Commit

Pick some org that proceeds with committing chaincode.

```bash
# commit enclave registry
peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name ercc --version 1.0 --sequence 1 -o orderer0:7050 --tls --cafile $ORDERER_CA
peer lifecycle chaincode commit -o orderer0:7050 --channelID mychannel --name ercc --version 1.0 --sequence 1 --tls true --cafile $ORDERER_CA --peerAddresses peer0-org1:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/ca.crt --peerAddresses peer0-org2:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2/peers/peer0-org2/tls/ca.crt

# commit FPC chaincode
peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name fpccc --version  $FPC_MRENCLAVE --sequence 1 -o orderer0:7050 --tls --cafile $ORDERER_CA
peer lifecycle chaincode commit -o orderer0:7050 --channelID mychannel --name fpccc --version  $FPC_MRENCLAVE --sequence 1 --tls true --cafile $ORDERER_CA --peerAddresses peer0-org1:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/ca.crt --peerAddresses peer0-org2:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2/peers/peer0-org2/tls/ca.crt
```

### Run our hello world example

```bash
# and join into the container
kubectl exec -it fpcclient_orgX_pod_name bash -n hyperledge

# init our enclave
fpcclient init peer0-org1

# interact with the FPC chaincode
fpcclient store apple 100
fpcclient retrieve apple
```

Alternatively, you can also run `just run fpcclient-org1` to enter the container

### Shutdown cluster
```bash
kubectl delete --all all --namespace=hyperledger
```

Alternatively, you can also call `just down`.


## Troubleshooting

```bash
discover --configFile conf.yaml --peerTLSCA /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/ca.crt --tlsKey /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/server.key --tlsCert /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/server.crt --userKey /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/users/User1@org1/msp/keystore/priv_sk --userCert /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/users/User1@org1/msp/signcerts/User1@org1-cert.pem --MSP org1MSP saveConfig
discover --configFile conf.yaml peers --channel mychannel --server peer0-org1:7051
```
