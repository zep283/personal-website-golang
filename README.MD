# self.zep-lab.dev
Simple golang website (:

## Set up guide from scratch for golang module
1. Download golang https://go.dev/dl/
2. `go version`
3. `mkdir person-web-golang`
4. `go mod init`
5. Create source code files (main.go, Dockerfile, etc)
6. Dowload Rancher Desktop for docker https://rancherdesktop.io
7. `docker build -f ./build/Dockerfile -t self-zep-lab --progress=plain --no-cache .`
8. `docker run -d -p 8080:8080 self-zep-lab`
9. GH Actions support - add .github/workflows/ci.yml and add corresponding secrets in repository settings
    - https://github.com/marketplace/actions/build-and-push-docker-images#inputs 
    - use tags to control which commits trigger build
    - `docker pull zep283/personal-web-golang:latest`
    - `docker run -d -p 8080:8080 zep283/personal-web-golang`