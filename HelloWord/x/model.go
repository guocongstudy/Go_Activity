package main
// 名师管理表
type ShixunFamousTeacher struct {
	FamousTeacherId int `json:"famous_teacher_id"` // 名师id
	FamousTeacherPhoto string `json:"famous_teacher_photo"` // 名师头像
	FamousTeacherName string `json:"famous_teacher_name"` // 名师姓名
	FamousTeacherMobile string `json:"famous_teacher_mobile"` // 名师手机号
	FamousTeacherInstitution string `json:"famous_teacher_institution"` // 名师单位/机构
	FamousTeacherJobTitle string `json:"famous_teacher_job_title"` // 名师职称
	FamousTeacherJobTitleProve string `json:"famous_teacher_job_title_prove"` // 名师职称证明
	FamousTeacherIntroduce string `json:"famous_teacher_introduce"` // 名师介绍
	FamousTeacherCreateTime int `json:"famous_teacher_create_time"` // 名师创建时间
	FamousTeacherUpdateTime int `json:"famous_teacher_update_time"` // 名师更新时间
	FamousTeacherIshow int `json:"famous_teacher_ishow"` // 是否显示0-否，1-是
}

