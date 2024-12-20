package home

// Image service层返回的Image结构体
type Image struct {
	ID       uint
	FileName string
	Url      string
}

// ActivityCard service层返回的Activity卡片信息的结构体
type ActivityCard struct {
	ID                    uint   // 主键
	ActivityName          string // 活动名称
	ActivityDate          string // 活动日期：YYYY-MM-DD
	StartTime             string // 活动开始时间： HH:MM
	EndTime               string // 活动结束时间： HH:MM
	ActivityPublisherName string // 发布者名称
	ActivityTypeName      string // 活动类型名称
	ActivityTypeImageUrl  string // 活动类型图片地址
	RemainingNumber       int    // 活动招募剩余人数
}

// Activity service层返回的Activity详细信息的结构体
type Activity struct {
	// 活动信息
	ID                   uint   // 主键，自动递增
	ActivityAddress      string // 活动指点
	ContactName          string // 联系人姓名
	ContactDetails       string // 联系人电话
	ActivityTypeName     string // 活动类别（标签）
	ActivityTypeImageUrl string // 活动类别图片地址

	// 志愿招募信息
	ActivityDate      string // 活动日期：YYYY-MM-DD
	StartTime         string // 开始时间：HH:MM
	EndTime           string // 结束时间：HH:MM
	RecruitmentNumber int    // 招募人数
	RecruitedNumber   int    // 已招募人数

	// 活动限制
	RegistrationRestrictions string // 报名限制：本学院成员、全校成员
	RegistrationRequirement  string // 报名要求
	RegistrationDeadline     string // 报名截止时间

	// 活动介绍
	ActivityIntroduction string // 活动简介
	ActivityContent      string // 活动内容

	ActivityName     string // 活动名称
	ActivityImageUrl string // 活动图片地址
	PublisherName    string // 活动发布者名称
	CreatedAt        string // 活动发布时间
	ActivityStatus   int    // 活动状态：

	// 报名状态
	ParticipateStatus int
}

// ActivityType service层返回的ActivityType结构体
type ActivityType struct {
	ID       uint
	TypeName string
	Url      string
}

type BiData struct {
	ActivityTotal int
	StudentTotal  int
	CollegeTotal  int
}

type BiDataLeaderboard struct {
	CollegeName  string
	StudentTotal int
}

type Help struct {
	Problem string
	Answer  string
}

// SearchActivityParams 活动查询参数
type SearchActivityParams struct {
	ActivityPublisherID uint   // 发布者ID
	ActivityNature      int    // 活动性质 0 - 全部 1 - 个人活动 2 - 学院活动 || 0 - 全部 1 - 我的发布 2 - 我的参与, 其余非法
	ActivityStatus      int    // 活动状态 0 - 全部 2 - 招募中 3 - 活动开始 4 - 活动结束, 其余非法
	ActivityDateStart   string // 活动日期 YYYY-MM-DD
	ActivityDateEnd     string // 活动日期 YYYY-MM-DD
	ActivityTypeID      uint   // 活动类别ID 0 - 全部
	Keyword             string // 搜索关键字，活动名称相关
	Page                int    // 页码
}

// 活动状态流转：Activity 表
// 1 - 审核中 (刚发布的活动)
// 2 - 招募中 (审核通过，开始招募)
// 3 - 活动开始
// 4 - 活动结束
// 5 - 审核失败
