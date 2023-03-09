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

- To remove a container: `docker rm <container id>`
- To remove an image : `docker rmi <image id>`
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

- Stream the logs - `docker logs <container id> -f`

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

## Docker Compose

- .yaml configuration file
- Docker compose takes care of creating a common network
- Docker volume gives data persistence
    - Folder in physical host file system is mounted into the virtual file system of Docker

```yaml
version:'3'  # version of docker-compose
services:
    mongodb:    # container name
        image: mongo    # image name
        ports:
            - 27017:27017
        environment:
            - ME_CONFIG_MONGODB_ADMINUSERNAME=admin
        volumes:
            - mongo-data:/data/db   # specify the directory to store persistence inside container

volumes:    # list all the volumes used 
    mongo-data:
        driver:local
```

- Run with docker compose

    `docker-compose -f <compose.yaml file> up`

- Stop all the containers; also removes the network

    `docker-compose -f <compose.yaml file> down`

## Dockerfile

- Blueprint for our image
- Can have multiple RUN command but only one CMD command
- Always called Dockerfile

```Dockerfile
FROM node:13-alpine   # image with node version installed (base image)

ENV MONGO_DB_USERNAME=admin MONGO_DB_PWD=password   # set environmental variable in the image environment

RUN mkdir -p /home/app  # run any linux command

COPY . /home/app    # copy host files to /home/app on container

CMD ["node","/home/app/server.js"]    #executes the entry point of the application
```

- Build an image from dockerfile
    - -t : Tag (name)

`docker build -t my-app:1.0 .`