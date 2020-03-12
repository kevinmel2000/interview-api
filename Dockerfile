FROM golang:latest
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN chmod +x ./interview-api
EXPOSE 8080
CMD ./interview-api
