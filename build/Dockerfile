FROM golang:latest as builder

ENV GOOS linux
ENV CGO_ENABLED 0

COPY . /home/
WORKDIR /home

RUN go mod download

RUN go build -o self-zep-lab

FROM alpine:3.16 as production

COPY --from=builder /home/self-zep-lab .
COPY --from=builder /home/web ./web

EXPOSE 8080 
CMD ./self-zep-lab