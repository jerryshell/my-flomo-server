FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --update add tzdata && \
    rm -rf /var/cache/apk/*
COPY ./my-flomo-server-amd64-linux /my-flomo-server/my-flomo-server-amd64-linux
COPY ./config.json /my-flomo-server/config.json
EXPOSE 8060
ENV GIN_MODE=release
WORKDIR /my-flomo-server
CMD ["./my-flomo-server-amd64-linux"]
