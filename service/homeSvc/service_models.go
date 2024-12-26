package homeSvc

// Image service层返回的Image结构体
type Image struct {
	ID       uint   `json:"id"`
	FileName string `json:"fileName"`
	URL      string `json:"url"`
}

// ActivityCard service层返回的Activity卡片信息的结构体
type ActivityCard struct {
	ID                    uint   `json:"id"`                    // 主键
	ActivityName          string `json:"activityName"`          // 活动名称
	ActivityDate          string `json:"activityDate"`          // 活动日期：YYYY-MM-DD
	StartTime             string `json:"startTime"`             // 活动开始时间： HH:MM
	EndTime               string `json:"endTime"`               // 活动结束时间： HH:MM
	ActivityPublisherName string `json:"activityPublisherName"` // 发布者名称
	ActivityTypeName      string `json:"activityTypeName"`      // 活动类型名称
	ActivityTypeImageUrl  string `json:"activityTypeImageUrl"`  // 活动类型图片地址
	RemainingNumber       int    `json:"remainingNumber"`       // 活动招募剩余人数
}

// Activity service层返回的Activity详细信息的结构体
type Activity struct {
	// 活动信息
	ID                   uint   `json:"id"`                   // 主键，自动递增
	ActivityAddress      string `json:"activityAddress"`      // 活动指点
	ContactName          string `json:"contactName"`          // 联系人姓名
	ContactDetails       string `json:"contactDetails"`       // 联系人电话
	ActivityTypeName     string `json:"activityTypeName"`     // 活动类别（标签）
	ActivityTypeImageUrl string `json:"activityTypeImageUrl"` // 活动类别图片地址

	// 志愿招募信息
	ActivityDate      string `json:"activityDate"`      // 活动日期：YYYY-MM-DD
	StartTime         string `json:"startTime"`         // 开始时间：HH:MM
	EndTime           string `json:"endTime"`           // 结束时间：HH:MM
	RecruitmentNumber int    `json:"recruitmentNumber"` // 招募人数
	RecruitedNumber   int    `json:"recruitedNumber"`   // 已招募人数

	// 活动限制
	RegistrationRestrictions string `json:"registrationRestrictions"` // 报名限制：本学院成员、全校成员
	RegistrationRequirement  string `json:"registrationRequirement"`  // 报名要求
	RegistrationDeadline     string `json:"registrationDeadline"`     // 报名截止时间

	// 活动介绍
	ActivityIntroduction string `json:"activityIntroduction"` // 活动简介
	ActivityContent      string `json:"activityContent"`      // 活动内容

	ActivityName     string `json:"activityName"`     // 活动名称
	ActivityImageUrl string `json:"activityImageUrl"` // 活动图片地址
	PublisherName    string `json:"publisherName"`    // 活动发布者名称
	CreatedAt        string `json:"createdAt"`        // 活动发布时间
	ActivityStatus   int    `json:"activityStatus"`   // 活动状态：

	// 报名状态
	ParticipateStatus int `json:"participateStatus"`
}

// ActivityType service层返回的ActivityType结构体
type ActivityType struct {
	ID       uint   `json:"id"`
	TypeName string `json:"typeName"`
	Url      string `json:"url"`
}

type BiData struct {
	ActivityTotal int `json:"activityTotal"`
	StudentTotal  int `json:"studentTotal"`
	CollegeTotal  int `json:"collegeTotal"`
}

type BiDataLeaderboard struct {
	CollegeName  string `json:"collegeName"`
	StudentTotal int    `json:"studentTotal"`
}

type Help struct {
	Problem string `json:"problem"`
	Answer  string `json:"answer"`
}

// SearchActivityParams 活动查询参数
type SearchActivityParams struct {
	ActivityPublisherID uint   `json:"activityPublisherID"` // 发布者ID
	ActivityNature      int    `json:"activityNature"`      // 活动性质 0 - 全部 1 - 个人活动 2 - 学院活动 || 0 - 全部 1 - 我的发布 2 - 我的参与, 其余非法
	ActivityStatus      int    `json:"activityStatus"`      // 活动状态 0 - 全部 2 - 招募中 3 - 活动开始 4 - 活动结束, 其余非法
	ActivityDateStart   string `json:"activityDateStart"`   // 活动日期 YYYY-MM-DD
	ActivityDateEnd     string `json:"activityDateEnd"`     // 活动日期 YYYY-MM-DD
	ActivityTypeID      uint   `json:"activityTypeID"`      // 活动类别ID 0 - 全部
	Keyword             string `json:"keyword"`             // 搜索关键字，活动名称相关
	Page                int    `json:"page"`                // 页码
}

type StuInfo struct {
	Name        string `json:"name"`        // 学生姓名
	ID          string `json:"id"`          // 学号
	Phone       string `json:"phone"`       // 电话号码
	Email       string `json:"email"`       // 邮箱
	CollegeName string `json:"collegeName"` // 学院名称
}

// 活动状态流转：Activity 表
// 1 - 审核中 (刚发布的活动)
// 2 - 招募中 (审核通过，开始招募)
// 3 - 活动开始
// 4 - 活动结束
// 5 - 审核失败
