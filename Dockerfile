FROM golang:1.21

WORKDIR /app
COPY . .

RUN go build -o awesomeProject1 ./awesomeProject1/main.go

CMD ["./awesomeProject1"]