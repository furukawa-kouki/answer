FROM node:15.7.0-alpine3.12
WORKDIR /usr/app/fe
COPY package.json ./
COPY yarn.lock ./
RUN yarn install
COPY ./ ./
