FROM golang:1.19

ARG arch=amd64

WORKDIR /

COPY powerdns-remote-backend-linux-${arch} dabdns.yaml /

RUN mv /powerdns-remote-backend-linux-${arch} /dabdns

EXPOSE 5353

CMD [ "/dabdns" ]