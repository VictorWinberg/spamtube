FROM golang:1.18-alpine

WORKDIR /app

# install FFmpeg
RUN apk add --no-cache ffmpeg

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /generator

CMD [ "/generator" ]
