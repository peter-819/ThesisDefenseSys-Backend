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
COPY excel/cmd/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/excel excel/cmd/rpc/excel.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/excel /app/excel
COPY --from=builder /app/etc /app/etc

CMD ["./excel", "-f", "etc/excel.yaml"]
