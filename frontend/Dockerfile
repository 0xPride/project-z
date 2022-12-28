FROM node:current-alpine3.17 AS BUILDER
# RUN apt update
# RUN apt install nodejs npm ca-certificates -y
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm i
COPY . .
RUN npm run build
CMD ["node", "prod/index.js"]