
FROM golang:1.14

WORKDIR /go/src/app
COPY . .
RUN apt-get update && \
      apt-get -y install sudo
RUN go get -d -v ./...
RUN go install -v ./...
RUN cd  /go/src/app && go build -o app .

CMD ["app"]