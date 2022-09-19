FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY .env .env.local ./

RUN go build -o ./server

EXPOSE 3000

CMD [ "./server" ]
