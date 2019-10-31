FROM node:12.13.0-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH

COPY serverapi.js /usr/src/app/serverapi.js

RUN npm install express

CMD ["node", "serverapi.js"]
