# ml_server

Go server with frontend and backends that, in theory, can talk to each other. Go dependencies are vendored (i.e. fully contained in the checkout). Lots of hacks, hardcoding, etc, but end result is nice.

Presetup (once per machine):
```bash
# 1. Install jdk8: http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html
# 2. Install brew (from: https://brew.sh/)

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# 3. Install bazel (from: https://docs.bazel.build/versions/master/install.html) 

brew install bazel

# 4. Install minikube (from: https://kubernetes.io/docs/tutorials/stateless-application/hello-minikube/)

curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 && \
  chmod +x minikube && \
  sudo mv minikube /usr/local/bin/

brew install docker-machine-driver-xhyve
sudo chown root:wheel $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
sudo chmod u+s $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve

brew install kubectl
```

Setup (once per reboot/logon):
```bash
# 1. Start minikube
minikube start --vm-driver=xhyve
# 2. Set minikube to default env for kubectl
eval $(minikube docker-env)
```

Build/Run server:
```bash
# This will build the go code, build docker images, push them to jdelfino's docker hub, then launch the cluster in your local minikube.
# Time for successive rebuild of frontend & backend and redeploy: 7.5s!
bazel run --cpu=k8 go/src/ml_server:server.apply
# Opens a web browser to the frontend, which should show some hardcoded text.
minikube service ml-server-frontend
```
