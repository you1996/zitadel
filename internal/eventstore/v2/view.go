package eventstore

import "time"

//View is the minimum representation of a View model.
// It implements a basic reducer
// it might be saved in a database or in memory
type View struct {
	CreationDate      time.Time     `gorm:"column:creation_date"`
	ChangeDate        time.Time     `gorm:"column:change_date"`
	ProcessedSequence uint64        `gorm:"column:sequence"`
	Events            []EventReader `json:"-"`
}

//AppendEvents adds all the events to the read model.
// The function doesn't compute the new state of the read model
func (rm *View) AppendEvents(events ...EventReader) *View {
	rm.Events = append(rm.Events, events...)
	return rm
}

//ReduceEvent sets the data
func (rm *View) ReduceEvent(event EventReader) {
	if rm.CreationDate.IsZero() {
		rm.CreationDate = rm.Events[0].CreationDate()
	}
	rm.ChangeDate = event.CreationDate()
	rm.ProcessedSequence = event.Sequence()
}
