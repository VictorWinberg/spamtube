FROM node:14-alpine

WORKDIR /app

RUN apk add curl

COPY package*.json ./

RUN npm ci

COPY . .

CMD ["npm", "start"]
