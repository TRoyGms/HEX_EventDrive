package application

type ISocketUseCase interface {
    NotifySocket(message string) error
}
