# Kafka

- 先安装zookeeper

```docker
docker run -d --name zookeeper -p 2181:2181 -v /etc/localtime:/etc/localtime zookeeper
```

- 再安装kafka

```docker
docker run  -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=192.168.206.1:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.206.1:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -t bitnami/kafka
```

- 进入容器内部

```bash
# exec kafka
docker exec -it 2b009cbc32dc /bin/sh

# inside
cd /opt/bitnami/kafka/bin 

# create topic
./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic obudata

# start producer
./kafka-console-producer.sh --broker-list localhost:9092 --topic obudata

# consume the producer
./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --from-beginning --topic obudata

# how to delete the producer message?
# so confuse
```

## Kafka 操作指南

安装kafka driver

```bash
go get -u  github.com/confluentinc/confluent-kafka-go/v2/kafka
```

连接kafka

Producer 生产者

```go
import (
 "fmt"
 "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

 p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
 if err != nil {
  panic(err)
 }

 defer p.Close()

 // Delivery report handler for produced messages
 go func() {
  for e := range p.Events() {
   switch ev := e.(type) {
   case *kafka.Message:
    if ev.TopicPartition.Error != nil {
     fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
    } else {
     fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
    }
   }
  }
 }()

 // Produce messages to topic (asynchronously)
 topic := "myTopic"
 for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
  p.Produce(&kafka.Message{
   TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
   Value:          []byte(word),
  }, nil)
 }

 // Wait for message deliveries before shutting down
 p.Flush(15 * 1000)
}

```

消费者

```go
import (
 "fmt"
 "time"

 "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

 c, err := kafka.NewConsumer(&kafka.ConfigMap{
  "bootstrap.servers": "localhost",
  "group.id":          "myGroup",
  "auto.offset.reset": "earliest",
 })

 if err != nil {
  panic(err)
 }

 c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

 // A signal handler or similar could be used to set this to false to break the loop.
 run := true

 for run {
  msg, err := c.ReadMessage(time.Second)
  if err == nil {
   fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
  } else if !err.(kafka.Error).IsTimeout() {
   // The client will automatically try to recover from all errors.
   // Timeout is not considered an error because it is raised by
   // ReadMessage in absence of messages.
   fmt.Printf("Consumer error: %v (%v)\n", err, msg)
  }
 }

 c.Close()
}
```
