FROM alpine:latest
COPY ./my-flomo-server-linux-amd64 /my-flomo-server-linux-amd64
COPY ./config.json ./config.json
EXPOSE 8060
CMD ["/my-flomo-server-linux-amd64"]