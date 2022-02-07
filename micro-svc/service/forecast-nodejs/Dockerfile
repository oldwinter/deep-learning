FROM node:11-slim

COPY . /app
WORKDIR /app
RUN npm i -g curl \
    && npm i
EXPOSE 3002
CMD ["node", "main.js"]