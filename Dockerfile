FROM alpine:3.9

LABEL maintainer="mritd <mritd1234@gmail.com>"

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

RUN apk upgrade \
    && apk add tar wget bash tzdata libc6-compat ca-certificates \
    && wget http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz \
    && tar -zxvf GeoLite2-City.tar.gz -C /tmp --strip-components 1 \
    && mv /tmp/GeoLite2-City.mmdb /geoip.mmdb \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tar wget \
    && rm -rf GeoLite2-City.tar.gz /var/cache/apk/* /tmp/*

COPY dist/myip_linux_amd64 /usr/bin/myip

CMD ["myip","--db","/geoip.mmdb"]
