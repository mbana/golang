FROM golang:1.15-alpine3.13 AS builder

# Use an unprivileged user. Create the user and group files that will be used in the running container to
# run the process an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' >/user/passwd && \
    echo 'nobody:x:65534:' >/user/group

WORKDIR /go/src/project_template_v2

ENV TERM=xterm-256color
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum /go/src/project_template_v2/
RUN go mod download
RUN go mod verify
COPY . /go/src/project_template_v2

RUN apk add --update --no-cache ca-certificates bash make git
RUN make build-binary-docker

FROM scratch

COPY --from=builder /go/src/project_template_v2/main /usr/bin/project_template_v2
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Use an unprivileged user. Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

USER nobody:nobody

EXPOSE 8080
EXPOSE 8443

ENTRYPOINT ["/usr/bin/project_template_v2"]
