# AjoCard DevOps and SRE Challenge
This is a golang and nodejs application responding "Hi"
- This stimulates long-lived connection and a random 10 seconds sleep has been
- The golang retrieves a response from a nodejs backend, when the HTTPclient issues a request, it starts a TCP commection with the HTTP server(nodejs) and maintains a long-lived TCP connection

# Deployment Architecture.
We have deployed a gitops approach.

# Continious Integration
We implemented a github-action to handle the continious integration of our project when the push is done to the repository, the github action builds the image and deploys it to my public docker account.

# Continious Delivery.
In order to implement an automated delivery process, we have implemented the use of fluxcd "gitops approach" in a folder called Flux.
The instructions their gives you a minimalist approach to installing flux on your kubernetes cluster and integrating it to you repository.
It takes at least 5 mins for the deployment to be initiated once code has been pushed to your repository.

However, you can simply run the yaml file in the workflow to simply build the application on your k8s cluster after you have pushed and github action has built the image.

# Stucture of Files
├── FLUX
│   ├── flux1.yaml
│   └── installation\ instructions
├── README.md
├── client
│   ├── Dockerfile
│   └── client.go
├── docker-compose.yml
├── namespaces
│   └── ajo.yaml
├── server
│   ├── Dockerfile
│   └── server.js
└── workloads
    ├── golang.yaml
    └── node-dep.yaml

# Number of apps
In each of the dep files contains the deployment manifest files, HorizontalScalling manifest files and Services manifest files all attached as one kubernetes manifest files.
Also in our deployment file we implemented a RollingUpdate strategy which ensures the applications uses a canary approach to deployments.

# View Application
in order to view the Nodejs response 
kubectl port-forward -n ajo deployment/nodebackend 3000:3000

to view the goland response
Because, the golang app sends its response through the teminal, the best approach is to get it through the terminal 

kubectl get pods --all-namespaces | grep client 

to copy any of the running pods name

kubectl logs -n ajo client***

# Install the apps directly using the kubectl approach

Kubectl apply -f namespaces
kubectl apply -f workloads