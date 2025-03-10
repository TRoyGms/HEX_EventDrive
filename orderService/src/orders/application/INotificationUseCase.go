package application

type INotificationUseCase interface {
    SendNotification(message string) error
}
