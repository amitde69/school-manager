# school-manager
school-manager is a collection of Kubernetes based services to manage a school database system.

## Running Locally
```
minikube start
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
## Development
Run the following before building locally: 
```
export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
```
