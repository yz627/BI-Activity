package college

import (
	"bi-activity/models"
	"bi-activity/models/label"
	"log"
	"strings"
	"time"
)

type ActivityReleaseRequest struct {
	//ID                       uint           // 前端无需传递,后端直接设置           `gorm:"primaryKey;autoIncrement" json:"id"`                 // 主键，自动递增
	//ActivityNature           int            // 前端无需传递,后端直接设置,学院活动默认为2            `gorm:"type:tinyint;null" json:"activity_nature"`           // 活动性质(学生活动\学院活动)
	//ActivityStatus           int            // 前端无需传递,后端直接设置int            `gorm:"type:tinyint;null" json:"activity_status"`           // 活动状态
	//ActivityPublisherID      uint           // 前端需要传递,后端从token中获取           `gorm:"type:bigint;null" json:"activity_publisher_id"`      // 发布者 ID
	ActivityName         string // 前端需要传递         `gorm:"type:varchar(255);null" json:"activity_name"`        // 活动名称
	ActivityTypeID       uint   // 前端需要传递           `gorm:"type:bigint;null" json:"activity_type_id"`           // 活动类型 ID
	ActivityAddress      string // 前端需要传递         `gorm:"type:varchar(255);null" json:"activity_address"`     // 活动地址
	ActivityIntroduction string // 前端需要传递         `gorm:"type:text" json:"activity_introduction"`             // 活动简介
	ActivityContent      string // 前端需要传递         `gorm:"type:text" json:"activity_content"`                  // 活动内容
	ActivityImageUrl     string // 前端需要传递
	//ActivityImageID          uint           // 前端无需传递,后端直接设置           `gorm:"type:bigint" json:"activity_image_id"`               // 活动图片 ID
	ActivityDate             string // 前端需要传递         `gorm:"type:datetime;null" json:"activity_date"`            // 活动日期
	StartTime                string // 前端需要传递         `gorm:"type:datetime;null" json:"start_time"`               // 活动开始时间
	EndTime                  string // 前端需要传递         `gorm:"type:datetime;null" json:"end_time"`                 // 活动结束时间
	RecruitmentNumber        int    // 前端需要传递            `gorm:"type:tinyint;null" json:"recruitment_number"`        // 招募人数
	RegistrationRestrictions int    // 前端需要传递            `gorm:"type:tinyint;null" json:"registration_restrictions"` // 报名限制
	RegistrationRequirement  string // 前端需要传递         `gorm:"type:text" json:"registration_requirement"`          // 报名要求
	RegistrationDeadline     string // 前端需要传递         `gorm:"type:datetime;null" json:"registration_deadline"`    // 报名截止时间
	ContactName              string // 前端需要传递         `gorm:"type:varchar(10);null" json:"contact_name"`          // 联系人姓名
	ContactDetails           string // 前端需要传递         `gorm:"type:varchar(20);null" json:"contact_details"`       // 联系人电话
	//CreatedAt                time.Time      // 前端无需传递,后端直接设置      `gorm:"type:datetime;null" json:"created_at"`               // 创建时间
	//UpdatedAt                time.Time      // 前端无需传递,后端直接设置      `gorm:"type:datetime;null" json:"updated_at"`               // 更新时间
	//DeletedAt                gorm.DeletedAt // 前端无需传递,后端无需设置 `gorm:"index" json:"-"`
}

func (a *ActivityReleaseRequest) GetActivity() *models.Activity {
	var activity = models.Activity{}
	// 活动性质
	activity.ActivityNature = label.ActivityNatureCollege
	// 活动状态
	now := time.Now().Format("2006-01-02 15:04:05")
	if now < activity.StartTime { // 招募中
		activity.ActivityStatus = label.ActivityStatusRecruiting
	} else if now < activity.EndTime { // 活动审核通过，但还未结束：进行中
		activity.ActivityStatus = label.ActivityStatusProceeding
	} else if now >= activity.EndTime { // 活动审核通过，但已结束：已结束
		activity.ActivityStatus = label.ActivityStatusEnded
	}
	// 活动发布者

	// 活动名称
	activity.ActivityName = a.ActivityName
	// 活动类型
	activity.ActivityTypeID = a.ActivityTypeID
	// 活动地址
	activity.ActivityAddress = a.ActivityAddress
	// 活动介绍
	activity.ActivityIntroduction = a.ActivityIntroduction
	// 活动内容
	activity.ActivityContent = a.ActivityContent
	// 活动图片地址

	// 活动图片id

	// 活动日期
	log.Println("活动日期：", a.ActivityDate)
	activity.ActivityDate = a.ActivityDate
	// 开始时间
	log.Println("开始日期：", a.StartTime)
	activity.StartTime = a.StartTime
	// 结束时间
	log.Println("结束日期：", a.EndTime)
	activity.EndTime = a.EndTime
	// 招募人数
	activity.RecruitmentNumber = a.RecruitmentNumber
	// 招募限制
	activity.RegistrationRestrictions = a.RegistrationRestrictions
	// 招募要求
	activity.RegistrationRequirement = a.RegistrationRequirement
	// 招募截止时间
	log.Println("招募截止时间：", a.RegistrationDeadline)
	activity.RegistrationDeadline = a.RegistrationDeadline
	// 联系人
	activity.ContactName = a.ContactName
	// 联系方式
	activity.ContactDetails = a.ContactDetails
	// 创建时间
	cur := time.Now()
	activity.CreatedAt = cur
	// 更新时间
	activity.UpdatedAt = cur

	return &activity
}

func (a *ActivityReleaseRequest) GetImage() *models.Image {
	if len(a.ActivityImageUrl) <= 0 {
		return nil
	}
	// 获取file_name, url
	var image = models.Image{}
	var now = time.Now()
	lastSlashIndex := strings.LastIndex(a.ActivityImageUrl, "/")
	image.URL = a.ActivityImageUrl[:lastSlashIndex+1]
	image.FileName = a.ActivityImageUrl[lastSlashIndex+1:]
	image.CreatedAt = now
	image.UpdatedAt = now
	image.Type = models.ImageTypeActivity

	return &image
}
