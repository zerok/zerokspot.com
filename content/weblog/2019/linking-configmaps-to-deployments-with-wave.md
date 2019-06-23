---
title: "Linking ConfigMaps to Deployments with Wave"
date: 2019-06-23T14:04:42+02:00
tags:
- kubeconeu2019
- kubernetes
- ops
---

In Kubernetes, most of the configuration for an application should be stored inside so-called ConfigMaps (or Secrets). You mount such ConfigMaps either as files into a pod's container or expose their value through environment variables to the container.

This is a quite common setup that is easy to grasp, prepare, and that usually also works with very small applications. But what happens when you change the configuration? By default: nothing. The application already has its settings and doesn't notice the change, nor does Kubernetes notify the application.

There are a couple of ways to solve this:

* The application could periodically fetch the configuration directly through the Kubernetes API.
* Some event system notifies the application of the change in configuration (it could also implement a Kubernetes controller itself and listen for changes to ConfigMaps).

All these are usually overkill for smaller applications or aren't even possible for legacy systems where you cannot really modify the application to make it more dynamically configured.

A third approach would be a controller that listens for changes to ConfigMaps (or Secrets) and basically triggers a rolling update on all Deployments, StatefulSets, etc. that have that a ConfigMap or Secret mounted.

This is pretty much what [Pusher's Wave](https://github.com/pusher/wave) project does. It registers controllers to listen for changes to Deployments, StatefulSets, and DaemonSets that have a `wave.pusher.com/update-on-config-change: "true"` annotation. Furthermore, it registers handlers for changes to secrets and ConfigMaps. When the content of either of those changes, the connected Deployments/StatefulSets/DaemonSets get updated.

If you want to learn more about the motivation behind this project,
[Joel Speed](https://joelspeed.co.uk/) gave a talk about that and Wave
at this year's [KubeCon Europe in
Barcelona](https://zerokspot.com/weblog/2019/05/27/kubecon-cloudnativecon-europe-2019/). You
can find the recording of that on
[YouTube](https://www.youtube.com/watch?v=8P7-C44Gjj8&list=PLj6h78yzYM2PpmMAnvpvsnR4c27wJePh3&index=7&t=0s).

<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/8P7-C44Gjj8" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

## Demo setup

Let's work on a little example here (you can find the complete set of resource definitions on [Github](https://github.com/zerok/wave-demo). I have a simple NGINX instance which serves a single `index.html`. The content of that `index.html` is generated out of the content of a ConfigMap:

```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  annotations:
    wave.pusher.com/update-on-config-change: "true"
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      volumes:
        - name: cfg
          configMap:
            name: application-config
      containers:
        - name: nginx
          image: nginx:latest
          ports:
            - name: http
              containerPort: 80
          volumeMounts:
            - name: cfg
              subPath: "index.html"
              mountPath: /usr/share/nginx/html/index.html
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: application-config
data:
  index.html: "Hello world"
```

I've put these two into a separate namespace, `wave-demo`. 

Next, let's setup wave inside the `wave-system` namespace:

```
$ k kustomize github.com/pusher/wave.git/config/default | k apply -f -

clusterrole.rbac.authorization.k8s.io/wave-manager-role created
clusterrolebinding.rbac.authorization.k8s.io/wave-manager-rolebinding created
secret/wave-webhook-server-secret created
service/wave-controller-manager-service created
statefulset.apps/wave-controller-manager created
```

Since I'm using RBAC in the setup, I had to modify the wave-manager-role a bit to also grant access to StatefulSets and DaemonSets. The cluster role binding also needed to have access to my wave-demo namespace. At the time of writing this, Wave's [kustomize](https://github.com/kubernetes-sigs/kustomize/) configuration contained a little bug that I've fixed with [this PR](https://github.com/pusher/wave/pull/57). Nothing fancy but you'd need to change the configuration for the controller a bit and then restart the the manager's pod:

```
# Add the extended cluster role settings
kubectl kustomize wave | kubectl apply -f -

# Restart wave
kubectl --namespace wave-system delete pod \
    wave-controller-manager-0
```

Now, when you change the data inside the application-config ConfigMap, Wave will update the deployment and you will see the new content in the frontend:

```
$ kubectl --namespace wave-demo patch configmap \
    application-config \
    -p '{"data": {"index.html": "updated"}}'
configmap/application-config patched

$ kubectl --namespace wave-demo port-forward \
    svc/nginx 8888:80 &

$ curl http://localhost:8888
updated
```

If that doesn't work for you, make sure to check the log of the wave-controller-manager-0 pod. You might have missed giving the service access to your namespace or to the Deployment kind 🙂

