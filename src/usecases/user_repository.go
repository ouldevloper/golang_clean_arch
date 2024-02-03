package usecases

import "go_clean/src/infrastructure"

type UserRepository struct {
	Sqlhandler *infrastructure.SqlHandler
}
