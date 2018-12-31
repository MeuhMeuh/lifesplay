package lifesplay

import (
	"fmt"
	"log"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/meuhmeuh/lifesplay/internal/pkg/events"
	"github.com/spf13/viper"
	"google.golang.org/api/calendar/v3"
)

// HandleMessages handles messages coming from the electron window.
func (lifesplay *Lifesplay) HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	o := OutboundPayload{}

	switch m.Name {
	case "ui.ready":
		type uiReady struct {
			FirstName string            `json:"firstName"`
			Events    []*calendar.Event `json:"events"`
		}
		s := uiReady{FirstName: viper.Get("me.firstName").(string)}

		log.Println(viper.Get("me.firstName").(string))

		// Retrieving the events to send them back straight after the UI has been booted.
		// In case of an error, we push it
		events, err := events.GetNextEvents(lifesplay.EventsClient, lifesplay.CalendarID)
		if err != nil {
			o.Error = fmt.Sprintf("Could not retrieve the events: %s", err)
		} else {
			s.Events = events.Items
		}

		o.Body = s

		payload = o
	}
	return
}
