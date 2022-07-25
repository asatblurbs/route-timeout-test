# route-timeout-test

* minikube start  --memory=2200mb --kubernetes-version=v1.22.12 --force

* kubectl create deploy golang-saiful --image=docker.io/asatblurbs/golang-saiful:1.0 --port=8080 -n default
* kubectl expose deploy golang-saiful   --type=NodePort --port=8080 -n default 
* minikube service golang-saiful --url
