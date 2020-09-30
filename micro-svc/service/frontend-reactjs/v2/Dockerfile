FROM node:11-slim

COPY . /app
WORKDIR /app
RUN npm i -g curl \
    && npm i \
    && npm run build \
    && rm -rf node_modules \
    && cd scripts \
    && npm i \
    && cd -
EXPOSE 3000
CMD ["node", "scripts/server.js"]