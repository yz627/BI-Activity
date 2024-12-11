package testcase

import (
	"bi-activity/dao/student_dao"
	"bi-activity/models"
	"log"

	"gorm.io/gorm"
)

func RunStudentTest(db *gorm.DB) {
	//log.Println("Running student table tests")

	//testCreateStudent(db)

	//testGetStudentByID(db)

	//testDelStudentByID(db)
}

func testCreateStudent(db *gorm.DB) {
	student := &models.Student{
		StudentPhone: "1234567890",
		StudentEmail: "test@student.com",
		StudentID:    "S12345",
		Password:     "securepassword",
		StudentName:  "John Doe",
		Gender:       1,
		Nickname:     "Johnny",
	}

	err := student_dao.CreateStudent(db, student)
	if err != nil {
		log.Fatalf("testCreateStudent failed: %v", err)
	}
	log.Println("testCreateStudent passed!")
}

func testGetStudentByID(db *gorm.DB) {
	id := uint(1)
	student, err := student_dao.GetStudentByID(db, id)
	if err != nil {
		log.Fatalf("testGetStudentByID failed: %v", err)
	}
	log.Printf("testGetStudentByID passed! Student: %+v\n", student)
}

func testDelStudentByID(db *gorm.DB) {
	id := uint64(1)
	err := student_dao.DeleteStudentByID(db, id)
	if err != nil {
		log.Fatalf("testDelStudentByID failed: %v", err)
	}
	log.Println("testDelStudentByID passed! Student")
}
