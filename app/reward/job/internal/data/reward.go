package data

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	walletClientV1 "sign-in/api/virtualwallet/v1"
	"strconv"
	"strings"
	"sync"
)

var (
	topic = "reward"
	ErrIncorrectFormat = errors.New("error: incorrect message format")

)


type rewardInfo struct {
	userId int64
	amount float64
}


type RewardRepo struct {
	data *Data
	log  *log.Helper
}

func NewRewardRepo(data *Data, logger log.Logger) *RewardRepo {
	return &RewardRepo{data: data, log: log.NewHelper(logger)}
}

func (r *RewardRepo) Consume(ctx context.Context) error {

	partitionList, err := r.data.kc.Partitions(topic)
	if err != nil {
		r.log.Errorf("fail to get list of partition,err: %v\n", err)
		return err
	}
	wg := sync.WaitGroup{}
	for partition := range partitionList {
		// 针对每个分区创建一个对应的消费者
		pc, err := r.data.kc.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			r.log.Errorf("failed to start consumer for partition %d, err: %v\n",
				partition, err)
			return err
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				r.log.Infof("Partition:%d Offset:%d Key:%v Value:%v\n",
					msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				err := r.doReward(ctx, string(msg.Value))
				if err != nil {
					r.log.Fatalf("do reward to user failed, please check err: %v", err)
				}
				defer pc.AsyncClose()
				wg.Done()
			}
		}(pc)
		wg.Wait()

	}
	return nil
}

func (r *RewardRepo) doReward(ctx context.Context, message string) error {

	rewardInfo, err := transferMessage(message)
	if err != nil {
		return err
	}
	_, err = r.data.wc.Credit(ctx, &walletClientV1.CreditRequest{
		User: rewardInfo.userId,
		Amount: rewardInfo.amount})
	return err

}

func transferMessage(message string) (*rewardInfo, error) {
	msg := strings.Split(message, "-")
	if len(msg) < 2 {
		return nil, ErrIncorrectFormat
	}
	userId, err := strconv.Atoi(msg[0])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	amount, err := strconv.ParseFloat(msg[1],64)
	if err != nil {
		return nil, ErrIncorrectFormat
	}

	return &rewardInfo{
		userId: int64(userId),
		amount: amount,
	}, nil
}

