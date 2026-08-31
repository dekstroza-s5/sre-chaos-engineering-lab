FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY app ./app
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/demo ./app

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /out/demo /demo
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/demo"]
