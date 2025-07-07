FROM alpine:3.16

# dumb-init needed for injected secure bootstrapping entrypoint script when run in secure mode.
RUN apk add --update --no-cache zeromq dumb-init

WORKDIR /app
COPY data /app/
COPY docker/data/res /app/res
COPY cmd/data/assets/swagger /app/assets/swagger


EXPOSE 9090
EXPOSE 9091

ENTRYPOINT ["/app/data"]
CMD ["--conf=/app/res/configuration.yaml"]
