# BUILDER
FROM golang:1.12 AS builder
ARG SERVICE=hue-tools
WORKDIR /opt/$SERVICE
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o /opt/$SERVICE/$SERVICE

# RUNNER
FROM alpine:3.8
WORKDIR /usr/local/bin
ARG SERVICE=hue-tools
COPY --from=builder /opt/$SERVICE/$SERVICE .

# it does accept the variable $SERVICE
CMD ["hue-tools"]
