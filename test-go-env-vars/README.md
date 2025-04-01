Test KimMachineGun/automemlimit and go.uber.org/automaxprocs
============================================================

The program tests how the automemlimit and automaxprocs behave

## Build

`docker buildx build -t myapp .`

## Run

`docker run --rm --cpus="4" --memory="64m" myapp`

## Deploy

1. Select your k8s context, e.g. 
`kubectl config use-context docker-desktop`

2. Invoke deployment
`./deploy.sh`

3. View logs in pod
`kubectl logs -f $(kubectl get pods  | grep -i myapp | awk '{print $1}')`
