FROM node:alpine
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY .env.example .env
COPY ./ ./
RUN npm i sqlite3
RUN npm i
CMD ["npm", "run", "start"]