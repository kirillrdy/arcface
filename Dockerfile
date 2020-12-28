FROM ubuntu:20.10
ENV DEBIAN_FRONTEND noninteractive
RUN apt update -y && apt upgrade -y
RUN apt install -y golang ca-certificates make libgl1-mesa-glx zlib1g-dev libssl-dev libffi-dev libbz2-dev
ADD . .
RUN go run main.go
