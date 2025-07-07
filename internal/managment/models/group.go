package models

type GroupInfo struct {
	Owner       string
	Name        string
	CreatedTime string
	UpdatedTime string

	DisplayName  string
	Manager      string
	ContactEmail string
	Type         string
	ParentId     string
	IsTopGroup   bool
	Users        []string
	UsersInfo    []UserInfo

	Title     string
	Key       string
	Children  []*GroupInfo
	IsEnabled bool
}
