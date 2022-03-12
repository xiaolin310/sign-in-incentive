package job

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

// 生产一条kafka消息
func TestDoReward(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"172.19.92.170:9092"}, config)
	defer client.Close()
	if err != nil {
		t.Error(err.Error())
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = "reward"
	msg.Value= sarama.ByteEncoder("1-150.00")

	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Producer message sent to Kafka, partition:%v offset:%v\n", partition, offset)

}
