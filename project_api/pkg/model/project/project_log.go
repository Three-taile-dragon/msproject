package project

type ProjectLog struct {
	Content      string `json:"content"`
	Remark       string `json:"remark"`
	CreateTime   string `json:"create_time"`
	SourceCode   string `json:"source_code"`
	ProjectCode  string `json:"project_code"`
	IsComment    int    `json:"is_comment"`
	ProjectName  string `json:"project_name"`
	MemberAvatar string `json:"member_avatar"`
	MemberName   string `json:"member_name"`
	TaskName     string `json:"task_name"`
}
