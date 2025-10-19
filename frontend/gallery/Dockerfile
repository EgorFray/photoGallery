FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm i -g serve
COPY . .
RUN npm run build
RUN npm install -g serve
CMD [ "serve", "-s", "dist" ]