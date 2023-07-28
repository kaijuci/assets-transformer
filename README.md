# Kaiju Asset Transformer

In course of mobile app development, a developer receives an asset file from the designer. In order to use the asset in the app, the developer needs to resize the asset to the required sizes. This is a tedious and error-prone process. Kaiju Asset Transformer is a tool that resizes the asset to the required sizes and generates the asset dictionary that can be used in the app. It also conforms to the asset file naming convention of the target platform.

Initially, only Android is supported. iOS support is planned.

# Dockerfile

The Dockerfile is used to build the container image that can be used to run the asset transformer. At the moment `deploy/ci/Dockerfile` is used for CI builds. It has `test` and `build` targets. In future `deploy/cd/Dockerfile` will be used for CD builds. It will have `build` and `publish` targets. Additionally `deploy/prod/Dockerfile` will provide a container image which can be used to transform assets in any container environment.

# Plan

Refer to [plan.md](plan.md) for the planned tasks and ongoing progress.
