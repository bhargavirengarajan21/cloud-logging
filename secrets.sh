#!/bin/bash

set -e

echo "Creating basic-auth secrets..."

kubectl delete secret basic-auth -n openfaas || true
kubectl create secret generic basic-auth \
  --from-literal=basic-auth-user=admin \
  --from-literal=basic-auth-password=admin \
  -n openfaas

echo "Basic-auth secret created!"

echo "Creating NATS config..."

# Delete existing config map (if any)
kubectl delete configmap nats-config -n openfaas || true

# Create new NATS config map
kubectl create configmap nats-config \
  --from-literal=nats_url="nats://nats.openfaas.svc.cluster.local:4222" \
  -n openfaas

echo "NATS config created!"

echo "Secrets and ConfigMap applied successfully!"
