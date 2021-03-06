FROM golang:1.13 as builder

WORKDIR /build

RUN apt-get update && apt-get install -y upx

ENV GOPROXY=https://goproxy.io \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY . .

RUN go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
RUN ./build/build.sh k8s-event-logger-operator cmd/manager/main.go

# application image
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /opt/go

LABEL maintainer="bakito <github@bakito.ch>"
EXPOSE 8080
USER 1001
ENTRYPOINT ["/opt/go/k8s-event-logger-operator"]

COPY --from=builder /build/k8s-event-logger-operator /opt/go/k8s-event-logger-operator
