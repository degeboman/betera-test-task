FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-docker-demo

EXPOSE 8000

CMD [ "/go-docker-demo" ]