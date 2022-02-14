FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o fcc-timestamp-ms

EXPOSE 3000

CMD ./fcc-timestamp-ms