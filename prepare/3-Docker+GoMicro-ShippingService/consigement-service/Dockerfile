# 若运行环境是Linux则需要把alpine换成debian
# 使用最新版的alpine作为基础镜像
FROM alpine:latest

# 在容器的根目录下创建app目录
RUN mkdir /app

# 将工作目录切换到 /app下
WORKDIR /app

# 将微服务的服务端运行文件拷贝到 /app下
ADD consignment-service /app/consignment-service

# 运行服务端
CMD ["./consignment-service"]