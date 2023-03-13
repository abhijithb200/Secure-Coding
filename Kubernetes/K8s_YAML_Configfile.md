## Deployment config

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:                   # deployment
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:               # pods - image detail
      containers:
      - name: nginx     
        image: nginx:1.16
        ports:
        - containerPort: 8080
```

## Service Config

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  selector:
    app: nginx          # look for the same label in deployments
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080  # target port need to match the container port
```
