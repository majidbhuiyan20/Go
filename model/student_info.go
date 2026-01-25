package model

type StudentInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	School    string `json:"school"`
	ClassName string `json:"class_name"`
	RollNo    string `json:"roll_no"`
}
