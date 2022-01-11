FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    rm -rf /var/cache/apk/*
COPY ./my-flomo-server-linux-amd64 /my-flomo-server-linux-amd64
COPY ./config.json /config.json
EXPOSE 8060
CMD ["/my-flomo-server-linux-amd64"]