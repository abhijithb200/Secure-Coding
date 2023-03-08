## Difference b/w docker and VM

- VM create a complete virtualized operating system environment, including a guest operating system and virtualized hardware layer that emulates physical hardware

- Docker create lightweight and portable environments for running applications. It shares the same host operating system, but it has its own isolated file system, network interfaces and process space. Docker run on the same host operating system, but on its own isolated environment.

## Docker Commands

- Download an image from docker hub - `docker pull redis`

- List all the downloaded docker images - `docker images`
    - Tags :  Version of the image

- To start running the redis image. Container is the running image - `docker run redis`
    - To run an image in a detached mode (background mode) - `docker run -d redis`

- To restart the container 
    ```cmd
        docker stop <container id>
        docker start <container id> #start -> restart the stopped container
    ```

- List all the running container - `docker ps`

- List running and stopped containers - `docker ps -a`
- Pull the image and run the container - `docker run redis:4.0`
- Host port is different from container host
    - Binding host port 6000 to container port 6379

```cmd
docker run -p6000:6379 -d redis
docker run -p6001:6379 -d redis:4.0
```

- Run a new container with custom name set

`docker run -d -p6001:6379 --name redis-older redis:4.0`

## Container Debugging

- View the logs - `docker logs <container id>`

- Get an interactive shell  - `docker exec -it <container id> /bin/bash`

## Docker Network

- List the networks - `docker network ls`
- Create a new network - `docker network create mongo-network`

## Project

- Download mongo and mongo-express images

```cmd
docker pull mongo
docker pull mongo-express
```

- Run mongo image
  - -e : environmental variable

`docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password --name mongodb --net mongo-network mongo`

- Run mongo express

`docker run -d -p 8081:8081 -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin -e ME_CONFIG_MONGODB_ADMINPASSWORD=password --net mongo-network --name mongo-express -e ME_CONFIG_MONGODB_SERVER=mongodb mongo-express`