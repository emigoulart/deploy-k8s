# deploy-k8s
A simple go app to deploy with K8s
To create the app:
go mod github.com/yourusername/deploy-k8s
create a main.go file with a "Hello World"
go run main.go and in a new tab in the terminal curl localhost:8080 

Create a Dockerfile and a docker-compose.yaml (allow multiples containers) file. 
docker rm ( to kill all containers)
docker-compose up -d ( to up a container) 

docker-compose ps ( to check the container running )

docker-compose exec goapp bash ( this will execute the go app into the container, that is based in the latest go image )
