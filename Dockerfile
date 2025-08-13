FROM alpine:latest

COPY ./out/bgptfrm /app/bgptfrm

WORKDIR /app

CMD ["/app/bgptfrm"]