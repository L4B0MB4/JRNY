package consumer

type Consumer interface {
	Initialize() error
	Consume()
}
