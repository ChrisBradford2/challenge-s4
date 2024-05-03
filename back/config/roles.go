package config

const (
	RoleUser      uint8 = 1 << iota
	RoleOrganizer uint8 = 1 << 1
	RoleAdmin     uint8 = 1 << 2
)
