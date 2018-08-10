# BUILDER
FROM golang:1.10 AS builder
ARG SERVICE=hue-tools
WORKDIR /go/src/github.com/redkite1/$SERVICE
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/$SERVICE

# RUNNER
FROM alpine:3.8
WORKDIR /usr/local/bin
ARG SERVICE=hue-tools
COPY --from=builder /go/bin/$SERVICE .

# it does accept the variable $SERVICE
CMD ["hue-tools"]