# simple API webserver
# Linux x64
FROM ubuntu

# LABEL key="value"

# install Golang
RUN sudo apt update
RUN sudo apt upgrade
RUN sudo apt install golang-go
RUN mkdir -p ~/go/src


# copy app code to correct location in container
COPY . ~/go/src

# set working directory
WORKDIR ~/go/src

# install dependencies
RUN go get

# command for container to execute
ENTRYPOINT [ "go", "main.go" ]