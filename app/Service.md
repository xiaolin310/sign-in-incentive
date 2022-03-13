## Service

#### user， RPC: 9001， 依赖：mysql [redis]

#### record, RPC: 9003, 依赖：mysql, redis, kafka

#### virtualwallet, RPC: 9002, 依赖：mysql


### signin/interface，RPC：9000， HTTP：8000

### signin/admin， RPC： 9004，HTTP：8001

### reward, Job, 不占用端口， 依赖：Kafka，wallet