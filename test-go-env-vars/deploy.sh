#!/usr/bin/env bash

kubectl apply -f infra/k8s/deployment.yaml
kubectl apply -f infra/k8s/service.yaml
