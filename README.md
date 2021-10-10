# distributed-web

# create docker image
docker build -t financial-web . 

# check if image is created
docker image ls

# run application from docker
docker run -p 29999:29999 -tid financial-web

sudo lsof -nP -iTCP -sTCP:LISTEN

# Push to dockerhub
docker login --username=echenim007
docker tag distributed echenim007/financial-web:latest
docker push echenim007/financial-web:latest


# copy to kubernetes server
scp -r kubernetes/distributed-web elk@192.168.1.212:dev

docker rmi -f imageID