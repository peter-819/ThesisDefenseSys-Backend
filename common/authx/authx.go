package authx

import "context"

type Role int

const (
	Student   = 1 << iota
	Teacher   = 1 << iota
	Secretary = 1 << iota
	Admin     = 1 << iota
)

func Verify(ctx context.Context, role Role) bool {
	inRole := Role(0)
	switch ctx.Value("Role").(string) {
	case "Student":
		inRole = Student
	case "Teacher":
		if ctx.Value("is_secretary").(bool) {
			inRole = Secretary
		} else {
			inRole = Teacher
		}
	case "Admin":
		inRole = Admin
	}
	return (inRole | role) > 0
}
