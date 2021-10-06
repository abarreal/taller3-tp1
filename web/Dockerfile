FROM node:16

WORKDIR /usr/src/app

# Copy package*.json to install dependencies before copying the rest of the code.
COPY package*.json ./
# Do install dependencies.
RUN npm install

# Copy everything else.
COPY . .

ENTRYPOINT ["node", "index.js"]