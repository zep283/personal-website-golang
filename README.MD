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