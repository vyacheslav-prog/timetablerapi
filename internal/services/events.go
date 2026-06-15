package services

type Event uint

type EventsSource interface {
	Events() []Event
}

func AppendEvents(sources []EventsSource) error {
	return nil
}
