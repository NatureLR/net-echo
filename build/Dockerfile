# 编译镜像 golang:x.y.z-alpine3.13
ARG BUILD_IMAGE=golang:1.21-alpine
ARG RUN_IMAGE=alpine:3

# 编译镜像
FROM ${BUILD_IMAGE} as build

ENV ROOT_DIR=/build
WORKDIR /build

COPY . .

# 国内替换镜像源
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \ 

# 修改源为国内阿里
# 修改时区为上海
# 安装make和git工具
RUN  apk add --no-cache ca-certificates tzdata  && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk add make && \
    apk add git

# 国内使用的goproxy
#ENV GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

RUN make build-in-docker

# 运行镜像
FROM ${RUN_IMAGE}

WORKDIR /root/

# RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list.d/debian.sources
RUN apt update && \
    apt install -y \
    telnet \
    dnsutils \
    iputils-ping \
    iproute2 \
    wget \
    curl

COPY --from=build /build/net-echo .

EXPOSE 80

#ENTRYPOINT ["./net-echo"]

CMD ["./net-echo"]
