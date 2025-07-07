FROM alpine:3.16

# dumb-init needed for injected secure bootstrapping entrypoint script when run in secure mode.
RUN apk add --update --no-cache zeromq dumb-init

WORKDIR /app
COPY storage /app/
COPY docker/storage/res /app/res
COPY cmd/storage/assets/swagger /app/assets/swagger


EXPOSE 9090
EXPOSE 9091

ENTRYPOINT ["/app/storage"]
CMD ["--conf=/app/res/configuration.yaml"]
