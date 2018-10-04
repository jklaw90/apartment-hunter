FROM alpine:latest
RUN apk add --update ca-certificates
WORKDIR /
ADD ./bin/writer /writer
RUN chmod 755 /writer
ENTRYPOINT ["/writer"]
