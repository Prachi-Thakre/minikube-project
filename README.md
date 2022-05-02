# minikubeproject
 
---
title: minikubeproject

---

{ % capture overview % }

The goal of this assignmet is to turn a simple Hello World go app
into an application running on Kubernetes. The tutorial shows you how to
take code that you have developed on your machine, turn it into a Docker
container image and then run that image on [Minikube]
Minikube provides a simple way of running Kubernetes on your local machine for free.

{% endcapture %}

{% capture objectives %}

* Run a hello world go application.
* Deploy the application to Minikube.
* View application logs.
* Update the application image.



## Create a Minikube cluster

 start the Minikube cluster:

```shell
minikube start --vm-driver=xhyve
```

The `--vm-driver=xhyve` flag specifies that you are using Docker for Mac. The
default VM driver is VirtualBox.

Now set the Minikube context. The context is what determines which cluster
`kubectl` is interacting with. You can see all your available contexts in the
`~/.kube/config` file.

```shell
kubectl config use-context minikube
```

Verify that `kubectl` is configured to communicate with your cluster:

```shell
kubectl cluster-info
```

## Create your go application

The next step is to write the application. Save this code in a folder named `minikubeproject`
with the filename `main.go`:


Run your application:

```shell
go run main.go
```

You should be able to see your "Hello World!" message at http://localhost:3000/.

Stop the running go server by pressing **Ctrl-C**.

The next step is to package your application in a Docker container.

## Create a Docker container image

Create a file, also in the `minikubeproject` folder, named `Dockerfile`. A Dockerfile describes
the image that you want to build. 


since we are using Minikube, instead of pushing your Docker image to a
registry, you can simply build the image using the same Docker host as
the Minikube VM, so that the images are automatically present. To do so, make
sure you are using the Minikube Docker daemon:

```shell
eval $(minikube docker-env)
```

**Note:** Later, when you no longer wish to use the Minikube host, you can undo
this change by running `eval $(minikube docker-env -u)`.

Build your Docker image, using the Minikube Docker daemon:

```shell
docker build -t minikubeproject:v1 .
```
Now the Minikube VM can run the image we built.

## Create a Deployment

Use the `kubectl run` command to create a Deployment that manages a Pod. The
Pod runs a Container based on your `minikubeproject:v1` Docker image:

```shell
kubectl run minikubeproject --image=minikubeproject:v1 --port=
```

View the Deployment:


```shell
kubectl get deployments
```

Output:


```shell
NAME             DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
minikubeproject   1         1         1            1           3m
```

View the Pod:


```shell
kubectl get pods
```

Output:


```shell
NAME                             READY     STATUS    RESTARTS   AGE
minikubeproject-714049816-ztzrb   1/1       Running   0          6m
```

View cluster events:

```shell
kubectl get events
```

View the `kubectl` configuration:

```shell
kubectl config view
```



## Create a Service

By default, the Pod is only accessible by its internal IP address within the
Kubernetes cluster. To make the `hello-node` Container accessible from outside the
Kubernetes virtual network, you have to expose the Pod as a
Kubernetes 

From your development machine, you can expose the Pod to the public internet
using the `kubectl expose` command:

```shell
kubectl expose deployment hello-node --type=LoadBalancer
```

View the Service you just created:

```shell
kubectl get services
```

Output:

```shell
NAME              CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
minikubeproject   10.0.0.71    <pending>     3000/TCP   6m
kubernetes        10.0.0.1     <none>        443/TCP    14d
```

The `--type=LoadBalancer` flag indicates that you want to expose your Service
outside of the cluster. On cloud providers that support load balancers,
an external IP address would be provisioned to access the Service. On Minikube,
the `LoadBalancer` type makes the Service accessible through the `minikube service`
command.

```shell
minikube service minikubeproject
```

This automatically opens up a browser window using a local IP address that
serves your app and shows the "Hello World" message.

Assuming you've sent requests to your new web service using the browser or curl,
you should now be able to see some logs:

```shell
kubectl logs <POD-NAME>
```

## Update your app

Edit your `main.go` file to return a new message:

```golang
fmt.Fprintf(w, "<h1>Hello World Again</h1>")

```

Build a new version of your image:

```shell
docker build -t minikubeproject:v2 .
```

Update the image of your Deployment:

```shell
kubectl set image deployment/minikubeproject hello-node=minikubeproject:v2
```

Run your app again to view the new message:

```shell
minikube service minikubeproject
```

## Clean up

Now you can clean up the resources you created in your cluster:

```shell
kubectl delete service minikubeproject
kubectl delete deployment minikubeproject
```

Optionally, stop Minikube:

```shell
minikube stop
```

{% endcapture %}



{% endcapture %}

