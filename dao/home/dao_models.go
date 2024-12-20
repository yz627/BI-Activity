package home

type SearchParams struct {
	ActivityPublisherID uint   // 发布者ID
	ActivityNature      int    // 活动性质 0 - 全部 1 - 个人活动 2 - 学院活动 || 0 - 全部 1 - 我的发布 2 - 我的参与, 其余非法
	ActivityStatus      int    // 活动状态 0 - 全部 2 - 招募中 3 - 活动开始 4 - 活动结束, 其余非法
	ActivityDateStart   string // 活动日期 YYYY-MM-DD
	ActivityDateEnd     string // 活动日期 YYYY-MM-DD
	ActivityTypeID      uint   // 活动类别ID 0 - 全部
	Keyword             string // 搜索关键字，活动名称相关
	Page                int    // 页码
}
