FROM node:12.7-stretch

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/
RUN yarn
COPY . /usr/src/app
RUN yarn build

EXPOSE 4200

CMD [ "yarn", "start" ]
