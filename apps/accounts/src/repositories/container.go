package repositories

import "backend/apps/accounts/libs/notification"

type Container struct {
	AccountRepository Repository
	UserRepository    Repository
	Notifications     notification.NotificationInterface
}
