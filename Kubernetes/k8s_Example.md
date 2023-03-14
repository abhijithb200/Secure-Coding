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
    app: mongodb
  ports:
    - protocol: TCP
      port: 27017
      targetPort: 27017

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