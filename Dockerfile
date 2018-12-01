#-- Stage: Builder
FROM golang:1.11.1-alpine as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /go/src/gitlab.com/zerok/zerokspot.com/
WORKDIR /go/src/gitlab.com/zerok/zerokspot.com/cmd/blogsearch
RUN mkdir -p /opt/bin && \
    go build -o /opt/bin/blogsearch

#-- Stage: Final result
FROM alpine:3.8
RUN apk add --no-cache git
COPY --from=builder /opt/bin/blogsearch /usr/local/bin/
