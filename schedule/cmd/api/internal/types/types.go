// Code generated by goctl. DO NOT EDIT.
package types

type Member struct {
	TeacherID   string `json:"teacher_id"`
	TeacherName string `json:"teacher_name,optional"`
}

type Student struct {
	StudentID   string `json:"student_id"`
	StudentName string `json:"student_name,optional"`
}

type Defense struct {
	Id                string    `json:"id,optional"`
	Classroom         string    `json:"classroom,optional"`
	ClassroomFullName string    `json:"classroom_full_name,optional"`
	StartTime         string    `json:"start_time,optional"`
	EndTime           string    `json:"end_time,optional"`
	Committee         []Member  `json:"committee,optional"`
	Students          []Student `json:"students, optional"`
}

type AddDefenseReply struct {
	Id string `json:"id"`
}

type QueryDefenseReq struct {
	Id string `path:"id"`
}

type QueryDefenseReply struct {
	Defenses []Defense `json:"defenses"`
}

type ModifyDefenseReq struct {
	ModifyID   string  `json:"modify_id"`
	NewDefense Defense `json:"new_defense"`
}

type RemoveDefenseReq struct {
	ID string `json:"id"`
}
