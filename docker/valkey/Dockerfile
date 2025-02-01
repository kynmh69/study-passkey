FROM valkey/valkey:8.0
LABEL authors="hirokikayanuma"
COPY ./docker/valkey/valkey.conf /etc/valkey/valkey.conf
CMD ["valkey-server", "/etc/valkey/valkey.conf"]
