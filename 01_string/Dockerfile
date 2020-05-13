################################################################################
# /usr/local/go           ... GOROOT (go runtime binary is arranged)
# /go                     ... GOPATH
#   /src
#     /github.com
#       /krdhrk10 ... GO_APP_ROOT
#         go.mod
#         main.go
#         ...etc...
################################################################################
FROM centos:latest

ARG APP_DOMAIN=github.com/krdhrk10
ARG GOLANG_VERSION=1.13.1

RUN yum update -y && yum install -y git postgresql

# golang
RUN curl -o /usr/local/src/golang.tar.gz https://dl.google.com/go/go$GOLANG_VERSION.linux-amd64.tar.gz
RUN tar xf /usr/local/src/golang.tar.gz -C /usr/local/
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
WORKDIR $GOPATH
RUN mkdir -p src bin && chmod -R 777 "$GOPATH"
ENV GO_APP_ROOT=$GOPATH/src/$APP_DOMAIN/
