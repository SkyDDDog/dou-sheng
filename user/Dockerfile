FROM alpine:3.12

MAINTAINER Sky_Dog<362664609@qq.com>

RUN mkdir "/app"
WORKDIR "/app"

COPY main /app/app
RUN chmod +x /app/app
ENTRYPOINT ["./app"]