FROM golang:1.14-alpine as builder

ARG VERSION

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY version/ version/
COPY collector/ collector/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags="-w -s -X github.com/scnewma/ynab-exporter/version.Version=${VERSION}" -o ynab-exporter main.go


FROM gcr.io/distroless/static:nonroot

ARG VCS_REF
ARG VERSION
ARG BUILD_DATE

LABEL org.opencontainers.image.created=${BUILD_DATE} \
      org.opencontainers.image.authors="github:@scnewma" \
      org.opencontainers.image.source="https://github.com/scnewma/ynab-exporter" \
      org.opencontainers.image.version=${VERSION} \
      org.opencontainers.image.revision=${VCS_REF} \
      docker.cmd="docker run -e YNAB_TOKEN=token -p 9721:9721 scnewma/ynab-exporter"


WORKDIR /
COPY --from=builder /workspace/ynab-exporter .
USER nonroot:nonroot

ENTRYPOINT ["/ynab-exporter"]