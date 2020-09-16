# Simple emptyDir volume example using SSE

# Description
Example of emptyDir volume sharing between two containers within a Pod. A file will be created with the current timestamp by the `producer` and the `consumer` is a web server which is reading the file every second.

Through Server Sent Events and a basic web page the current value is show to the user and pushed on every change.

# Running through Docker-Compose
Within the root folder simply run:

```shell
docker-compose up --build --detach consumer producer
```

The web server can now be reached through `localhost:8080`.

To shutdown the system run following command:

```shell
docker-compose down
```

# Running within a local Kubernetes
- Build both container images by running following commands from the root folder:

```shell
docker build -t consumer:latest consumer/.
docker build -t producer:latest producer/.
```

Apply the provided manifest files within the k8s folder:

```shell
kubectl apply -f k8s/
```

Forward the port to the created service:

```shell
kubectl port-forward svc/k8s-emptydir 8080:8080
```

The web server can now be reached through `localhost:8080`.

To shutdown the system run following command:

```shell
kubectl delete -f k8s/
```
