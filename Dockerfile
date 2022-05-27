#-- Stage: Builder
FROM golang:1.18-alpine as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /go/src/gitlab.com/zerok/zerokspot.com/
WORKDIR /go/src/gitlab.com/zerok/zerokspot.com/cmd/blog
RUN mkdir -p /opt/bin && \
    go build -o /opt/bin/blog

#-- Stage: Final result
FROM alpine:3.16
RUN apk add --no-cache git imagemagick
COPY --from=builder /opt/bin/blog /usr/local/bin/
