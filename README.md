# school-manager
school-manager is a collection of Kubernetes based services and operators to manage a school database system using solely Kubernetes objects without a database.

## Running Locally
make sure you have minikube installed and run
```
./start-dev.sh
```
## Development
Run the following before building locally: 
```
export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
```
