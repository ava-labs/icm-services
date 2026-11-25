FROM node:24-alpine
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci --ignore-scripts
COPY send_ethereum_message.mts poll_boundless_proofs.mts generate_fixture.mts ./
RUN mkdir -p /app/testdata
