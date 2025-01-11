package college

type ActivityReleaseRequest struct {
	ActivityNature       int    `json:"activity_nature" binding:"required"`
	ActivityName         string `json:"activity_name" binding:"required"`
	ActivityTypeID       uint   `json:"activity_type_id" binding:"required"`
	ActivityAddress      string `json:"activity_address" binding:"required"`
	ActivityIntroduction string `json:"activity_introduction"`
	ActivityContent      string `json:"activity_content" binding:"required"`
	ActivityImageID      uint   `json:"activity_image_id"`
	ActivityDate         string `json:"activity_date" binding:"required"`
	StartTime            string `json:"start_time" binding:"required"`
	EndTime              string `json:"end_time" binding:"required"`
	RecruitmentNumber    int    `json:"recruitment_number" binding:"required"`
	RegistrationDeadline string `json:"registration_deadline" binding:"required"`
	ContactName          string `json:"contact_name" binding:"required"`
	ContactDetails       string `json:"contact_details" binding:"required"`
}
