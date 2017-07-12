FROM ubuntu:14.04
COPY main.go /usr/share/main.go
RUN apt-get update && apt-get -y install golang-go

