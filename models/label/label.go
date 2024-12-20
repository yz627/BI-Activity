package label

// 审核状态 Participate 表
// 1 - 待审核 (学生刚报名)
// 2 - 已录取 (活动发布者通过)
// 3 - 未录取 (活动发布者拒绝)
// 4 - 已取消报名 (学生主动取消)
const (
	ParticipateStatusPending  = iota + 1 // 待审核
	ParticipateStatusPassed              // 已通过
	ParticipateStatusRejected            // 已拒绝
	ParticipateStatusCanceled            // 已取消报名
)

// 其余审核表
// 1 - 审核中
// 2 - 审核通过
// 3 - 审核不通过
const (
	AuditStatusPending  = iota + 1 // 审核中
	AuditStatusPassed              // 已通过
	AuditStatusRejected            // 已拒绝
)

// 性别
const (
	GenderFemale = iota + 1 // 女生
	GenderMale              // 男生
)

// 邀请码状态
const (
	InviteCodeStatusUnused = iota + 1 // 未使用
	InviteCodeStatusUsed              // 已使用
)

// 图片类型
const (
	ImageTypeAvatar   = iota + 1 // 1-头像
	ImageTypeCollege             // 2-活动图片
	ImageTypeBanner              // 3-轮播图
	ImageTypeActivity            // 4-学院图标
)

// 校园
const (
	CampusZhuHai    = iota + 1 // 珠海校区
	CampusGuangZhou            // 广州校区
	CampusShenZhen             // 深圳校区
)

// 活动性质
const (
	ActivityNatureStudent = iota + 1 // 学生活动
	ActivityNatureCollege            // 学院活动
)

const (
	ActivityMyPublish     = iota + 1 // 我的发布
	ActivityMyParticipate            // 我的参与
)

// 活动状态流转:Activity表
// 1-审核中(刚发布的活动)
// 2-招募中(审核通过,开始招募)
// 3-活动开始
// 4-活动结束
// 5-审核失败
const (
	ActivityStatusPending     = iota + 1 // 待审核
	ActivityStatusRecruiting             // 招募中
	ActivityStatusProceeding             // 活动进行中
	ActivityStatusEnded                  // 活动已结束
	ActivityStatusAuditFailed            // 审核失败
)

// 招募限制
const (
	RecruitmentRestrictionUnlimited = iota + 1 // 招募人员无限制
	RecruitmentRestrictionCollege              // 学院内招募
)

var (
	RecruitmentRestriction = map[int]string{
		RecruitmentRestrictionUnlimited: "招募人员无限制",
		RecruitmentRestrictionCollege:   "学院内招募",
	}
)

// 管理员权限
const (
	RoleSuperAdmin = iota + 1 // 一级管理员
	RoleAdmin                 // 二级管理员
)
