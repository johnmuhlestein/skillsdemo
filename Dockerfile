# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY --chown=app:app go.mod go.sum ./
RUN go mod download
COPY --chown=app:app cmd cmd
COPY --chown=app:app ent ent
COPY --chown=app:app internal internal
COPY --chown=app:app rpc rpc

RUN go build -C cmd/server -o /skillsdemo

FROM scratch
WORKDIR /
COPY --from=build /skillsdemo /skillsdemo
EXPOSE 8080

ENTRYPOINT ["/skillsdemo"]