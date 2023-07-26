FROM golang:1.21rc3-alpine3.18

WORKDIR /app
COPY . .

RUN go build -o app
EXPOSE 8080

CMD ["./app"]