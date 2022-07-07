FROM golang:alpine as user-server
ENV GO111MODULE=on
WORKDIR /server
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /go/bin/user-service cmd/main.go

FROM scratch
COPY --from=user-server /go/bin/user-service /go/bin/user-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/user-service"]