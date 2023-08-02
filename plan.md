# Plan

07/19/2023

Setup the project with core libraries that allow to exercise the main functionality of resizing image assets. First is Android.

- [x] 1. Add `cobra` for CLI
- [x] 2. Add `ImageMagick` for image processing
- [x] 3. Define unit tests and test data for Android
- [x] 4. Add internal Kaiju CI pipeline
- [x] 5. Implement to pass unit tests for Android
- [x] 6. Add `Dockerfile` and document usage in `README.md`
- [x] 7. Add internal Kaiju CD pipeline to publish container image to `Docker Hub`
- [ ] 8. Define `Tekton` task to run asset transformer
- [ ] 9. Add `Tekton` task to `Tekton Hub`
- [ ] 10. Start work on iOS asset transformer and asset dictionary
