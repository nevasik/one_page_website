FROM golang:1.22.0 as gobuild

WORKDIR /app

RUN apt-get -y install tzdata
ENV TZ=Europe/Moscow

COPY . .
RUN go mod download
RUN go build -o main ./cmd

CMD [ "./main"]
