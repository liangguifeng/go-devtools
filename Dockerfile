FROM golang:1.17-alpine3.16 AS builder

WORKDIR /go/src/github.com/liangguifeng/go-devtools

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
ADD . /go/src/github.com/liangguifeng/go-devtools
RUN go mod tidy && go mod vendor

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o go-devtools .

FROM alpine:3.10 AS final

WORKDIR /app

COPY --from=builder /go/src/github.com/liangguifeng/go-devtools/go-devtools /app/

CMD go-devtools
