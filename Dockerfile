FROM golang:1.19

WORKDIR /dist

COPY go.mod .

RUN go mod download

COPY . .

EXPOSE 3000
RUN go build -o ./bin/main
ENTRYPOINT ["./bin/main"]
