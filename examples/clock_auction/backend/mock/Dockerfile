FROM node:12.13.0-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH

COPY serverapi.js /usr/src/app/serverapi.json

RUN npm install -g json-server

CMD ["json-server", "--watch", "serverapi.json", "--host", "0.0.0.0"]
