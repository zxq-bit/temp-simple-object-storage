FROM cargo.caicloudprivatetest.com/caicloud/debian:jessie

# Set the default timezone to Shanghai
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

WORKDIR /caicloud/simple-object-storage
COPY ./bin/gateway /caicloud/simple-object-storage/

EXPOSE 80

CMD ["/caicloud/simple-object-storage/gateway"]
