package repo

import "gorm.io/gorm"

// Service provides all databases
type Service struct {
	GDB     *gorm.DB
	HTTPLog *HTTPLogRepo
}

// NewService returns a new service instance
func NewService(gdb *gorm.DB) *Service {
	return &Service{
		GDB:     gdb,
		HTTPLog: NewHTTPLogRepo(gdb),
	}
}
