# distributed-web

# create docker image
docker build -t fin-tech . 

# check if image is created
docker image ls

# run application from docker
docker run -p 29999:29999 -tid fin-tech

sudo lsof -nP -iTCP -sTCP:LISTEN

# Push to dockerhub
docker login --username=echenim007
docker tag distributed echenim007/fin-tech:1.0.0
docker push echenim007/fin-tech:1.0.0


# copy to kubernetes server
scp -r kubernetes/distributed-web elk@192.168.1.212:dev

docker rmi -f imageID