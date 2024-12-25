package student_mask

import "strings"

func MaskPhone(phone string) string {
    if len(phone) != 11 {
        return phone
    }
    return phone[:3] + "****" + phone[7:]
}

func MaskEmail(email string) string {
    if email == "" {
        return ""
    }
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return email
    }
    name := parts[0]
    if len(name) < 2 {
        return email
    }
    return name[:1] + "****" + "@" + parts[1]
}