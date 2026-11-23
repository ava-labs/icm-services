FROM node:24-alpine
WORKDIR /app
COPY package.json ./
RUN npm install
COPY send_ethereum_message.mts wait_for_finality.mts generate_fixture.mts ./
RUN mkdir -p /app/testdata
