## Node

- Types :
    1. Master Node
    2. Slave Node / Worker Node
- Each node has multiple pods on it
- 💡 3 process installed on every Node:
    1. Container Runtime : eg:- Docker
    2. Kubelet : interacts with both container runtime and node. Assigning cpu / ram resources.
    3. Kube Proxy
        - Forwarding communication from services to pod  

## Master node

- Manage cluster state and control the worker nodes
- Each k8 cluster made upof multiple master node
- 4 processes:

    1. Api server
        - Cluster gateway
        - Acts as a gatekeeper for athentication 
        - Entry point into the cluster
    2. Scheduler
        - Handle the request coming from api server
        -  Determine where to put the pod on the node according to the resource availability using kubelets on the worker node
    3. Controller Manager
        -  If pods die on any node, detects it and resheduel as soon as possible
        - It detects cluster state changes
    4. etcd
        - Key-value store on the cluster state
        - 🧠 Cluster brain
        - Cluster changes get stored in the key value store
        - All the other process uses etcd to look for state change