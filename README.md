Simple GO application that returns the current time.

Run it with:

`go run time.go`

Open following url: `http://localhost:8080/time`

App can also be dockerized with following commands:

Create image from Dockerfile: 

`docker build --tag cavke/go-time:1.0 .`

Run image (map local port 8000 to port exposed by docker 8080): 

`docker run --detach --publish 8000:8080 cavke/go-time:1.0`

Check running containers:

`docker ps`

Stop and remove running container:

`docker rm --force containerID`

