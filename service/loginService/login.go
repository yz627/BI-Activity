package loginService

import (
	"bi-activity/dao/loginDao"
	"bi-activity/utils"
	"errors"
)

func StudentLoginService(username, password string) (string, error) {
	// 1. 对数据进行验证、密码hash加密、权限判断等

	// 2. 调用dao层获取数据
	student, err := loginDao.GetStudentByUsername(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}
	// 3. 验证密码
	if student.Password != password {
		return "", errors.New("密码错误")
	}
	//4. 验证通过后生成token
	token, err := utils.GenerateJWT(student.ID, "student")
	if err != nil {
		return "", errors.New("无法生成token")
	}
	// 5. 返回处理后的数据
	return token, nil
}

func CollegeLoginService(username, password string) (string, error) {
	// 1. 对数据进行验证、密码hash加密、权限判断等

	// 2. 调用dao层获取数据
	college, err := loginDao.GetCollegeByUsername(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}
	// 3. 验证密码
	if college.Password != password {
		return "", errors.New("密码错误")
	}
	//4. 验证通过后生成token
	token, err := utils.GenerateJWT(college.ID, "college")
	if err != nil {
		return "", errors.New("无法生成token")
	}
	// 5. 返回处理后的数据
	return token, nil
}
