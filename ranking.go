package main

import (
	"fmt"
	"math"
)

type PlayerInfo struct {
	Uid   string
	Name  string
	Score uint32
	Rank  int32
	VIPLv uint32
}

type RankingInfo struct {
	Id      uint32 // 排行榜Id
	Adapter uint32
	PLayers []PlayerInfo // 这里应该返回不可变的slice
}

// GenerateScore
//
//	@Description: 生成玩家上榜分数，注意用起服时候的时区或者用游戏里的时间就行
func (p *PlayerInfo) GenerateScore() {
	//组合成 分数.(下个月开始时间-当前时间)
}

// GetPLayerInfoRank
//
//	@Description:获取玩家当前排名。更新和获取应当由业务层保证，或者直接用下面的update获取
//	@param uid 玩家uid
//	@return uint32 玩家当前排名
func (r *RankingInfo) GetPLayerInfoRank(uid string) uint32 {
	//ZREVRANK 方法拿到玩家自己的排名。
	return 0
}

// UpdateRanking
//
//	@Description: 更新玩家排名
//	@param uid 玩家唯一Id
//	@param score 玩家分数
//	@return int32 返回玩家此时排名,这里其实可以返回名次+错误
func (r *RankingInfo) UpdateRanking(uid string, score uint32) int32 {
	// 这里应该由业务去保证，不应该在这里做
	if score > 10000 {
		score = 100000
	}
	//	ADD方法
	return 0 // 返回玩家此时的排名，ADD失败返回错误或者-1，由调用方去控制流程
}

// GetRankingRange
//
//	@Description:这里返回玩家的
//	@receiver r
//	@param uid
//	@return []PlayerInfo
func (r *RankingInfo) GetRankingRange(uid string) []PlayerInfo {
	//ZREVRANK 方法拿到玩家自己的排名
	var rank float64 //例子,这是玩家的排名
	start := math.Max(0, rank-10)
	//ZCARD返回排行榜长度
	var length float64
	end := math.Min(rank+10, length)
	//ZREVRANGE传入rank
	fmt.Sprintf("start:%v,end:%v", start, end)
	//players := ZREVRANGE返回值，组装好排名返回即可。更多信息可以由业务层组装
	return make([]PlayerInfo, 0)
}

// DeletePLayerInfo
//
//	@Description: 从排行榜删除玩家
//	@param uid
func (r *RankingInfo) DeletePLayerInfo(uid string) {
	//ZREM删除玩家
}

// TODO：优化
// 考虑抽一个排行榜的通用结构出来，通过构造器传入排行榜Id和排序权重进行初始化，考虑红黑树当参数做内存缓存，需要保证缓存一致性
// 业务层调用考虑增加配置，展示和缓存的条数通过配置实现
type RankingManager struct {
}

// GetRankingInfo
//
//	@Description:类似builder，通过传入的参数生成一个排行榜对象
//	@param id 排行榜Id
//	@param adapter 权重控制。
//	@return RankingInfo 排行榜对象
func (service *RankingManager) GetRankingInfo(id, adapter uint32) RankingInfo {
	return RankingInfo{
		Id:      id,
		Adapter: adapter,
	}
}

func (r *RankingInfo) adapterControl() {
	//组装好数据之后在这里写对应的更深一层的权重控制
}
