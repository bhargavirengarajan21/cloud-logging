#!/bin/bash

set -e

echo "Creating namespaces..."
kubectl apply -f openfaas-deployment.yaml

echo "Waiting for OpenFaaS to be ready..."
kubectl rollout status -n openfaas deploy/gateway --timeout=60s


# Create Secrets and ConfigMap
./secrets.sh

# Login to OpenFaaS
PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode)
echo $PASSWORD | faas-cli login -g http://127.0.0.1:8080 -u admin --password-stdin

# Deploy function
echo "Deploying function..."
faas-cli deploy -f pubsub-function.yaml

# Get External IP
EXTERNAL_IP=$(kubectl get svc -n openfaas gateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
echo "OpenFaaS Gateway available at: http://$EXTERNAL_IP:8080"

# Test function
echo "Testing function..."
curl -v http://127.0.0.1:8080/function/log-pubsub-handler
