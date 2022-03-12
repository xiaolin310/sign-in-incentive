package main

import (
    "context"
    "fmt"

    "github.com/Shopify/sarama"
)

var Consumer sarama.Consumer

func main() {
    var err error
    Consumer, err = sarama.NewConsumer([]string{"172.19.92.170:9092"}, nil)
    if err != nil {
        fmt.Printf("fail to start consumer,err:%v\n", err)
        return
    }
    // topics := []string{"mytopic"}
    //topics := []string{"nsq"}
    topic := "mytopic"
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    err = testTopic(topic, ctx)
    if err != nil {
        fmt.Printf("consume message failed")
    }
    select {}
}
func testTopic(topic string, ctx context.Context) error {
    partitionList, err := Consumer.Partitions(topic)
    fmt.Println(partitionList)
    if err != nil {
        fmt.Printf("fail to start consumer partition,err:%v\n", err)
        return err
    }
    for partition := range partitionList {
        //  遍历所有的分区，并且针对每一个分区建立对应的消费者
        pc, err := Consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
        if err != nil {
            fmt.Printf("fail to start consumer for partition %d,err:%v\n", partition, err)
            return err
        }
        go func(sarama.PartitionConsumer) {
            defer pc.AsyncClose()
            for msg := range pc.Messages() {
                fmt.Printf("Partition:%d Offset:%v Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
                select {
                case <-ctx.Done():
                    return
                default:
                    continue
                }
            }
        }(pc)
    }
    return err
}
