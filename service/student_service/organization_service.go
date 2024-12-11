package student_service

import (
	"bi-activity/dao/student_dao"
	"bi-activity/global"
	"fmt"
)

func GetCollegeNameByStudentID(studentId uint) (string, error) {
	student, err := student_dao.GetStudentByID(global.Db, studentId)

	if err != nil {
		return "", fmt.Errorf("无法查询学生信息: %v", err)
	}

	college, err := student_dao.GetCollegeNameByID(global.Db, student.CollegeID)

	if  err != nil {
		return "", fmt.Errorf("无法查询学院信息: %v", err)
	}

	return college.CollegeName, nil
}