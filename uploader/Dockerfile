FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

COPY .credentials .credentials

RUN go build -o ./uploader

CMD [ "./uploader" ]
