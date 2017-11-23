FROM alpine:3.6

ADD cleanistry /usr/bin/cleanistry
ARG DOCKER_CLI_VERSION="17.06.2-ce"
ENV DOWNLOAD_URL="https://download.docker.com/linux/static/stable/x86_64/docker-$DOCKER_CLI_VERSION.tgz"
RUN apk --update add curl \
    && mkdir -p /tmp/download \
    && curl -L $DOWNLOAD_URL | tar -xz -C /tmp/download \
    && mv /tmp/download/docker/docker /usr/local/bin/ \
    && rm -rf /tmp/download \
    && apk del curl \
    && rm -rf /var/cache/apk/* \
    && chmod +x /usr/bin/cleanistry
CMD ["/usr/bin/cleanistry"]
