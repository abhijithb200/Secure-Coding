`kubectl exec -i -t nginx-deployment-6845b4b8b9-9zz2b --container php  -- sh`

`kubectl logs  nginx-deployment-67f9d469bc-h5s56`

### Pushing docker into dockerhub

- `docker build -t codeguardian-xss .`
- `docker image tag codeguardian-xss sherloc/codeguardian-xss:1.2`
- `docker image push sherloc/codeguardian-xss:1.2`
