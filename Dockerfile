FROM  golang:1.16.1-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /root

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /dmos -ldflags "-s -w -X github.com/linuxsuren/cobra-extension/version.version=latest  -X github.com/linuxsuren/cobra-extension/version.commit= -X github.com/linuxsuren/cobra-extension/version.date=" main.go

FROM alpine AS UPX
COPY --from=builder /dmos /dmos
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk add --update upx && upx /dmos
FROM alpine
COPY --from=UPX /dmos /bin/dmos
ENTRYPOINT ["/bin/dmos"]
