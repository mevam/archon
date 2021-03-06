Archon
======

[![CircleCI](https://circleci.com/gh/kubeup/archon/tree/master.svg?style=shield)][circleci]

Archon is an open source tool for cluster creation and daily operations.
It is primarily targeted for Kubernetes clusters but extending to other
distributed clusters should be easy due to its declarative nature.

Archon is designed following the principles of Kubernetes and it works as
an extension of Kubernetes using [ThirdPartyResource]. You can define your
cluster with yaml files then use `kubectl` to create and scale it.

WARNING: Archon is currently in alpha status.

Features
--------

Archon itself is an execution engine. It provides following fundamental capabilities:

  - Certificates management. Creating CA, signing new certificates with `Secret`.
  - Network management. Create VPC from definition resource and manage its lifecyle.
  - User management. Create user definitions used as default users for instances.
  - Instance group management. Manage a group of instances with similar configurations.
    Scale the group by changing the `replicas` field.
  - Instance management. Create instance with generated `userdata` for [cloudinit]
    from definition resource. Manage the instance lifecyle by watching its status.

We have put together some examples to showcase various ways to bootstrap and manage
a Kubernetes cluster with Archon with following features:

  - HA master to prevent single point of failure
  - Rolling update by creating a new instance group and delete the old one
  - Support various operating systems using [bootkube] or [kubeadm]
  - Scale the cluster with one command

Why Archon
----------

We already have tools like [kubeadm] and [kops]. Why do we need another tool
for a similar job? Here are bunch of reasons:

  1. Tools exist today are mainly imperative. Declarative tools will help
  users share their experience by publish their own cluster definitions.
  Also users could easily make small customization based on existing template
  using declarative tools.
  2. [Self-hosted Kubernetes] sounds like a promising idea and we need cluster
  management tools which could be easily integrated into current Kubernetes
  architecture.
  3. There're many different environment to launch a Kubernetes cluster in,
  an unified experience and easy to extend codebase will help developers to
  collaborate by adding more cloud provider and os support.

How it works
------------

Archon use Kubernetes as its base. Just like the way you use Kubernetes,
define resources in your cluster by creating yaml files using our customized
resource types. Then manage its lifecycle using `kubectl`. Until now, we
support these resources:

  - User
  - Network
  - InstanceGroup
  - Instance

The `archon-controller` which you should launch beforehand in the cluster will
create and manage its status based on your definition.

You might find the idea that bootstrapping a new Kubernetes cluster with an existing
one too complex. But after your cluster got initialized, you can move all the
definitions and controllers into the new cluster and let it manage itself. It will
be very convenient that you can manage both applications running in your cluster
and the cluster itself with just the `kubectl` cli tool.

Supported platforms
-------------------

Supported cloud providers:

  - [AWS]
  - [Aliyun]

Supported operating system:

  - [CoreOS][bootkube-example]
  - [Ubuntu][ubuntu-example]
  - [CentOS][centos-example]
  - [RedHat][redhat-example]

Supported bootstrapping tool:

  - [bootkube][bootkube-example]
  - [kubeadm][ubuntu-example]

At the moment, we only support limited cloud providers and oses. More cloud providers
and operating systems support will be added when the core is stable.

Installation
------------

You could launch Archon locally or install it into your cluster.

### Download

Download the latest release from Github.

```
wget https://github.com/kubeup/archon/releases/download/v0.2.0/archon-controller-v0.2.0-linux-amd64.gz
gunzip archon-controller-v0.2.0-linux-amd64.gz
chmod +x archon-controller-v0.2.0-linux-amd64
mv archon-controller-v0.2.0-linux-amd64 /usr/local/bin/archon-controller
```

On OSX, just change `linux` to `darwin`:

```
wget https://github.com/kubeup/archon/releases/download/v0.2.0/archon-controller-v0.2.0-darwin-amd64.gz
gunzip archon-controller-v0.2.0-darwin-amd64.gz
chmod +x archon-controller-v0.2.0-darwin-amd64
mv archon-controller-v0.2.0-darwin-amd64 /usr/local/bin/archon-controller
```

### Launch locally

Then config AWS credentials and start running the embeded Kubernetes server along with `archon-controller`:

```
export AWS_ACCESS_KEY_ID=YOUR_AWS_KEY_ID
export AWS_SECRET_ACCESS_KEY=YOUR_AWS_SECRET
export AWS_ZONE=YOUR_CLUSTER_ZONE
archon-controller --local --cloud-provider aws
```

You can also use Aliyun:

```
export ALIYUN_ACCESS_KEY=YOUR_ALIYUN_KEY_ID
export ALIYUN_ACCESS_KEY_SECRET=YOUR_ALIYUN_SECRET
archon-controller --local --cloud-provider aliyun
```

The server will listen on `localhost:8080` and server data will be saved to `./.localkube` folder.

### Create cluster resource with kubectl

By default, `kubectl` will talk with the server on `localhost:8080`, you can create an example cluster with:

```
kubectl create -f examples/k8s-simple
```

After a while, you could get the ip address for the server:

```
kubectl get instance -o yaml
```

And ssh into the server with the default password `archon`:

```
ssh core@SERVER_IP
```

### In cluster deployment

Please follow these instructions if you plan to deploy Archon into your cluster:

  - [AWS installation instructions]
  - [Aliyun installation instructions]


Example
-------

We believe every cluster is different. So we provides some ways to bootstrap a
Kubernetes cluster as examples. You could choose one as the starting point and
customize it to match your needs.

  - [Simple one machine Kubernetes cluster][simple-example]
  - [One master multiple nodes Kubernetes cluster][master-node-example]
  - [Self-hosted Kubernetes cluster with bootkube][bootkube-example]
  - [Kubernetes cluster with Ubuntu and kubeadm][ubuntu-example]
  - [Kubernetes cluster with CentOS and kubeadm][centos-example]
  - [Kubernetes cluster with RedHat and kubeadm][redhat-example]
  - [Aliyun one machine Kubernetes cluster][aliyun-example]
  - [Three nodes etcd cluster][etcd-example]

[ThirdPartyResource]: http://kubernetes.io/docs/user-guide/thirdpartyresources/
[kubeadm]: https://github.com/kubernetes/kubeadm
[bootkube]: https://github.com/kubernetes-incubator/bootkube
[kops]: https://github.com/kubernetes/kops
[cloudinit]: http://cloudinit.readthedocs.io/en/latest/
[AWS]: https://aws.amazon.com
[Aliyun]: https://www.aliyun.com
[simple-example]: https://github.com/kubeup/archon/tree/master/example/k8s-simple
[master-node-example]: https://github.com/kubeup/archon/tree/master/example/k8s-master-node
[bootkube-example]: https://github.com/kubeup/archon/tree/master/example/k8s-bootkube
[ubuntu-example]: https://github.com/kubeup/archon/tree/master/example/k8s-ubuntu
[centos-example]: https://github.com/kubeup/archon/tree/master/example/k8s-centos
[redhat-example]: https://github.com/kubeup/archon/tree/master/example/k8s-redhat
[aliyun-example]: https://github.com/kubeup/archon/tree/master/example/k8s-aliyun
[etcd-example]: https://github.com/kubeup/archon/tree/master/example/etcd-cluster
[Self-hosted Kubernetes]: https://github.com/kubernetes/community/pull/206
[AWS installation instructions]: https://github.com/kubeup/archon/blob/master/docs/installation.md
[Aliyun installation instructions]: https://github.com/kubeup/archon/blob/master/docs/installation_aliyun.md
[circleci]: https://circleci.com/gh/kubeup/archon
