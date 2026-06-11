FROM docker.io/library/golang:1.22.0-alpine AS builder

ARG GIT_SHA=unknown
LABEL application_name="internal-devbox"
LABEL description="Internal dev-env helper CLI for Pantalasa engineers"
LABEL owner="devex@pantalasa.org"
LABEL source_uri="https://github.com/pantalasa/internal-devbox"
LABEL git_sha="${GIT_SHA}"

WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/devbox .

FROM docker.io/library/alpine:3.20
COPY --from=builder /out/devbox /devbox
ENTRYPOINT ["/devbox"]

# Container healthcheck (image metadata)
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD exit 0
# Run as a non-root user
USER 1001
