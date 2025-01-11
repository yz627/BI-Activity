package collegeDAO

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/college"
	"fmt"
	"log"
	"strings"
	"time"
)

type PcDAO struct {
	// 数据库连接
	data *dao.Data
}

func NewPcDAO(data *dao.Data) *PcDAO {
	return &PcDAO{
		data: data,
	}
}

func (p *PcDAO) GetCollegeInfo(id uint) *college.CollegeInfo {
	//collegeInfo := &college.CollegeInfo{}
	//sql := fmt.Sprintf("SELECT c.id, c.college_account, c.college_name, c.campus, c.college_address, c.college_introduction, concat(i.url, i.file_name) as college_avatar_url " +
	//	"FROM college c, image i " +
	//	"WHERE c.id = ? AND c.college_avatar_id = i.id;")
	//db := p.data.DB()
	//db.Raw(sql, id).Scan(collegeInfo)
	//return collegeInfo

	db := p.data.DB()
	// 查询 college 表中的 college_avatar_id
	var clg models.College
	if err := db.Where("id = ?", id).First(&clg).Error; err != nil {
		log.Println("学院头像信息查询失败: ", err)
		return nil
	}

	// 如果 college_avatar_id 为空，先在 image 表中插入一条空记录
	if clg.CollegeAvatarID == 0 {
		var newImage models.Image
		if err := db.Create(&newImage).Error; err != nil {
			log.Println("image记录创建失败: ", err)
			return nil
		}
		// 更新 college 表中的 college_avatar_id
		if err := db.Model(&models.College{}).Where("id = ?", id).Update("college_avatar_id", newImage.ID).Error; err != nil {
			log.Println("学员头像更新失败：", err)
			return nil
		}
		clg.CollegeAvatarID = newImage.ID
	}

	// 进行多表查询
	var collegeInfo college.CollegeInfo
	sql := fmt.Sprintf("SELECT c.id, c.college_account, c.college_name, c.campus, c.college_address, c.college_introduction, concat(i.url, i.file_name) as college_avatar_url " +
		"FROM college c " +
		"JOIN image i ON c.college_avatar_id = i.id " +
		"WHERE c.id = ?")
	if err := db.Raw(sql, id).Scan(&collegeInfo).Error; err != nil {
		log.Println("学员信息查询失败：", err)
	}

	return &collegeInfo
}

func (p *PcDAO) UpdateCollegeInfo(collegeInfo *college.CollegeInfo) {
	// 1. 数据库连接实例
	db := p.data.DB()
	// 2. 更新时间
	now := time.Now()
	// 3. 更新图片表
	// 获取college_avatar_id
	college := &models.College{}
	sql1 := fmt.Sprintf("SELECT college_avatar_id FROM college where id = ?;")
	db.Raw(sql1, collegeInfo.ID).Scan(college)
	var college_avatar_id = college.CollegeAvatarID
	log.Println("college_avatar_id: ", college_avatar_id)
	// 获取file_name, url
	lastSlashIndex := strings.LastIndex(collegeInfo.CollegeAvatarUrl, "/")
	url := collegeInfo.CollegeAvatarUrl[:lastSlashIndex+1]
	file_name := collegeInfo.CollegeAvatarUrl[lastSlashIndex+1:]
	fmt.Println("路径:", url)
	fmt.Println("文件名:", file_name)
	// 更新图片表
	sql2 := fmt.Sprintf("UPDATE image SET file_name = ?, url = ?, updated_at = ? WHERE id = ?;")
	db.Exec(sql2, file_name, url, now, college_avatar_id)
	// 4. 更新学院表
	sql3 := fmt.Sprintf("UPDATE college SET campus = ?, college_address = ?, college_introduction = ?, updated_at = ? WHERE id = ? ;")
	db.Exec(sql3, collegeInfo.Campus, collegeInfo.CollegeAddress, collegeInfo.CollegeIntroduction, now, collegeInfo.ID)
}

func (p *PcDAO) GetAdminInfo(id uint) *college.AdminInfo {
	adminInfo := &college.AdminInfo{}
	sql := fmt.Sprintf("SELECT c.id, c.admin_name, c.admin_id_number, c.admin_image_id, c.admin_phone, c.admin_email, concat(i.url, i.file_name) as admin_image_url " +
		"FROM college c, image i " +
		"WHERE c.id = ? AND c.admin_image_id = i.id;")
	db := p.data.DB()
	db.Raw(sql, id).Scan(adminInfo)
	return adminInfo
}

func (p *PcDAO) UpdateAdminInfo(adminInfo *college.AdminInfo) {
	// 1. 数据库连接实例
	db := p.data.DB()
	// 2. 更新时间
	now := time.Now()
	// 3. 更新图片表
	// 获取admin_image_id
	college := &models.College{}
	sql1 := fmt.Sprintf("SELECT admin_image_id FROM college where id = ?;")
	db.Raw(sql1, adminInfo.ID).Scan(college)
	var admin_image_id = college.AdminImageID
	log.Println("admin_image_id: ", admin_image_id)
	// 获取file_name, url
	lastSlashIndex := strings.LastIndex(adminInfo.AdminImageUrl, "/")
	url := adminInfo.AdminImageUrl[:lastSlashIndex+1]
	file_name := adminInfo.AdminImageUrl[lastSlashIndex+1:]
	fmt.Println("路径:", url)
	fmt.Println("文件名:", file_name)
	// 更新图片表
	sql2 := fmt.Sprintf("UPDATE image SET file_name = ?, url = ?, updated_at = ? WHERE id = ?;")
	db.Exec(sql2, file_name, url, now, admin_image_id)
	// 4. 更新学院表
	sql3 := fmt.Sprintf("UPDATE college SET admin_name = ?, admin_id_number = ?, updated_at = ? WHERE id = ?;")
	db.Exec(sql3, adminInfo.AdminName, adminInfo.AdminIDNumber, now, adminInfo.ID)
}
