package student_controller

import (
	"bi-activity/service/student_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCollegeNameByStudentID(c *gin.Context) {
	studentIDStr := c.Param("id")
	studentID, err := strconv.Atoi(studentIDStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的学生 ID"})
		return
	}

	collegeName, err := student_service.GetCollegeNameByStudentID(uint(studentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student_id":   studentID,
		"college_name": collegeName,
	})
}