# distributed-web

# create docker image
docker build -t distributed . 

# check if image is created
docker image ls

# run application from docker
docker run -p 29999:29999 -tid distributed

sudo lsof -nP -iTCP -sTCP:LISTEN

# Push to dockerhub
docker login --username=echenim007
docker tag distributed echenim007/distributed:v1.0.1
docker push echenim007/distributed:v1.0.1


# copy to kubernetes server
scp -r kubernetes/distributed-web elk@192.168.1.212:dev

docker rmi -f imageID