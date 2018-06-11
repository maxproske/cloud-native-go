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
docker rm $(docker ps -a -q) -f	# Remove all containers
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
docker-compose up -d
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
minikube start --vm-driver="hyperv" --hyperv-virtual-switch="Primary Virtual Switch"
minikube status
kubectl cluster-info # Check we are working locally

# Run the provided command to work with the docker daemon
minikube docker-env
docker ps

# Create pod
kubectl create -f k8s-pod.yml
kubectl get pods    # List
kubectl get pods -o wide # Detailed list
kubectl describe pod cloud-native-go    # Describe
kubectl delete pod cloud-native-go  # Delete

# Access service
kubectl port-forward cloud-native-go 8080:8080 # View at http://localhost:8080/api/books

# Create namespace
kubectl get ns # Get namespace
kubectl get pods --namespace kube-system # Pods running in this namespace
kubectl create -f k8s-namespace.yml # Create own namespace
kubectl create -f k8s-pod.yml --namespace cloud-native-go # Recreate with namespace
kubectl get pods --namespace cloud-native-go # Second pod created
kubectl delete pod cloud-native-go # Delete first pod
kubectl delete --all pods --namespace=cloud-native-go # Delete second pod

# Create deployment
kubectl create -f k8s-deployment.yml
kubectl get deployments,pods,rs
kubectl apply -f k8s-deployment.yml # Update
kubectl describe deployment cloud-native-go # View constraints

# Create service
kubectl create -f k8s-service.yml

# Access pods running behind load balancing service.
# (eg http://192.168.0.29:32270/)
minikube ip
kubectl get services 

kubectl apply -f k8s-service.yml # Update
kubectl describe service cloud-native-go # Describe
kubectl delete services cloud-native-go # Delete services
kubectl delete deployments cloud-native-go # Delete deployments

# Horizontal scaling
kubectl get pods -w # In a seperate window
kubectl create -f k8s-deployment.yml --record=true # Record history
kubectl scale deployment cloud-native-go --replicas=5 # Scale up
kubectl get pods
kubectl scale deployment cloud-native-go --replicas=3 # Scale down
kubectl get pods
kubectl rollout history deployment cloud-native-go # See history

# Update image as rolling update
# kubectl set image deployment (deployment) (container)=(image)
kubectl set image deployment cloud-native-go cloud-native-go=maxproske/cloud-native-go:1.0.2-alpine

# Undo rollout
kubectl rollout undo deployment cloud-native-go