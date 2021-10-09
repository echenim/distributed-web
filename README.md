# distributed-web

# create docker image
docker build -t distributed-web . 

# check if image is created
docker image ls

# run application from docker
docker run -p 30000:30000 -tid distributed-web

sudo lsof -nP -iTCP -sTCP:LISTEN