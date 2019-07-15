package services

import "backend/apps/severino/src/interfaces"

type Container struct {
	Album      interfaces.AlbumService
	Shelf      interfaces.Service
	Track      interfaces.Service
	User       interfaces.Service
	Account    interfaces.Service
	Product    interfaces.Service
	Home       interfaces.Service
	Preference interfaces.Service
	Subject    interfaces.Service
}
