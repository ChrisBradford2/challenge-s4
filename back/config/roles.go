package config

import "challenges4/models"

const (
	RoleUser      uint8 = 1 << iota
	RoleOrganizer uint8 = 1 << 1
	RoleAdmin     uint8 = 1 << 2
)

func HasRole(user *models.User, role uint8) bool {
	return user.Roles&role != 0
}
