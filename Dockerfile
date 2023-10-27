FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build ./cmd/betera-test-task /app

EXPOSE 8000

CMD [ "/app" ]
