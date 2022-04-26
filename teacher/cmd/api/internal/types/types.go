// Code generated by goctl. DO NOT EDIT.
package types

type Schedule struct {
	ScheduleId string `json:"schedule_id"`
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type Teacher struct {
	Id                string     `json:"id"`
	Name              string     `json:"name"`
	IsSecretary       string     `json:"is_secretary"`
	MaxDefensePerWeek int        `json:"max_defense_per_week"`
	PreferKeywords    []string   `json:"prefer_keywords"`
	Schedules         []Schedule `json:"schedules"`
}

type QueryTeacherReq struct {
	Id string `path:"id"`
}

type QueryTeachersReq struct {
	Teachers []Teacher `json:"teachers"`
}

type ModifyTeacherReq struct {
	Id         string  `json:"id"`
	NewTeacher Teacher `json:"new_teacher"`
}