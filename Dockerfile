ARG PHP_VERSION
FROM ghcr.io/heyframe/heyframe-cli-base:${PHP_VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY ${TARGETOS}/${TARGETARCH}/heyframe-cli /usr/local/bin/heyframe-cli

ENTRYPOINT ["/usr/local/bin/heyframe-cli"]
CMD ["--help"]