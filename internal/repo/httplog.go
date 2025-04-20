package repo

import (
	"aegir/internal/model"

	"gorm.io/gorm"

	repoutil "github.com/M15t/gram/pkg/util/repo"
)

// HTTPLogRepo represents the client for http log table
type HTTPLogRepo struct {
	*repoutil.Repo[model.HTTPLog]
}

// NewHTTPLogRepo returns a new http log database instance
func NewHTTPLogRepo(gdb *gorm.DB) *HTTPLogRepo {
	return &HTTPLogRepo{repoutil.NewRepo[model.HTTPLog](gdb)}
}
