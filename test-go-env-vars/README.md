Test KimMachineGun/automemlimit and go.uber.org/automaxprocs
============================================================

The program tests how the automemlimit and automaxprocs behave

## Build

`docker buildx build -t myapp .`

## Run

`docker run --rm --cpus="4" --memory="64m" myapp`
