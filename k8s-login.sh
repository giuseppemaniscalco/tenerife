#!/gin/bash

kubectl config set-cluster ws-${USER} \
    --embed-certs=true \
    --server=https://35.205.229.158 \
    --certificate-authority=${CA_PATH}
kubectl config set-credentials ws-${USER} --token=${TOKEN}
kubectl config set-context ws-${USER} \
    --cluster=ws-${USER} \
    --user=ws-${USER} \
    --namespace=${USER}
kubectl config use-context ws-${USER}