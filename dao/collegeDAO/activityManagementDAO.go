package collegeDAO

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/college"
	"bi-activity/models/label"
	cr "bi-activity/response/college"
	"log"
	"time"
)

type ActivityManagementDAO struct {
	// 数据库连接
	data *dao.Data
}

func NewActivityManagementDAO(data *dao.Data) *ActivityManagementDAO {
	return &ActivityManagementDAO{data: data}
}

func (a *ActivityManagementDAO) GetAuditRecord(collegeId uint, status, page, size int) *cr.Result {
	db := a.data.DB()
	var total int64 = 0
	// 计算总数
	query := db.Table("student_activity_audit")                                              // 从student_activity_audit表开始
	query = query.Select("count(*)")                                                         // 选择计数
	query = query.Joins("join activity on activity.id = student_activity_audit.activity_id") // 连接Activity表
	query = query.Joins("join student on student.id = activity.activity_publisher_id")       // 连接Student表
	query = query.Where("student_activity_audit.college_id = ?", collegeId)                  // 添加学院ID条件
	query = query.Where("student_activity_audit.status = ?", status)                         // 添加状态条件
	query = query.Where("activity.activity_nature = ?", label.ActivityNatureStudent)
	query.Count(&total)

	var result = []college.ActivityAuditRecord{}
	// 构建查询
	query2 := db.Table("student_activity_audit") // 从StudentActivityAudit表开始
	// 选择需要的字段
	query2 = query2.Select("student_activity_audit.id as RecordId," +
		" activity.activity_name as ActivityName," +
		" activity.start_time as StartTime," +
		" activity.end_time as EndTime," +
		" activity.activity_address as ActivityAddress," +
		" student.student_name as Organizer," +
		" student_activity_audit.created_at as ApplicationTime," +
		" student_activity_audit.status as Status")
	query2 = query2.Joins("join activity on activity.id = student_activity_audit.activity_id") // 连接Activity表
	query2 = query2.Joins("join student on student.id = activity.activity_publisher_id")       // 连接Student表
	query2 = query2.Where("student_activity_audit.college_id = ?", collegeId)                  // 添加学院ID条件
	query2 = query2.Where("student_activity_audit.status = ?", status)                         // 添加状态条件
	query2 = query2.Where("activity.activity_nature = ?", label.ActivityNatureStudent)
	// 排序
	query2 = query2.Order("student_activity_audit.updated_at desc")
	// 添加分页
	offset := (page - 1) * size
	query2 = query2.Limit(size).Offset(offset)
	// 执行查询
	query2.Scan(&result)

	log.Println(total)
	log.Println(result)

	return cr.NewResult(int(total), result)
}

func (a *ActivityManagementDAO) UpdateAuditRecord(id uint, status int) {
	db := a.data.DB()
	now := time.Now().Format("2006-01-02 15:04:05")
	// student_activity_audit表
	log.Println("审核结果：")
	log.Println(status)
	studentActivityAudit := models.StudentActivityAudit{}
	db.Where("id = ?", id).First(&studentActivityAudit)
	db.Model(&studentActivityAudit).Update("status", status).Update("updated_at", now)
	log.Println("审核记录")
	log.Println(studentActivityAudit)
	// activity表
	activity := models.Activity{}
	db.Where("id = ?", studentActivityAudit.ActivityID).First(&activity)
	var activityStatus = 0
	if status == label.AuditStatusRejected { // 活动审核未通过
		activityStatus = label.ActivityStatusRejected
	} else if status == label.AuditStatusPassed && now < activity.StartTime { // 活动审核通过，但还未开始：招募中
		activityStatus = label.ActivityRecruiting
	} else if status == label.AuditStatusPassed && now < activity.EndTime { // 活动审核通过，但还未结束：进行中
		activityStatus = label.ActivityStatusProceeding
	} else if status == label.AuditStatusPassed && now >= activity.EndTime { // 活动审核通过，但已结束：已结束
		activityStatus = label.ActivityStatusEnded
	}
	log.Println("活动新状态：")
	log.Println(activityStatus)
	db.Model(&activity).Update("activity_status", activityStatus).Update("updated_at", now)
}

func (a *ActivityManagementDAO) GetAdmissionRecord(collegeId uint, status, page, size int) *cr.Result {
	db := a.data.DB()
	var total int64 = 0
	// 计算总数
	query := db.Table("participant")                                              // 从student_activity_audit表开始
	query = query.Select("count(*)")                                              // 选择计数
	query = query.Joins("join activity on activity.id = participant.activity_id") // 连接Activity表
	query = query.Joins("join student on student.id = participant.student_id")    // 连接Student表
	query = query.Where("student.college_id = ?", collegeId)                      // 添加学院ID条件
	query = query.Where("participant.status = ?", status)                         // 添加状态条件
	query = query.Where("activity.activity_nature = ?", label.ActivityNatureCollege)
	query.Count(&total)

	var result = []college.ActivityAdmissionRecord{}
	// 构建查询
	query2 := db.Table("participant") // 从StudentActivityAudit表开始
	// 选择需要的字段
	query2 = query2.Select("participant.id as RecordId," +
		" activity.activity_name as ActivityName," +
		" student.student_name as StudentName," +
		" student.student_id as StudentId," +
		" participant.created_at as ApplicationTime," +
		" participant.status as Status")
	query2 = query2.Joins("join activity on activity.id = participant.activity_id") // 连接Activity表
	query2 = query2.Joins("join student on student.id =participant.student_id")     // 连接Student表
	query2 = query2.Where("student.college_id = ?", collegeId)                      // 添加学院ID条件
	query2 = query2.Where("participant.status = ?", status)                         // 添加状态条件
	query2 = query2.Where("activity.activity_nature = ?", label.ActivityNatureCollege)
	// 排序
	query2 = query2.Order("participant.updated_at desc")
	// 添加分页
	offset := (page - 1) * size
	query2 = query2.Limit(size).Offset(offset)
	// 执行查询
	query2.Scan(&result)

	log.Println(total)
	log.Println(result)

	return cr.NewResult(int(total), result)
}

func (a *ActivityManagementDAO) UpdateAdmissionRecord(id uint, status int) {
	db := a.data.DB()
	now := time.Now().Format("2006-01-02 15:04:05")
	// participant表
	log.Println("审核结果：")
	log.Println(status)
	participant := models.Participant{}
	db.Where("id = ?", id).First(&participant)
	db.Model(&participant).Update("status", status).Update("updated_at", now)
	log.Println("审核记录")
	log.Println(participant)
}
