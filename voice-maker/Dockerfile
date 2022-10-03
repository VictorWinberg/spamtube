FROM golang:1.18-alpine

WORKDIR /app

# install dependencies
RUN apk add --no-cache pkgconfig alsa-lib-dev gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /voice-maker

CMD [ "/voice-maker" ]
