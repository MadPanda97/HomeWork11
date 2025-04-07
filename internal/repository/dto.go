package repository

type UpdateUserRequest struct {
	ID       int
	Name     *string
	Email    *string
	Phone    *string
	Password *string
	Balance  *int64
}
