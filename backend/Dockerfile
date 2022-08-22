FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build .

EXPOSE 3000

CMD [ "/app/backend" ]
