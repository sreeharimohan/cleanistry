FROM alpine:latest

ADD cleanistry /usr/bin/cleanistry
RUN chmod +x /usr/bin/cleanistry
CMD ["/usr/bin/cleanistry"]
