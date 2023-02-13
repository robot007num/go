package pkg

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"time"
)

//CreateSnowID 创建唯一UserID 强唯一性
func CreateSnowID() (int64, error) {

	//设置时间
	startTime, err := time.Parse("2006/01/02", "2023/01/11")
	if err != nil {
		panic("time parse err :" + err.Error())
	}
	snowflake.Epoch = startTime.UnixNano() / 1000000

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}

	// Generate and print, all in one.
	//fmt.Printf("ID       : %d\n", node.Generate().Int64())
	return node.Generate().Int64(), nil
}

//CreateUUID 创建随机用户名数字 有一定几率重复
func CreateUUID() (uint32, error) {
	// V4 基于随机数
	u1, err := uuid.NewRandom()
	if err != nil {
		return 0, err
	}
	return uuid.Must(u1, err).ID(), nil
}
