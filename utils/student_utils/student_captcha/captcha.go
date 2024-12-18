package student_captcha

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 定义验证码存储器
var store = base64Captcha.DefaultMemStore

// 配置验证码参数
var driverString = base64Captcha.DriverString{
    Height:          40,
    Width:           120,
    NoiseCount:      0,
    ShowLineOptions: 2,
    Length:          4,
    Source:         "1234567890", // 只使用数字
    BgColor: &color.RGBA{R: 255, G: 255, B: 255, A: 255},
    Fonts:   []string{"wqy-microhei.ttc"},
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha() (string, string, error) {
    driver := driverString.ConvertFonts()
    c := base64Captcha.NewCaptcha(driver, store)
    id, b64s, _, err := c.Generate()
    return id, b64s, err
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id, code string) bool {
    return store.Verify(id, code, true)
}