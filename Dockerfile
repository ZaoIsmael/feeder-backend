FROM golang:alpine AS build

# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org

WORKDIR /app
COPY . .
# Set OS as linux
RUN GOOS=linux go build -o /go/bin/main cmd/feeder-service/main.go
RUN mkdir -p tmp

EXPOSE 4000

ENTRYPOINT ["/go/bin/main"]
