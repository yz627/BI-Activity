package student_service

import (
	"bi-activity/dao/student_dao"
	"bi-activity/models"
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
)

type ActivityService interface {
	// 发布活动
	CreateActivity(publisherID uint, req *student_response.CreateActivityRequest) error

	// 获取活动信息
	GetMyActivities(publisherID uint) (*student_response.ActivityListResponse, error)
	GetActivityByID(id uint) (*student_response.ActivityResponse, error)

	// 活动状态更新
	UpdateActivityStatus(id uint, status int) error

	// 参与者管理
	GetParticipants(activityID uint) ([]*student_response.ParticipantResponse, error)
	UpdateParticipantStatus(participantID uint, status int) error
}

type ActivityServiceImpl struct {
	activityDao      student_dao.ActivityDao
	participantDao   student_dao.ParticipantDao
	activityAuditDao student_dao.ActivityAuditDao
	studentDao       student_dao.StudentDao
	collegeDao       student_dao.CollegeDao
}

func NewActivityService(activityDao student_dao.ActivityDao, participantDao student_dao.ParticipantDao, activityAuditDao student_dao.ActivityAuditDao, studentDao student_dao.StudentDao, collegeDao student_dao.CollegeDao) ActivityService {
	return &ActivityServiceImpl{
		activityDao:      activityDao,
		participantDao:   participantDao,
		activityAuditDao: activityAuditDao,
		studentDao:       studentDao,
		collegeDao:       collegeDao,
	}
}

// CreateActivity 创建活动
func (s *ActivityServiceImpl) CreateActivity(publisherID uint, req *student_response.CreateActivityRequest) error {
	// 获取发布者信息
	student, err := s.studentDao.GetByID(uint(publisherID))
	if err != nil {
		return student_error.ErrStudentNotFoundError
	}

	// 创建活动记录
	activity := &models.Activity{
		ActivityNature:           req.ActivityNature,
		ActivityStatus:           1, // 1: 审核中
		ActivityPublisherID:      publisherID,
		ActivityName:             req.ActivityName,
		ActivityTypeID:           req.ActivityTypeID,
		ActivityAddress:          req.ActivityAddress,
		ActivityIntroduction:     req.ActivityIntroduction,
		ActivityContent:          req.ActivityContent,
		ActivityImageID:          req.ActivityImageID,
		ActivityDate:             req.ActivityDate,
		StartTime:                req.StartTime,
		EndTime:                  req.EndTime,
		RecruitmentNumber:        req.RecruitmentNumber,
		RegistrationRestrictions: req.RegistrationRestrictions,
		RegistrationRequirement:  req.RegistrationRequirement,
		RegistrationDeadline:     req.RegistrationDeadline,
		ContactName:              req.ContactName,
		ContactDetails:           req.ContactDetails,
	}

	// 保存活动信息
	if err := s.activityDao.Create(activity); err != nil {
		return err
	}

	// 创建审核记录
	audit := &models.StudentActivityAudit{
		CollegeID:  student.CollegeID,
		ActivityID: activity.ID,
		Status:     1, // 1: 审核中
	}

	if err := s.activityAuditDao.Create(audit); err != nil {
		return err
	}

	return nil
}

