## Pods

- Smallest unit of K8s
- Group of one more containers
- Abstraction over container
- Each pod gets its own IP address
- New ip address on recreation (😖 headache!)

## Service

- Permanent ip address that can be attached on each pod
- Pods communicate with each other with service
- Lifecycle of pod and service NOT connected. Even if recreated the ip address stay same
- Also work as a load balancer🔨
    - Replica of a pod is created by defining blueprints for pods - Deployment
- Types:

    1. External Service
    - Opens the connection to external sources

    2. Internal Service
    - Only accessible internally. eg:- Datbase endpoint❗

## Ingress

- Make external service endpoint like https://my-app.com instead of http://123.89.8.1:8090

- Perform forwarding to service

## ConfigMap

- External configuration of the application
- Avoid re-building image after a minor change
- But putting credentials into configmap is not recommended🚫 - use Secret

## Secret

- Used to store secret data
- Stored in base64 encoded
- Can be use like configmap

## Volumes

- Support data persistence
- Attach a physical hard drive
- Storage on local or remote storage - not part of k8 cluster

## Deployment

- Blueprint for creating the pods
- `kubectl create deployment nginx-depl --image=nginx` : most basic configuration for deployment(only name and image to use😊)
- Abstraction of pods - create replica of pod

## StatefulSet

- For databases in load-balancing to avoid data inconsistency 