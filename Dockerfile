FROM alpine:3.18

RUN apk add gcompat

WORKDIR /service

COPY ./build/database ./

RUN chmod +x /service/database

CMD ["/service/database"]