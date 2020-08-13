package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var Consumer sarama.Consumer

func Init() (err error) {
	Consumer, err = sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("sarama.NewConsumer err: ", err)
		return err
	}
	partitionList, err := Consumer.Partitions("web")
	if err != nil {
		fmt.Println("Consumer.Partitions err: ", err)
		return err
	}
	fmt.Println(partitionList)
	for partion := range partitionList {
		//对应每一个分区创建一个对应的分区消费者
		pc, err := Consumer.ConsumePartition("nginx", int32(partion), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Consumer.ConsumePartition err: ", err)
			return err
		}
		defer pc.AsyncClose()
		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d offset:%d key:%v value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	for {

	}
	return
}
