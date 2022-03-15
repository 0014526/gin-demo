package model

func (User) TableName() string {
	return "tbl_user"
}

type User struct {
	ID int	`gorm:"primary_key" form:"id"`
	UserId int	`form:"user_id"`
	UserName string `form:"user_name"`
	UserNike string `form:"user_nike"`
	Status int	`form:"status" default:"1"`
}
