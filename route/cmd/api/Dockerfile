FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY route/cmd/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/route route/cmd/api/route.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/route /app/route
COPY --from=builder /app/etc /app/etc

CMD ["./route", "-f", "etc/route-api.yaml"]
