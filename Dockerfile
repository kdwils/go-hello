FROM golang:1.22.3-alpine3.20 as builder

ARG TARGETOS
ARG TARGETARCH
ENV GOARCH ${TARGETARCH:-amd64}
ENV GOOS ${TARGETOS:-linux}

WORKDIR /app

COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 go build -o go-hello .

FROM alpine
COPY --from=builder /app/go-hello .

ENTRYPOINT [ "./go-hello" ]