FROM alpine:latest

RUN mkdir /app

COPY connectApp /app

CMD [ "/app/connectApp"]