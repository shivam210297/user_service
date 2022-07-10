FROM golang:alpine as user-server
ENV GO111MODULE=on
WORKDIR /server
COPY go.mod /server/
COPY go.sum /server/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o /go/bin/user-service cmd/main.go

FROM scratch
COPY --from=user-server /go/bin/user-service /go/bin/user-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/user-service"]