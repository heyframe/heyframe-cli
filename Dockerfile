ARG PHP_VERSION

FROM ghcr.io/heyframe/heyframe-cli-base:${PHP_VERSION}

ARG TARGETPLATFORM

COPY $TARGETPLATFORM/heyframe-cli /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/heyframe-cli"]
CMD ["--help"]
