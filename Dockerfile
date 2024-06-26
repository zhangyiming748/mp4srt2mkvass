FROM golang:1.22.2-bookworm
ARG proxy="http://192.168.1.20:8889"
# 编译可执行文件
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir /root/app
WORKDIR /root/app
COPY . .
RUN CGO_ENABLED=1 go build -o /usr/local/bin/MergeSubtitle main.go

# 安装必要软件
RUN mv ./debian.sources /etc/apt/sources.list.d
RUN apt update

COPY install-retry.sh /usr/local/bin/install-retry.sh
RUN chmod +x /usr/local/bin/install-retry.sh

RUN install-retry.sh ffmpeg locales dos2unix mediainfo

RUN locale-gen zh_CN.UTF-8
RUN echo "export LC_ALL=zh_CN.UTF-8">> /etc/profile

RUN dos2unix entrypoint.sh
RUN chmod +x entrypoint.sh
CMD ["MergeSubtitle"]
# ENTRYPOINT ["/root/app/entrypoint.sh"]
# docker build -t merge:latest .
# docker run -dit --name=merge --cpus=1 --memory=2048M -v '/c/Users/zen/Videos/test/data:/data' merge:latest
# docker run -dit --name=merge --cpus=1 --memory=2048M -v '/c/Users/zen/Videos/test/data:/data' merge:latest bash