# ml_server

Go server with frontend and backends that, in theory, could talk to each other if a k8s config existed. Go dependencies are vendored (i.e. fully contained in the checkout).

To build:
```bash
cd <ml_server checkout root>
TAG="<username>/buildenv:latest"
docker build -t $TAG go/buildenv
# start the container, mount the cwd as a volume at /usr/src/app in the container
# TODO: needing to mount the bazel cache is funky
docker run -t -d --name build -u bazel-user -v "$PWD":/usr/src/app -v "$HOME/.cache/bazel:/home/bazel-user/.cache/bazel" $TAG
# run the build command
docker exec build bazel build go/src/...

...
# later, clean up:
docker stop build
# nuke the container if you don't want it anymore
docker rm build
```

It will take a few minutes the first time, but the docker container will stay running and maintain a cache, and subsequent builds will be fast. Build output will be written to the root of the project checkout.

Kubernetes/minikube deployment coming next...