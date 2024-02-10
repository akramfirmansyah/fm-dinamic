FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/fm-dinamic-api ./...

CMD [ "fm-dinamic-api" ]