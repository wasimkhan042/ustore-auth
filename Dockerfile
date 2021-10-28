FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/ustore-auth /usr/local/bin/ustore-auth

CMD ["/usr/local/bin/ustore-auth"]