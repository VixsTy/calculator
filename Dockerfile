# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

## -------------------------------------------------------------------------------------------------

FROM golang:1.12 as tools
RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx
WORKDIR /tmp/src/calculator
# Force go modules
ENV GO111MODULE=on
COPY tools tools
RUN cd tools && go run github.com/magefile/mage

## -------------------------------------------------------------------------------------------------

FROM tools AS deps
COPY magefile.go go.mod ./
RUN set -eux; \
    go run github.com/magefile/mage go:deps

FROM deps as source
COPY . .

## -------------------------------------------------------------------------------------------------

FROM source AS build
RUN go run github.com/magefile/mage
# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/base:latest
WORKDIR /
# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF
# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="Calculator" \
    org.label-schema.description="Calculator command line" \
    org.label-schema.url="https://github.com/VixsTy/calculator" \
    org.label-schema.vcs-url="https://github.com/VixsTy/calculator.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Kevin LARQUEMIN" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0" \
    org.zenithar.licence="MIT"
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /tmp/src/calculator/bin/* /app/bin
ENTRYPOINT [ "/app/bin" ]
CMD ["--help"]



