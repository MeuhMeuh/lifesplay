package lifesplay

import (
	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// HandleMessages handles messages coming from the electron window.
func (lifesplay *Lifesplay) HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	payload = "hello"

	return
	// o := OutboundPayload{}

	// switch m.Name {
	// case "ui.ready":
	// 	type uiReady struct {
	// 		FirstName string
	// 		Events    []*calendar.Event
	// 	}
	// 	s := uiReady{FirstName: viper.Get("me.firstName").(string)}

	// 	// Retrieving the events to send them back straight after the UI has been booted.
	// 	// In case of an error, we push it
	// 	events, err := events.GetNextEvents(lifesplay.EventsClient, lifesplay.CalendarID)
	// 	if err != nil {
	// 		o.Error = fmt.Sprintf("Could not retrieve the events: %s", err)
	// 	} else {
	// 		s.Events = events.Items
	// 	}

	// 	payload = o
	// }
	return
}
