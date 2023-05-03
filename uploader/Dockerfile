FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./uploader

# ---

FROM alpine:3.17

WORKDIR /app

COPY --from=build /app/uploader ./uploader

COPY .credentials .credentials

CMD [ "./uploader" ]
