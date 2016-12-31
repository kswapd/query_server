GOOS=linux go build -o bin/monitor-query-server
docker build -t mutemaniac/monitor-query-server:0.1 .
docker push mutemaniac/monitor-query-server:0.1
