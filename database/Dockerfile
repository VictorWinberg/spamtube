FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/*.sql ./migrations/

RUN ls

RUN go build -o /migrate

CMD [ "/migrate" ]