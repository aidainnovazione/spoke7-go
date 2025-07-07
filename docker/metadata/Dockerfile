FROM alpine:3.16

# dumb-init needed for injected secure bootstrapping entrypoint script when run in secure mode.
RUN apk add --update --no-cache zeromq dumb-init

WORKDIR /app
COPY metadata /app/
COPY docker/metadata/res /app/res
COPY cmd/metadata/assets/swagger /app/assets/swagger


EXPOSE 9090
EXPOSE 9091

ENTRYPOINT ["/app/metadata"]
CMD ["--conf=/app/res/configuration.yaml"]
