package services

type Event uint

type EventsSource interface {
	Events() []Event
}

type EventsLog struct {
}

func (l EventsLog) Collect(sources []EventsSource) error {
	return nil
}
