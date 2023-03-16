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

## Namespace

- Organise resources in namespace
- 4 namespace per default
    1. kube-system
        - system process

    2. kube-public
        - Publically accessible data
        - Configmap contains which contain cluster information

    3. kube-node-lease
        - Determine the pod availability

    4. default
        - resources we create are located here if we do not create any namespace

`kubectl create namespace my-namespace` : create a namespace

`kubectl get namespace` : list all the namespace

`kubectl get all -n my-namespace` : list all in a particular namespace

`kubectl apply -f <yaml file> --namespace=my-namespace` : create component in a particular namespace

or specifying in the configuration file itself👍

Use <b>kubens</b> to change the default active namespace to our own custom one



```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mongodb-configmap
  namespace: my-namespace
data:
  database_url: mongodb-service
```

- If you create all the resources in a default namespace, it is hard to get a view about it. Create namespace with identical resources together inside a cluster
- Separate namespace for separate team to avoid conflict
- Separate namespace on different version
- Access and resource limit

## Ingress

- Make external service endpoint like https://my-app.com instead of http://123.89.8.1:8090

- Perform forwarding to internalservice

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dashboard-ingress
  namespace: kubernetes-dashboard
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
  - host: dashboard.com # the endpoint url
    http:               # protocol used
      paths:
      - path: /
        pathType: Exact  
        backend:    # target incoming req is redirected
          service:
            name: kubernetes-dashboard  # internal service name
            port: 
              number: 80
```

- Ingress Controller : manage and process ingress rules. Also manages redirection.

`minikube addons enable ingress` : install ingress controller

- Ingress can be used to subdomain and multiple paths for same host 🛴

## Helm - Package Manager

- Package manager like apt, brew
- Avoid manually configuring yaml files
- Someone created the configuration file, package it and made available in somewhere so other people can be used
- Helm Charts :
    - Bundle of yaml files
    - Create helm charts and can be pushed them to helm repository
    - Can also download and use existing ones
    - Can be reuse the configuration someone used

`helm search <keyword>` : search for a helm chart
- Template engine
    - Define a common blueprint
    - Dynamic values are replaced by placeholders

## Kubernetes Volumes 

- Data Persistence
- Not namespaced
- 3 components of k8s storage
    1. Persistent Volume
        - needs actual physical storage,nfs(n/w file system) or cloud
        - created via YAML file
            - kind : PersistentVolume
    2. Persistent Volume Claim
        - Also created with yaml configuration
        - Pod requests the volume through the PV claim. Claim tries to find a volume in cluster
        - Volume has the actual storage backend
    3. Storage Class
        - automatically create persistent volume dynamically in the background

## StatefulSet

- For databases in load-balancing to avoid data inconsistency 
- Eg: all databases
- Stateful applications are deployed using statefulset
- Replica pods are not identical.<b>Pod Identity</b> 
  - created with same specification but not interchangeable
- Scaling databse applications
  - Master : used to write/read data
  - Worker : only used to read data 
  - Master and worker not use the same physical storage they individually keep separate replicas of storage. So they continously synchronize the data
  - Master stors the data and all the workers change according to master
  - If all pods die, the data will be lost. So it is recommended to use <b>data persistence</b>


## Kubernetes Service
 
- Each pod has its own ip address.
- Service give a stable IP address and also provide loadbalancing
- Types:
  1. ClusterIP service
    - default type
    - determine which pods to forward the request to by using a parameter called "selector" which is labels of pods
    - determine which port the pod need to forward is by using the parameter called "targePort". targetPort must match the port the container is listening at❗

  2. Headless Serirvice
    - client want to communicate with a specific pod
    - want to talk directly with specific pod and not randomly selected
    - used on stateful application
    - create it by setting `clusterIP : None`

  3. Nodeport
    - port that is exposed on each node externally (ranges 30000 to 32767) ❗
    - not secure to externally open a port outside

  4. Loadbalancer
    - accessible externally through cloud provider loadbalancer funtionality
    - nodeport and clusterip are created automatically  
