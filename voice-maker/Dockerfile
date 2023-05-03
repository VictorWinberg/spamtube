FROM golang:1.18-alpine AS build

WORKDIR /app

# install dependencies
RUN apk add --no-cache pkgconfig alsa-lib-dev gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./voice-maker

# ---

FROM alpine:3.17

WORKDIR /app

# install dependencies
RUN apk add --no-cache pkgconfig alsa-lib-dev gcc libc-dev

COPY --from=build /app/voice-maker ./voice-maker

CMD [ "./voice-maker" ]
