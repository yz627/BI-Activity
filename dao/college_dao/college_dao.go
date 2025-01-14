// dao/college_dao/college_dao.go
package college_dao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/response/errors/college_error"
	"errors"

	"gorm.io/gorm"
)

type CollegeDao struct {
    data *dao.Data
}

func NewCollegeDao(data *dao.Data) *CollegeDao {
    return &CollegeDao{data: data}
}

func (d *CollegeDao) GetCollegeByID(id uint) (*models.College, error) {
    var college models.College
    result := d.data.DB().Preload("AdminImage").Preload("CollegeAvatar").First(&college, id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, college_error.ErrCollegeNotFoundError
        }
        return nil, result.Error
    }
    return &college, nil
}

func (d *CollegeDao) UpdateCollege(college *models.College) error {
    result := d.data.DB().Model(college).Updates(map[string]interface{}{
        "college_name":          college.CollegeName,
        "college_account":       college.CollegeAccount,
        "campus":               college.Campus,
        "college_address":      college.CollegeAddress,
        "college_introduction": college.CollegeIntroduction,
    })
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// UpdateCollegeAdminInfo 更新管理员信息
func (d *CollegeDao) UpdateCollegeAdminInfo(college *models.College) error {
    // 只更新管理员相关字段
    result := d.data.DB().Model(college).Updates(map[string]interface{}{
        "admin_name":      college.AdminName,
        "admin_id_number": college.AdminIDNumber,
        "admin_phone":     college.AdminPhone,
        "admin_email":     college.AdminEmail,
    })

    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return college_error.ErrUpdateFailedError
    }
    return nil
}

// UpdateAdminAvatar 更新管理员头像
func (d *CollegeDao) UpdateAdminAvatar(collegeID uint, avatarID uint) error {
    result := d.data.DB().Model(&models.College{}).
                       Where("id = ?", collegeID).
                       Update("admin_image_id", avatarID)
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return college_error.ErrUpdateFailedError
    }
    return nil
}

// UpdateCollegeAvatar 更新学院头像
func (d *CollegeDao) UpdateCollegeAvatar(collegeID uint, avatarID uint) error {
    result := d.data.DB().Model(&models.College{}).
                       Where("id = ?", collegeID).
                       Update("college_avatar_id", avatarID)
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return college_error.ErrUpdateFailedError
    }
    return nil
}

// CheckPhoneExists 检查手机号是否已存在
func (d *CollegeDao) CheckPhoneExists(phone string, excludeID uint) (bool, error) {
    var count int64
    result := d.data.DB().Model(&models.College{}).
                       Where("admin_phone = ? AND id != ?", phone, excludeID).
                       Count(&count)
                       
    if result.Error != nil {
        return false, result.Error
    }
    return count > 0, nil
}

// CheckEmailExists 检查邮箱是否已存在
func (d *CollegeDao) CheckEmailExists(email string, excludeID uint) (bool, error) {
    var count int64
    result := d.data.DB().Model(&models.College{}).
                       Where("admin_email = ? AND id != ?", email, excludeID).
                       Count(&count)
                       
    if result.Error != nil {
        return false, result.Error
    }
    return count > 0, nil
}