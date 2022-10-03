FROM node:14-alpine

WORKDIR /app

# install dependencies
RUN apk add g++ make py3-pip

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]
