GOOS=linux go build -o bin/monitor-query-server
docker build -t registry.hnaresearch.com/tang/monitor-query-server:0.2 .
docker push registry.hnaresearch.com/tang/monitor-query-server:0.2
