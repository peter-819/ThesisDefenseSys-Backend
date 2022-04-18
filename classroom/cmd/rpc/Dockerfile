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
COPY classroom/cmd/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/classroom classroom/cmd/rpc/classroom.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/classroom /app/classroom
COPY --from=builder /app/etc /app/etc

CMD ["./classroom", "-f", "etc/classroom.yaml"]
