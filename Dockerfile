FROM golang:latest

ENV APP_HOME /app

RUN mkdir $APP_HOME
WORKDIR $APP_HOME

COPY . $APP_HOME%