FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /usr/bin

COPY evermos-project evermos-project

COPY .env .env

EXPOSE 8000

CMD [ "evermos-project" ]