// GetMyActivities 获取我的活动列表
func (s *ActivityServiceImpl) GetMyActivities(publisherID uint) (*student_response.ActivityListResponse, error) {
	// 获取活动列表
	activities, err := s.activityDao.GetByPublisherID(publisherID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	activityResponses := make([]student_response.ActivityResponse, 0, len(activities))
	for _, activity := range activities {
		activityResponses = append(activityResponses, student_response.ActivityResponse{
			ID:                       activity.ID,
			ActivityNature:           activity.ActivityNature,
			ActivityStatus:           activity.ActivityStatus,
			ActivityPublisherID:      activity.ActivityPublisherID,
			ActivityName:             activity.ActivityName,
			ActivityTypeID:           activity.ActivityTypeID,
			ActivityAddress:          activity.ActivityAddress,
			ActivityIntroduction:     activity.ActivityIntroduction,
			ActivityContent:          activity.ActivityContent,
			ActivityImageID:          activity.ActivityImageID,
			ActivityDate:             activity.ActivityDate,
			StartTime:                activity.StartTime,
			EndTime:                  activity.EndTime,
			RecruitmentNumber:        activity.RecruitmentNumber,
			RegistrationRestrictions: activity.RegistrationRestrictions,
			RegistrationRequirement:  activity.RegistrationRequirement,
			RegistrationDeadline:     activity.RegistrationDeadline,
			ContactName:              activity.ContactName,
			ContactDetails:           activity.ContactDetails,
		})
	}

	return &student_response.ActivityListResponse{
		Total:      int64(len(activities)),
		Activities: activityResponses,
	}, nil
}

// GetActivityByID 获取活动详情
func (s *ActivityServiceImpl) GetActivityByID(id uint) (*student_response.ActivityResponse, error) {
	// 获取活动详情
	activity, err := s.activityDao.GetByID(id)
	if err != nil {
		return nil, student_error.ErrActivityNotFoundError
	}

	// 转换为响应格式
	return &student_response.ActivityResponse{
		ID:                   activity.ID,
		ActivityNature:       activity.ActivityNature,
		ActivityStatus:       activity.ActivityStatus,
		ActivityPublisherID:  activity.ActivityPublisherID,
		ActivityName:         activity.ActivityName,
		ActivityTypeID:       activity.ActivityTypeID,
		ActivityAddress:      activity.ActivityAddress,
		ActivityIntroduction: activity.ActivityIntroduction,
		ActivityContent:      activity.ActivityContent,
		ActivityImageID:      activity.ActivityImageID,
		ActivityDate:         activity.ActivityDate,
		StartTime:            activity.StartTime,
		EndTime:              activity.EndTime,
		RecruitmentNumber:    activity.RecruitmentNumber,
        RegistrationRestrictions: activity.RegistrationRestrictions,
		RegistrationRequirement:  activity.RegistrationRequirement,
		RegistrationDeadline: activity.RegistrationDeadline,
		ContactName:          activity.ContactName,
		ContactDetails:       activity.ContactDetails,
	}, nil
}

// GetParticipants 获取活动的参与者列表
func (s *ActivityServiceImpl) GetParticipants(activityID uint) ([]*student_response.ParticipantResponse, error) {
	// 获取参与者列表
	participants, err := s.participantDao.GetByActivityID(activityID)
	if err != nil {
		return nil, err
	}

	// 构建响应
	var responses []*student_response.ParticipantResponse
	for _, p := range participants {
		// 获取学生信息
		student, err := s.studentDao.GetByID(p.StudentID)
		if err != nil {
			continue
		}

		// 获取学院信息
		college, err := s.collegeDao.GetByID(student.CollegeID)
		if err != nil {
			continue
		}

		responses = append(responses, &student_response.ParticipantResponse{
			ID:            p.ID,
			ActivityID:    p.ActivityID,
			StudentID:     p.StudentID,
			Status:        p.Status,
			StudentName:   student.StudentName,
			StudentPhone:  student.StudentPhone,
			CollegeName:   college.CollegeName,
			StudentNumber: student.StudentID, // 注意这里是学号
		})
	}

	return responses, nil
}

// UpdateParticipantStatus 更新参与者状态
func (s *ActivityServiceImpl) UpdateParticipantStatus(participantID uint, status int) error {
	// 获取参与记录
	participant, err := s.participantDao.GetByID(participantID)
	if err != nil {
		return student_error.ErrParticipantNotFoundError
	}

	// 获取活动信息
	activity, err := s.activityDao.GetByID(participant.ActivityID)
	if err != nil {
		return student_error.ErrActivityNotFoundError
	}

	// 检查活动状态是否允许更新参与者状态
	if activity.ActivityStatus != 2 { // 不是招募中状态
		return student_error.ErrActivityStatusInvalidError
	}

	// 如果是录取操作，需要检查是否超过招募人数
	if status == 2 { // 假设2表示已录取
		// 获取当前已录取人数
		participants, err := s.participantDao.GetByActivityID(activity.ID)
		if err != nil {
			return err
		}

		admittedCount := 0
		for _, p := range participants {
			if p.Status == 2 { // 已录取
				admittedCount++
			}
		}

		if admittedCount >= activity.RecruitmentNumber {
			return student_error.ErrActivityFullError
		}
	}

	// 更新状态
	if err := s.participantDao.UpdateStatus(participantID, status); err != nil {
		return err
	}

	return nil
}

// UpdateActivityStatus 更新活动状态
func (s *ActivityServiceImpl) UpdateActivityStatus(id uint, status int) error {
	// 检查活动是否存在
	activity, err := s.activityDao.GetByID(id)
	if err != nil {
		return student_error.ErrActivityNotFoundError
	}

	// 验证状态转换的合法性
	if !isValidStatusTransition(activity.ActivityStatus, status) {
		return student_error.ErrActivityStatusInvalidError
	}

	// 更新状态
	return s.activityDao.UpdateStatus(id, status)
}

// 辅助函数：验证状态转换的合法性
func isValidStatusTransition(currentStatus, newStatus int) bool {
	// 状态定义：
	// 1: 审核中
	// 2: 招募中（审核通过）
	// 3: 活动开始
	// 4: 活动结束
	// 5: 审核失败

	switch currentStatus {
	case 1: // 审核中 -> 招募中/审核失败
		return newStatus == 2 || newStatus == 5
	case 2: // 招募中 -> 活动开始
		return newStatus == 3
	case 3: // 活动开始 -> 活动结束
		return newStatus == 4
	default:
		return false
	}
}
