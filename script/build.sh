GOOS=linux go build -o bin/monitor-query-server
docker build -t registry.hnaresearch.com/public/monitor-query-server:1.1 .
#docker push registry.hnaresearch.com/public/monitor-query-server:1.0
