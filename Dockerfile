FROM golang:alpine

WORKDIR /src
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build

EXPOSE 8080
ENTRYPOINT [ "./stranger-album-api" ]


