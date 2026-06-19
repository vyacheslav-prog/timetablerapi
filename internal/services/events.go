package services

type EventsSource interface {
	Events() []uint
}

type EventsLog struct {
}

func (l EventsLog) Collect(sources []EventsSource) error {
	return nil
}
