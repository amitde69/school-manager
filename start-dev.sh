#!/bin/bash
minikube start

## Deploy ArgoCD
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

## Change default ArgoCD admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
while [[ $? != 0 ]]; do
    kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
done
argopass=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
kubectl port-forward svc/argocd-server 9090:80 -n argocd &
argocd login localhost:9090  --username admin --password $argopass --insecure
argocd account update-password --current-password $argopass --new-password Aa123456

## Deploy Istio infra
argocd app create  -f istio/argocd/operator-app.yaml 
## wait for CRDs to be created before using them
sleep 10
argocd app create  -f istio/argocd/infra-app.yaml

## Deploy students domain
argocd app create  -f students/argocd/api-app.yaml
argocd app create  -f students/argocd/controller-app.yaml
## Deploy classes domain
argocd app create  -f classes/argocd/api-app.yaml
argocd app create  -f classes/argocd/controller-app.yaml

## expose istio-ingressgateway on port 8080
kubectl port-forward svc/istio-ingressgateway 8080:80 -n istio-system &


