FROM golang:1.17.2-alpine AS build

RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

RUN mkdir /app

COPY ./code /app
RUN chmod +x /app/entrypoint.sh
RUN chmod -R 755 /app/config
WORKDIR /app


RUN go mod download
RUN go mod verify

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/server .
COPY --from=build /app/entrypoint.sh .
COPY --from=build /app/config/*.yml ./config/
EXPOSE 80
ENTRYPOINT ["./entrypoint.sh"]