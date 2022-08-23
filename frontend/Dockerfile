FROM node:14-alpine

WORKDIR /app

COPY package*.json ./

RUN apk add g++ make py3-pip

RUN npm ci

COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]
