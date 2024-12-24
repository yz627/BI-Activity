package collegeDAO

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/college"
	"bi-activity/models/label"
	cr "bi-activity/response/college"
	"fmt"
	"log"
	"time"
)

type MmDAO struct {
	// 数据库连接
	data *dao.Data
}

func NewMmDAO(data *dao.Data) *MmDAO {
	return &MmDAO{
		data: data,
	}
}

func (m *MmDAO) GetAuditRecord(id, status, page, size int) *cr.Result {
	// 数据库连接实例
	db := m.data.DB()
	// 查询结果
	total := 0
	records := []college.JAuditRecord{}
	// sql操作
	sql1 := fmt.Sprintf("SELECT count(*) FROM student s, join_college_audit j " +
		"WHERE s.id = j.student_id AND j.status = ? AND j.college_id = ?;")
	db.Raw(sql1, status, id).Scan(&total)
	sql2 := fmt.Sprintf("SELECT j.id, s.student_name, s.student_id, j.status, j.updated_at FROM student s, join_college_audit j " +
		"WHERE s.id = j.student_id AND j.status = ? AND j.college_id = ? " +
		"order by j.updated_at desc limit ?, ?;")
	db.Raw(sql2, status, id, (page-1)*size, size).Scan(&records)
	result := cr.NewResult(total, records)
	log.Println(result)
	return result
}

func (m *MmDAO) UpdateAuditRecord(audit *college.Audit) {
	// 数据库连接实例
	db := m.data.DB()
	// sql操作
	// 查询student.id
	sid := 0
	sql1 := fmt.Sprintf("SELECT id FROM student WHERE student_id = ?;")
	db.Raw(sql1, audit.StudentId).Scan(&sid)
	// 更新join_college_audit
	now := time.Now()
	sql2 := fmt.Sprintf("UPDATE join_college_audit SET status = ?, updated_at = ? " +
		"WHERE id = ?;")
	db.Exec(sql2, audit.Status, now, audit.AuditId)
	// 更新student.college_id
	if audit.Status == 2 {
		sql3 := fmt.Sprintf("UPDATE student SET college_id = ?, updated_at = ?" +
			" WHERE id = ?;")
		db.Exec(sql3, audit.CollegeId, now, sid)
	}
}

func (m *MmDAO) QueryMember(id, page, size int, studentName, studentId, start, end string) *cr.Result {
	// 数据库连接实例
	db := m.data.DB()
	// 查询结果
	total := 0
	records := []college.Member{}
	// sql操作
	sql1 := "SELECT count(*) FROM student s, join_college_audit j " + "WHERE s.college_id = j.college_id"
	// 条件分页查询条目
	sql2 := "SELECT s.student_name, s.student_id, j.updated_at FROM student s, join_college_audit j " + "WHERE s.college_id = j.college_id "
	if studentName != "" {
		sql1 += fmt.Sprintf(" AND s.student_name like '%%%s%%'", studentName)
		sql2 += fmt.Sprintf("AND s.student_name like '%%%s%%' ", studentName)
	}
	if studentId != "" {
		sql1 += fmt.Sprintf(" AND s.student_id like '%%%s%%'", studentId)
		sql2 += fmt.Sprintf("AND s.student_id like '%%%s%%' ", studentId)
	}
	if start != "" && end != "" {
		sql1 += fmt.Sprintf("AND j.updated_at BETWEEN '%s' AND '%s' ", start, end)
		sql2 += fmt.Sprintf("AND j.updated_at BETWEEN '%s' AND '%s' ", start, end)
	}

	sql1 += ";"
	sql2 += "ORDER BY j.updated_at DESC LIMIT ?, ?;"
	db.Raw(sql1).Scan(&total)
	db.Raw(sql2, (page-1)*size, size).Scan(&records)
	result := cr.NewResult(total, records)
	log.Println(result)
	return result
}

func (m *MmDAO) DeleteMember(collegeId uint, studentId string) {
	db := m.data.DB()
	// student表
	student := models.Student{}
	db.Where("student_id = ?", studentId).First(&student)
	db.Model(&student).Update("college_id", nil)
	// join_college_audit表
	joinCollegeAudit := models.JoinCollegeAudit{
		StudentID: student.ID,
		CollegeID: collegeId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    label.AuditStatusRemoved,
	}
	db.Create(&joinCollegeAudit)
}
