FROM node:boron

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json /usr/src/app/
RUN npm install

# Copy node modules for hopeful reuse
COPY ./node_modules /usr/src/app/node_modules

# Bundle app source
COPY . /usr/src/app/

EXPOSE 8080
CMD [ "npm", "start" ]

ARG VERSION
ARG BUILD_DATE
ENV VERSION $VERSION
ENV BUILD_DATE $BUILD_DATE