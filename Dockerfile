FROM alpine:latest

COPY ./hc2017 /root

CMD ["/root/hc2017"]

EXPOSE 80