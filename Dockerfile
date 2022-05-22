FROM golang:alpine

WORKDIR /src
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build

ENTRYPOINT [ "./stranger-album-api" ]


