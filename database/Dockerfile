FROM arm32v7/golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

COPY migrations/*.sql ./migrations/

CMD [ "go run ." ]
