FROM qimao-registry.cn-beijing.cr.aliyuncs.com/public/alpine:latest
ADD ./bin /bin
ENTRYPOINT [ "/bin/server" ]