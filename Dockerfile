FROM golang:1.21.1

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o app cmd/betera-test-task/main.go

RUN chmod +x app

EXPOSE 8000

CMD [ "./app" ]