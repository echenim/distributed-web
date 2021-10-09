# distributed-web

# create docker image
docker build -t distributed-web . 

# check if image is created
docker image ls

# run application from docker
docker run -p 30000:30000 -tid distributed-web

sudo lsof -nP -iTCP -sTCP:LISTEN

# Push to dockerhub
docker login --username=echenim007
docker tag distributed-web echenim007/distributed-web:latest
docker push echenim007/distributed-web:latest


# copy to kubernetes server
scp -r kubernetes/distributed-web elk@192.168.1.212:dev     