FROM golang:alpine as builder

MAINTAINER shivamxav@gmail.com

WORKDIR /go/src/lilagames/app

COPY app /go/src/lilagames/app

# Run the two commands below to install git and dependencies for the project.
RUN apk update && apk add --no-cache git
RUN go install

RUN go build .

FROM alpine as final
ENV PORT=3004
ENV GIN_MODE=release

COPY --from=builder /go/src/lilagames/app/LilaGames .
EXPOSE $PORT

ENTRYPOINT ["./LilaGames"]