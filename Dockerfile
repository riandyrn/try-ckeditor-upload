FROM golang:1.19.0-alpine3.16 as build

WORKDIR /go/src/github.com/riandyrn/try-ckeditor-upload

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3.16
COPY --from=build /go/src/github.com/riandyrn/try-ckeditor-upload/app ./app
COPY ./web ./web
COPY ./images ./images

CMD "./app"