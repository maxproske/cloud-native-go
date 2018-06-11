# Build image from Dockerfile
docker build -t cloud-native-go:1.0.1 .
docker images

# Tag image
docker tag cloud-native-go:1.0.1 maxproske/cloud-native-go:1.0.1
docker images

# Push to Docker Hub
docker login
docker push maxproske/cloud-native-go:1.0.1

# Run image
docker run -it -p 8080:80 maxproske/cloud-native-go:1.0.1	# Interactive terminal
docker run -d -p 8080:80 maxproske/cloud-native-go:1.0.1	# Detached
docker run --name cloud-native-go --cpu-quota 50000 --memory 64m --memory-swappiness 0 -d -p 8080:80 maxproske/cloud-native-go:1.0.1	# With quotas
docker ps

# Stats
docker stats

# Process Status
docker stop $(docker ps -q)		# Stop all containers gracefully
docker kill $(docker ps -q)		# Stop all containers immediately
docker rm $(docker ps -a -q)	# Remove all containers
docker rmi $(docker images -q)	# Remove all docker images

# Pull
docker pull maxproske/cloud-native-go:1.0.1

# Compile for Linux
GOOS=linux GOARCH=amd64 go build
docker build -t cloud-native-go:1.0.1-alpine .
docker tag cloud-native-go:1.0.1-alpine maxproske/cloud-native-go:1.0.1-alpine
docker push maxproske/cloud-native-go:1.0.1-alpine

# Compose up
docker-compose build
docker-compose.exe up -d
docker ps # One for nginx, one for microservice

# Compose kill all
docker-compose kill
docker-compose rm

# Install Kubernetes
https://github.com/kubernetes/minikube/releases

# Start Kubernetes using Hyper-V
#   If you accidentally use VirtualBox, delete C:\Users\Max\.minikube and start minikube again
#   If minikube exits with status 1, your machine doesn't have enough memory to run the VM.
#     Set: HyperV Manager > minikube > Settings > Memory > uncheck 'Enable Dynamic Memory'
minikube delete
minikube start --vm-driver="hyperv" --hyperv-virtual-switch="myswitch"
minikube status