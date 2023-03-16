<img  src="images\k8_example_components.png"/>


## Deployment for mongodb

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongodb-deployment
  labels:
    app: mongodb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:       # container 
      - name: mongodb
        image: mongo    #  pull mongo image from docker hub
        ports:
        - containerPort: 27017 # port in the container listen on
        env:
        - name: MONGO_INITDB_ROOT_USERNAME # will go to repo and not contain any password❗
            valueFrom:                     # pull the value from secret component
                secretKeyRef:
                    name: mongodb-secret
                    key: mongo-root-username
        - name: MONGO_INITDB_ROOT_PASSWORD
            valueFrom:
                secretKeyRef:
                    name: mongodb-secret
                    key: mongo-root-password
--- # syntax for document separation
apiVersion: v1
kind: Service
metadata:
  name: mongodb-service
spec:
  selector:
    app: mongodb  # to connect to pod through label
  ports:
    - protocol: TCP 
      port: 27017   
      targetPort: 27017 # same as container port

```

## Secret for mongodb

```yaml
apiVersion: v1
kind: Secret
metadata:
    name: mongodb-secret    # name of the secret
type: Opaque                # default cred type
data:
    mongo-root-username: dXNlcm5hbWU=   # value encoded in base64
    mongo-root-password: cGFzc3dvcmQ=
```
 
- `kubectl apply -f mongo-secret.yaml` : after secret is creted it can be referenced by the deployment

## Deployment for mongo-express

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongo-express
  labels:
    app: mongo-express
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongo-express
  template:
    metadata:
      labels:
        app: mongo-express
    spec:
      containers:
      - name: mongo-express
        image: mongo-express
        ports:
        - containerPort: 8081
        env:
        - name: ME_CONFIG_MONGODB_ADMINUSERNAME
          valueFrom:
            secretKeyRef:
              name: mongodb-secret
              key: mongo-root-username
        - name: ME_CONFIG_MONGODB_ADMINPASSWORD
          valueFrom: 
            secretKeyRef:
              name: mongodb-secret
              key: mongo-root-password
        - name: ME_CONFIG_MONGODB_SERVER
          valueFrom: 
            configMapKeyRef:
              name: mongodb-configmap
              key: database_url   # taken from configmap
---
apiVersion: v1
kind: Service # external service to access express externally
metadata:
  name: mongo-express-service
spec:
  selector:
    app: mongo-express
  type: LoadBalancer  # accepts external requests
  ports:
    - protocol: TCP
      port: 8081
      targetPort: 8081
      nodePort: 30000 # port that external ip is open (30000 - 32767)

```

`minikube service mongo-express-service` : to get a public ip adress in minikube and not in actual kubernetes

## Configmap for mongo-express

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mongodb-configmap
data:
  database_url: mongodb-service   # automatically detect the ip from servicename
```