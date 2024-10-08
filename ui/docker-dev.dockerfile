# https://github.com/cypress-io/cypress-docker-images/tree/master/base
FROM cypress/base:10.16.0

# Make the `app` folder the current working directory
WORKDIR /app

# Copy dependency-related files
COPY package.json ./

# Install project dependencies
RUN npm install

# Expose ports 8080, which the dev server will be bound to
EXPOSE 8080
