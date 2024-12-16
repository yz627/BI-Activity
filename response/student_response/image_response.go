package student_response

type ImageResponse struct {
    ID       uint   `json:"id"`
    FileName string `json:"file_name"`
    URL      string `json:"url"`
    Type     int    `json:"type"`
}