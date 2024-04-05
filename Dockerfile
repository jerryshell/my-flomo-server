FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --update add tzdata && \
    rm -rf /var/cache/apk/*

ARG TARGETOS

ARG TARGETARCH

ENV BINARY_NAME=my-flomo-server-${TARGETARCH}-${TARGETOS}

COPY ./${BINARY_NAME} /my-flomo-server/${BINARY_NAME}

EXPOSE 8060

ENV GIN_MODE=release

WORKDIR /my-flomo-server

CMD ["sh", "-c", "./${BINARY_NAME}"]
