package repositories

import "new-platform-dashboard/db/entities"

func Save(entity *entities.Example) (*[]entities.Example, error) {
	return &[]entities.Example{*entity}, nil
}
