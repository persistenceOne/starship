FROM golang:1.18-alpine3.16 AS go-builder

RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0

FROM alpine:3.16

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

COPY --from=go-builder /go/bin/cosmovisor /usr/bin/

# Set up dependencies
ENV PACKAGES --update ca-certificates curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
