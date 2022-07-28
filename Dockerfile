FROM golang:1.18-alpine

LABEL maintainer="Injila <antonyshikubu@gmail.com>"

WORKDIR /src

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8000

RUN go build 

CMD [ "./vision" ]
