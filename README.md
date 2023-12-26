# ranking
日活跃用户在10万⼈以上的排行榜
1. 在每月活动中 ，玩家得到的活动总分为 0 到 10000 之间的整数。
2. 在每月活动结束之后 ，需要依据这⼀活动总分 ，从高到低为玩家建立排行榜。
3. 如果多位玩家分数相同 ，则按得到指定分数顺序排序 ，先得到的玩家排在前面。
4. 系统提供玩家名次查询接口 ，玩家能够查询自己名次前后10位玩家的分数和名次。
5. 请使用UML图或线框图表达设计 ，关键算法可使用流程图或伪代码表达。
6. 如果玩家分数，触发时间均相同，则根据玩家等级，名字依次排序

如果用户量小的话直接slice+map排序，适合玩法内的小排行，省去redis的一些开销
这里考虑redis的zset结构

分析：
1. 活动总分不大，用redis的话需要用float64类型
2. 每月一清
3. 有特殊排序：
        初始化排行榜的时候获取到下个周期的时间
        用此时间和玩家分数组成玩家的排名
4. 提供简单的接口
5. 有特殊的排序--放在优化里，通用排行榜不考虑，留到业务层去做。

