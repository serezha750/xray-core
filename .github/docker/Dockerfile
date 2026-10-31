# syntax=docker/dockerfile:latest
FROM --platform=$BUILDPLATFORM golang:1.25-bookworm AS build

WORKDIR /src
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -o xray -trimpath -buildvcs=false \
    -ldflags "-s -w -buildid=" ./main

RUN mkdir -p /tmp/empty /tmp/usr/local/etc/xray /tmp/var/log/xray
RUN echo '{}' > /tmp/usr/local/etc/xray/config.json
RUN touch /tmp/var/log/xray/error.log

# 最终镜像使用 distroless，极小
FROM gcr.io/distroless/static:nonroot

COPY --from=build --chown=0:0 --chmod=755 /src/xray /usr/local/bin/xray
COPY --from=build --chown=65532:65532 --chmod=755 /tmp/empty /usr/local/share/xray
COPY --from=build --chown=0:0 --chmod=755 /tmp/empty /usr/local/etc/xray
COPY --from=build --chown=0:0 --chmod=644 /tmp/usr/local/etc/xray/config.json /usr/local/etc/xray/
COPY --from=build --chown=65532:65532 --chmod=600 /tmp/var/log/xray/error.log /var/log/xray/

VOLUME /usr/local/etc/xray
VOLUME /var/log/xray

ENTRYPOINT ["/usr/local/bin/xray"]
CMD ["run", "-c", "/usr/local/etc/xray/config.json"]
