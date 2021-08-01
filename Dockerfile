FROM golang:1.16.0 as builder
 
WORKDIR /go/src
ENV CGO_ENABLED=0
 
COPY go.mod go.sum ./
RUN go mod download
 
COPY . .
 
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/main -ldflags '-s -w'
 
FROM scratch
COPY --from=builder /go/bin/main /
CMD ["/main"]