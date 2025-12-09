FROM debian:bookworm-slim

WORKDIR /usr/local/sectfs-server

RUN set -e && mkdir config static

COPY ./bin/sectfs ./

COPY ./config/sectfs.conf ./config/

COPY ./static/index.html ./static/

RUN set -e && chmod +x ./sectfs

EXPOSE 5363

CMD ["./sectfs"]
