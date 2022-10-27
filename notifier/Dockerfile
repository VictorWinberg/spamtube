FROM alpine:3.14

WORKDIR /app

RUN apk --no-cache add curl

COPY . ./

CMD ["sh", "post_slack.sh"]
