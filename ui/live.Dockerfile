FROM node:16

ARG UID=1000
ARG GID=1000

USER root
RUN usermod -u $UID node
RUN groupmod -g $GID node
RUN npm install -g nodemon

USER node
WORKDIR /app

EXPOSE 5000

CMD ["nodemon", "-w", "./src/", "--exec", "vue-cli-service", "serve", "--port", "5000"]
