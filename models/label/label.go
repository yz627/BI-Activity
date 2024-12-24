package label

// 审核状态
const (
	AuditStatusPending  = iota + 1 // 待审核
	AuditStatusPassed              // 已通过
	AuditStatusRejected            // 已拒绝
	AuditStatusRemoved             // 已移出
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

// 活动状态
const (
	ActivityStatusPending    = iota + 1 // 活动状态 审核中
	ActivityRecruiting                  // 活动状态 招募中
	ActivityStatusProceeding            // 活动状态 进行中
	ActivityStatusEnded                 // 活动状态 已结束
	ActivityStatusRejected              // 活动状态 未通过
)

// 招募限制
const (
	RecruitmentRestrictionUnlimited = iota + 1 // 招募人员无限制
	RecruitmentRestrictionCollege              // 学院内招募
)

// 管理员权限
const (
	RoleSuperAdmin = iota + 1 // 一级管理员
	RoleAdmin                 // 二级管理员
)
