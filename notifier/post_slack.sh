#!/bin/sh

VIDEO_URL=$(cat data/video_url.txt)

curl -X POST $SLACK_WEBHOOK_URL \
    -H "Content-type: application/json" \
    --data '{ "text": "<'"${VIDEO_URL}"'|'"${SLACK_POST_TITLE}"'>\n'"${SLACK_POST_BODY}"'" }'
