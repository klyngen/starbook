FROM golang:1.18
#MAINTAINER Martin Klingenberg <klingen@alv.no>

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -v -o app
RUN go install github.com/mitranim/gow@latest

ENV DATABASE_NAME="starbook"
ENV DATABASE_USERNAME="starbook"
ENV DATABASE_PASSWORD="password"
ENV DATABASE_HOST="localhost"

EXPOSE 8080

CMD ["go", "run", "."]
