FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/*.sql ./migrations/

RUN go build -o /migrate

CMD [ "/migrate" ]
