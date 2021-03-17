FROM golang:alpine AS build-env
RUN apk add build-base

WORKDIR $HOME/app

ENV CGO_ENABLED 1

RUN export PATH="$PATH:$GOPATH/bin"

COPY . .

RUN go build

FROM alpine
WORKDIR /
COPY --from=build-env /app/Blibots /Blibots
ENV DISCORD_TOKEN=DISCORD_TOKEN_HERE
ENTRYPOINT ./Blibots