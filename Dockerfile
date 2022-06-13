# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ARG SENDGRID_API_KEY

ENV SENDGRID_API_KEY=${SENDGRID_API_KEY}

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /exec ./src/

EXPOSE 5896

CMD [ "/exec" ]