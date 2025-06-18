# deploy-k8s
A simple go app to deploy with K8s
To create the app:
go mod github.com/yourusername/deploy-k8s
create a main.go file with a "Hello World"
go run main.go and in a new tab in the terminal curl localhost:8080 (or open via browser)

* Create a Dockerfile and a docker-compose.yaml (allow multiples containers) file
* docker rm -f $(docker ps -a -q) (to kill all containers)
* docker-compose up -d ( to up a container) 

* docker-compose ps (to check the container running)

* docker-compose exec goapp bash (this will execute the go app into the container, that is based in the   latest go image)

* Log into your docker-hub https://app.docker.com/

* docker build -t yourusername/deploy-k8s:latest -f prod.Dockerfile . (to generate the image)

* docker images ( to check the images )

* docker run --rm -p 8080:8080 yourusername/deploy-k8s   (to run the app ) and access http://localhost:8080

### Now let's create a cluster using Kind, https://kind.sigs.k8s.io/ 

* kind create cluster --name=insert-kluster-name-here

* kubectl cluster-info --context kind-k8slab

* kubectl create namespace insert-namespace-name-here 

*  kubectl apply -f k8s/deployment.yaml -n yournamespacename (It will create the deployment and pod in the informed namespace)

* kubctl get pods -n yournamespacename (check the pods)

*  kubectl apply -f k8s/service.yaml -n yournamespacename (It will create the service in the informed namespace)

* kubectl get svc -n yournamespacename (to check the services)

* kubectl port-forward svc/goappserversvc -n yournamespacename 8080:8080 

### https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#