FROM golang:1.18-alpine

WORKDIR /app

# install dependencies
RUN apk add --no-cache ffmpeg

RUN apk --no-cache add msttcorefonts-installer fontconfig && \
    update-ms-fonts && \
    fc-cache -f

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /generator

CMD [ "/generator" ]
