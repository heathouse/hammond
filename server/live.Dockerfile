FROM golang:1.23
#FROM golang:1.15

ARG UID=1000
ARG GID=1000

ENV CONFIG=/config
ENV DATA=/assets

USER root
RUN go install github.com/air-verse/air@latest

RUN adduser --uid ${UID} --gecos "" --disabled-password golang
RUN groupmod -g $GID golang
RUN chown -R root:golang /go/
RUN chmod -R g+rwx /go/

USER golang
WORKDIR /app

EXPOSE 3000

RUN mkdir /home/golang/airtmp
CMD ["air"]
