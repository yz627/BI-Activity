// service/college_service/profile_service.go
package college_service

import (
    "bi-activity/dao/college_dao"
    "bi-activity/response/college_response"
    "bi-activity/response/errors/college_error"
)

type CollegeProfileService struct {
    collegeDao *college_dao.CollegeDao
}

func NewCollegeProfileService(collegeDao *college_dao.CollegeDao) *CollegeProfileService {
    return &CollegeProfileService{
        collegeDao: collegeDao,
    }
}

// GetCollegeProfile 获取学院资料
func (s *CollegeProfileService) GetCollegeProfile(collegeID uint) (*college_response.CollegeProfileResponse, error) {
    // 获取学院信息
    college, err := s.collegeDao.GetCollegeByID(collegeID)
    if err != nil {
        return nil, err
    }

    // 转换为响应结构
    return &college_response.CollegeProfileResponse{
        ID:                  college.ID,
        CollegeAccount:      college.CollegeAccount,
        CollegeName:         college.CollegeName,
        AdminName:           college.AdminName,
        AdminPhone:          college.AdminPhone,
        AdminEmail:          college.AdminEmail,
        AdminIDNumber:       college.AdminIDNumber,
        Campus:              college.Campus,
        CollegeAddress:      college.CollegeAddress,
        CollegeIntroduction: college.CollegeIntroduction,
        AdminImage: college_response.Image{
            ID:       college.AdminImage.ID,
            FileName: college.AdminImage.FileName,
            URL:      college.AdminImage.URL,
            Type:     college.AdminImage.Type,
        },
        CollegeAvatar: college_response.Image{
            ID:       college.CollegeAvatar.ID,
            FileName: college.CollegeAvatar.FileName,
            URL:      college.CollegeAvatar.URL,
            Type:     college.CollegeAvatar.Type,
        },
        UpdatedAt: college.UpdatedAt,
    }, nil
}

// UpdateCollegeProfile 更新学院资料
func (s *CollegeProfileService) UpdateCollegeProfile(collegeID uint, req *college_response.UpdateProfileRequest) error {
    // 获取学院信息
    college, err := s.collegeDao.GetCollegeByID(collegeID)
    if err != nil {
        return err
    }

    // 更新字段
    college.CollegeName = req.CollegeName
    college.CollegeAccount = req.CollegeAccount
    college.Campus = req.Campus
    college.CollegeAddress = req.CollegeAddress
    college.CollegeIntroduction = req.CollegeIntroduction

    // 保存更新
    return s.collegeDao.UpdateCollege(college)
}

// UpdateCollegeAdminInfo 更新管理员信息
func (s *CollegeProfileService) UpdateCollegeAdminInfo(collegeID uint, req *college_response.UpdateAdminInfoRequest) error {
    // 获取学院信息
    college, err := s.collegeDao.GetCollegeByID(collegeID)
    if err != nil {
        return err
    }

    // 检查手机号是否已被其他学院使用
    if college.AdminPhone != req.AdminPhone {
        exists, err := s.collegeDao.CheckPhoneExists(req.AdminPhone, collegeID)
        if err != nil {
            return err
        }
        if exists {
            return college_error.ErrPhoneExistsError
        }
    }

    // 检查邮箱是否已被其他学院使用
    if college.AdminEmail != req.AdminEmail {
        exists, err := s.collegeDao.CheckEmailExists(req.AdminEmail, collegeID)
        if err != nil {
            return err
        }
        if exists {
            return college_error.ErrEmailExistsError
        }
    }

    // 更新管理员信息
    college.AdminName = req.AdminName
    college.AdminIDNumber = req.AdminIDNumber
    college.AdminPhone = req.AdminPhone
    college.AdminEmail = req.AdminEmail

    return s.collegeDao.UpdateCollegeAdminInfo(college)
}

// UpdateAdminAvatar 更新管理员头像
func (s *CollegeProfileService) UpdateAdminAvatar(collegeID uint, avatarID uint) error {
    return s.collegeDao.UpdateAdminAvatar(collegeID, avatarID)
}

// UpdateCollegeAvatar 更新学院头像
func (s *CollegeProfileService) UpdateCollegeAvatar(collegeID uint, avatarID uint) error {
    return s.collegeDao.UpdateCollegeAvatar(collegeID, avatarID)
}