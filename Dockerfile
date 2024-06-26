FROM docker.io/golang:1.22.1-alpine AS build

RUN apk add --no-cache upx
ENV PORT=8080

WORKDIR /src
COPY backends/go/go.* .
RUN go mod download
COPY backends/go/. .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/site cmd/site/main.go
RUN upx /out/site

FROM scratch
COPY --from=build /out/site /
ENTRYPOINT ["/site"]