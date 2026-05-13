FROM node:24-alpine
WORKDIR /app
COPY package.json ./
RUN npm install
COPY send_sepolia_message.mts poll_boundless_proofs.mts generate_fixture.mts ./
RUN mkdir -p /app/testdata
