# Kaiju Asset Transformer

In course of mobile app development, a developer receives an asset file from the designer. In order to use the asset in the app, the developer needs to resize the asset to the required sizes. This is a tedious and error-prone process. Kaiju Asset Transformer is a tool that resizes the asset to the required sizes and generates the asset dictionary that can be used in the app. It also conforms to the asset file naming convention of the target platform.

Initially, only Android is supported. iOS support is planned.

# Docker

The Dockerfile is used to build the container image that can be used to run the asset transformer. At the moment `deploy/ci/Dockerfile` is used for CI builds. It has `test` and `build` targets. In future `deploy/cd/Dockerfile` will be used for CD builds. It will have `build` and `publish` targets. Additionally `deploy/prod/Dockerfile` will provide a container image which can be used to transform assets in any container environment.

The CI `Dockerfile` depends on a base image, `Dockerfile.gomagick`, so that it sources cached image during CI builds. This speeds up the build process by approx `30 seconds`, rest will require configuring the build and module cache. The base image is being published manually atm-

```bash
# build the image locally
docker build -t gomagick:2.0 -f deploy/ci/Dockerfile.gomagick .

# tag the image
docker tag a32623302329 kaijuci/gomagick:2.0

# push the image
docker push kaijuci/gomagick:2.0
```

To build prod:

```bash
docker build -t xfrmr:<tag> --build-arg BUILD_VERSION=`git rev-parse --short HEAD` -f deploy/prod/Dockerfile .
```

# Plan

Refer to [plan.md](plan.md) for the planned tasks and ongoing progress.
