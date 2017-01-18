FROM alpine
#FROM scratch
MAINTAINER qian.tang@hnair.com

#ENV INFLUXDB_HOST "54.223.149.26:8086"

# Grab cadvisor from the staging directory.
#ADD monitor-query-server /usr/bin/monitor-query-server
COPY ./bin/monitor-query-server /usr/bin/
RUN chmod 777 /usr/bin/monitor-query-server
EXPOSE 8100
#ENTRYPOINT ["/usr/bin/monitor-query-server -influxdb_driver_host=${INFLUXDB_HOST}"]
#ENTRYPOINT /usr/bin/monitor-query-server -influxdb_driver_host="${INFLUXDB_HOST}"
#ENTRYPOINT ["/usr/bin/monitor-query-server","-influxdb_driver_host=${INFLUXDB_HOST}"]
ENTRYPOINT ["/usr/bin/monitor-query-server"]
#CMD ["/usr/bin/monitor-query-server"]

