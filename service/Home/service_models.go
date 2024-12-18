package Home

import "time"

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
	StartTime             string // 活动开始时间
	EndTime               string // 活动结束时间
	ActivityPublisherName string // 发布者名称
	ActivityTypeName      string // 活动类型名称
	ActivityTypeImageUrl  string // 活动类型图片地址
	RemainingNumber       int    // 活动招募剩余人数
}

// Activity service层返回的Activity详细信息的结构体
type Activity struct {
	// 活动信息
	ID                       uint      // 主键，自动递增
	ActivityNature           int       // 活动性质
	ActivityStatus           int       // 活动状态
	ActivityName             string    // 活动名称
	ActivityAddress          string    // 活动地址
	ActivityIntroduction     string    // 活动简介
	ActivityContent          string    // 活动内容
	ActivityDate             string    // 活动日期
	StartTime                string    // 活动开始时间
	EndTime                  string    // 活动结束时间
	RecruitmentNumber        int       // 招募人数
	RegistrationRestrictions int       // 报名限制
	RegistrationRequirement  string    // 报名要求
	RegistrationDeadline     string    // 报名截止时间
	ContactName              string    // 联系人姓名
	ContactDetails           string    // 联系人电话
	CreatedAt                time.Time // 创建时间

	// 活动类型
	ActivityTypeName string

	// 活动图片信息
	Url string

	// 活动发起人信息
	PublisherName string
}

// ActivityType service层返回的ActivityType结构体
type ActivityType struct {
	ID       uint
	TypeName string
	Url      string
}
