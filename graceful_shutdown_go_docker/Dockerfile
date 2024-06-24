#=============== Build Stage ===============
FROM golang:1.22 as builder
WORKDIR /demo
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o /demo/demo

#=============== Final Stage ===============
FROM gcr.io/distroless/static:latest as final
COPY --from=builder /demo/demo /demo
ENTRYPOINT [ "/demo" ]
