# route-timeout-test

## Install docker

## Start minikube
* minikube start  --memory=2200mb --kubernetes-version=v1.22.12 --force

## Build go app
* git clone https://github.com/asatblurbs/route-timeout-test.git
* cd route-timeout-test
* docker login
* docker build -t docker.io/asatblurbs/golang-saiful:1.0 .
* docker push docker.io/asatblurbs/golang-saiful:1.0

## Deploy application
* kubectl create deploy golang-saiful --image=docker.io/asatblurbs/golang-saiful:1.0 --port=8080 -n default
* kubectl expose deploy golang-saiful   --type=NodePort --port=8080 -n default 
* minikube service golang-saiful --url
