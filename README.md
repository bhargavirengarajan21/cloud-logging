# Cloud Logging

Cloud Logging is a Kubernetes-native serverless logging solution built on top of OpenFaaS and NATS. This project provides a real-time, event-driven logging service that allows applications to publish and subscribe to log events via a simple HTTP interface.

## Features

- **Serverless Logging**: Uses OpenFaaS to deploy and manage the logging handler as a serverless function.
- **Real-Time Event Streaming**: Employs NATS as a lightweight message broker for real-time log event distribution.
- **HTTP API**: Accepts log events via HTTP POST requests, which are then published to a NATS topic.
- **Kubernetes Native**: All components are designed to run on Kubernetes, with easy deployment scripts provided.
- **Secure Secrets Management**: Uses Kubernetes secrets and config maps for managing sensitive information and configuration.

## Architecture

1. **OpenFaaS Function**: Handles HTTP requests and publishes log events to NATS.
2. **NATS**: Acts as the message broker for distributing log events.
3. **Kubernetes**: Orchestrates OpenFaaS, NATS, and network resources.

## Getting Started

### Prerequisites

- Kubernetes cluster
- `kubectl` configured to access your cluster
- [OpenFaaS CLI](https://docs.openfaas.com/cli/install/)
- [faas-cli](https://github.com/openfaas/faas-cli)
- [NATS](https://nats.io/)

### Deployment

1. **Deploy OpenFaaS and NATS**

   Use the provided deployment script to deploy OpenFaaS components and NATS to your cluster:

   ```bash
   ./openfass-deploy.sh
   ```

2. **Configure Secrets and NATS**

   The deployment script will automatically run `secrets.sh` to create the necessary Kubernetes secrets and NATS configuration.

3. **Deploy Logging Function**

   The function will be deployed via OpenFaaS using the provided YAML definition.

4. **Access OpenFaaS Gateway**

   The script will output the external IP of your OpenFaaS gateway. Access it at:

   ```
   http://<EXTERNAL_IP>:8080
   ```

### Usage

- **Publish a Log Event:**

  Send a POST request to the function endpoint:

  ```bash
  curl -X POST http://<GATEWAY_IP>:8080/function/log-pubsub-handler \
    -H "Content-Type: application/json" \
    -d '{"message": "Your log message"}'
  ```

- **Subscribe to Log Events:**

  The service subscribes to the `log-events` topic on NATS, printing all received messages.

## Project Structure

- `log-pubsub-handler/handler.go` - Core function for handling HTTP requests and publishing to NATS.
- `openfass-deploy.sh` - Main deployment script for OpenFaaS and the logging function.
- `secrets.sh` - Script to set up Kubernetes secrets and NATS config.
- `pubsub-function.yaml`, `openfaas-deployment.yaml` - Kubernetes and OpenFaaS deployment manifests.

## Example

```json
POST /function/log-pubsub-handler
{
  "message": "This is a test log event"
}
```
Logs will be published to the NATS topic and available to subscribers in real time.

## License

This project is open source. See the LICENSE file for details.

## Author

[bhargavirengarajan21](https://github.com/bhargavirengarajan21)
