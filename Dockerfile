# workspace (GOPATH) configured at /go
FROM golang:1.17 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/delever/delever_websocket_service_v2
WORKDIR $GOPATH/src/gitlab.udevs.io/delever/delever_websocket_service_v2

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/delever_websocket_service_v2 /

ENTRYPOINT ["/delever_websocket_service_v2"]
