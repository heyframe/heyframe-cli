ARG PHP_VERSION

FROM ghcr.io/heyframe/heyframe-cli-base:${PHP_VERSION}

ARG TARGETPLATFORM

COPY $TARGETPLATFORM/heyFrame-cli /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/heyFrame-cli"]
CMD ["--help"]
