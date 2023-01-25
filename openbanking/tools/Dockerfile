FROM docker.io/golang:1.18-alpine

LABEL maintainer="Mohamed Bana <mohamed@bana.io>"

ENV TERM xterm-color
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN apk --no-cache add ca-certificates git bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN set -e -o pipefail -x; \
  go build -ldflags='-w -s' -installsuffix 'static' -o ./openbanking_tools ./main.go && \
  ln -s $(pwd)/openbanking_tools /usr/bin/openbanking_tools

ENTRYPOINT ["/usr/bin/openbanking_tools"]
CMD ["--help"]
