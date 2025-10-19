FROM node:18-alpine
# Define a build argument
ARG VITE_BACKEND_URL

# Set it as an environment variable (Vite needs this at build time)
ENV VITE_BACKEND_URL=${VITE_BACKEND_URL}
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm i -g serve
COPY . .
RUN npm run build
RUN npm install -g serve
CMD [ "serve", "-s", "dist" ]