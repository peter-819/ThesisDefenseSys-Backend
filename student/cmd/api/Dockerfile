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
COPY student/cmd/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/student student/cmd/api/student.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/student /app/student
COPY --from=builder /app/etc /app/etc

CMD ["./student", "-f", "etc/student-api.yaml"]
