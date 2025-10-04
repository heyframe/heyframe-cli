ARG PHP_VERSION

FROM ghcr.io/heyframe/heyframe-cli-base:${PHP_VERSION}


COPY {{ .Binary }} /usr/local/bin/heyframe-cli

ENTRYPOINT ["/usr/local/bin/heyframe-cli"]
CMD ["--help"]
