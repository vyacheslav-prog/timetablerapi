package services

type EventsSource interface {
	Events() []uint
}

func FlushEvents(sources []EventsSource) error {
	return nil
}
