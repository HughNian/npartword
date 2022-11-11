FROM ubuntu

WORKDIR /root

ARG TARGETARCH

COPY server/data ./data
COPY server/npwserver ./npwserver

# 镜像启动服务自动被拉起配置
COPY run /etc/service/run
RUN chmod +x /etc/service/run