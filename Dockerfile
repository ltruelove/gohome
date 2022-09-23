FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/gohome
RUN mkdir -p "${APP_HOME}"

WORKDIR "${APP_HOME}"
EXPOSE 8082

CMD ["bee", "run"]