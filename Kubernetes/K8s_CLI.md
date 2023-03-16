`kubectl get nodes` - list all the nodes

`kubectl get pod` - check the pods

`kubectl get pod -o wide` - get extra informations on pods

`kubectl get services` - list the services

`kubectl get deployment` - list the deployments

`kubectl get replicaset` - manage the replicas of the pod

### Create components

`kubectl create deployment nginx-depl --image=nginx` - create a pod with name nginx-depl

`kubectl edit deployment nginx-depl` - edit the configuration file. Redeploy new one and old one will be terminated

### Debugging commands

- `kubectl logs [pod name]`

- `kubectl describe pod [pod name]` - list the statechanges inside the pod

- `kubectl exec -it [pod name] -- bin/bash` - return the interactive terminal inside the container

### Others

- `kubectl delete deployment [deployment name]` - delete the deployment

- `kubectl apply -f [configuration file name]` - deployment with configuration file

- `kubectl delete -f [configuration file name]` - delete the deployment specified in the configuration file

- `kubectl get deployment nginx-deployment -o yaml > nginx-result.yaml` - output the configuration status after pod creation