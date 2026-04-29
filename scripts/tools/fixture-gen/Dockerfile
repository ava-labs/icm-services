FROM node:24-alpine
WORKDIR /app
COPY package.json ./
RUN npm install
COPY send_message.mts prepare_fixtures.mts generate_fixtures.mts ./
RUN mkdir -p /app/testdata
