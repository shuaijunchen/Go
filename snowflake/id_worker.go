// id_worker.go
package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	twepoch        = int64(1417937700000) // 默认起始时间戳 1417937700000
	DistrictIdBits = uint(5)              // 区域 所占用位置
	NodeIdBits     = uint(9)              // 节点 所占位置
	sequenceBits   = uint(10)             // 自增Id 所占用位置

	maxNodeId     = -1 ^ (-1 << NodeIdBits)     // 节点Id 最大范围
	maxDistrictId = -1 ^ (-1 << DistrictIdBits) // 最大区域范围

	nodeIdShift        = sequenceBits // 左移次数
	DistrictIdShift    = sequenceBits + NodeIdBits
	timestampLeftShift = sequenceBits + NodeIdBits + DistrictIdBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)
	maxNextIdsNum      = 100 // 单次获取最大数量
)

type IdWorker struct {
	sequence      int64 // 序号
	lastTimestamp int64 // 最后时间戳
	nodeId        int64 // 节点Id
	twepoch       int64
	districtId    int64
	mutex         sync.Mutex
}

func NewIdWorker(NodeId int64) (*IdWorker, error) {
	var districtId int64
	districtId = 1 // 默认给1
	idWorker := &IdWorker{}
	if NodeId > maxNodeId || NodeId < 0 {
		fmt.Sprintf("NodeId Id can't be greater than %d or less than 0", maxNodeId)
		return nil, errors.New(fmt.Sprintf("NodeId Id: %d error", NodeId))
	}
	if districtId > maxDistrictId || districtId < 0 {
		fmt.Sprintf("District Id can't be greater than %d or less than 0", maxDistrictId)
		return nil, errors.New(fmt.Sprintf("District Id: %d error", districtId))
	}
	idWorker.nodeId = NodeId
	idWorker.districtId = districtId
	idWorker.lastTimestamp = -1
	idWorker.sequence = 0
	idWorker.twepoch = twepoch
	idWorker.mutex = sync.Mutex{}
	fmt.Sprintf("worker starting. timestamp left shift %d,District id bits %d, worker id bits %dworker id bits %d, sequence bits %d, workerid %d", timestampLeftShift, DistrictIdBits, NodeIdBits, sequenceBits, NodeId)
	return idWorker, nil
}

func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func tilNextMilis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}

func (id *IdWorker) NextId() (int64, error) {
	id.mutex.Lock()
	defer id.mutex.Unlock()
	return id.nextId()
}

func (id *IdWorker) nextId() (int64, error) {
	timestamp := timeGen()
	if timestamp < id.lastTimestamp {
		return 0, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", id.lastTimestamp-timestamp))
	}
	if id.lastTimestamp == timestamp {
		id.sequence = (id.sequence + 1) & sequenceMask
		if id.sequence == 0 {
			timestamp = tilNextMilis(id.lastTimestamp)
		}
	} else {
		id.sequence = 0
	}
	id.lastTimestamp = timestamp
	return ((timestamp - id.twepoch) << timestampLeftShift) | (id.districtId << DistrictIdShift) | (id.nodeId << nodeIdShift) | id.sequence, nil
}
