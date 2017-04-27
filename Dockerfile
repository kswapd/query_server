FROM debian:jessie
#FROM scratch
MAINTAINER xw.kong@hnair.com
ENV INFLUXDB_HOST "54.223.149.26:8086"
ENV ES_HOST "http://172.16.10.169:9200"
COPY ./bin/monitor-query-server /usr/bin/
RUN chmod 777 /usr/bin/monitor-query-server
EXPOSE 8100
ENTRYPOINT /usr/bin/monitor-query-server -influxdb_driver_host="${INFLUXDB_HOST}" -elasticsearch_cluster_host="${ES_HOST}"



