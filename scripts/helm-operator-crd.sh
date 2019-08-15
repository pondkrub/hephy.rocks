export REPO=https://raw.githubusercontent.com/fluxcd/flux/master

kubectl apply -f ${REPO}/deploy-helm/flux-helm-release-crd.yaml
# kubectl apply -f ${REPO}/deploy-helm/weave-cloud-helm-operator-deployment.yaml
