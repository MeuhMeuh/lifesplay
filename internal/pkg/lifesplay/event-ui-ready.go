package lifesplay

import (
	"fmt"

	"github.com/meuhmeuh/lifesplay/internal/pkg/communication"
	"github.com/meuhmeuh/lifesplay/internal/pkg/events"
	"github.com/spf13/viper"
	"google.golang.org/api/calendar/v3"
)

// uiReadyResponse defines what the body will look like for the event ui.ready.
type uiReadyResponse struct {
	FirstName string            `json:"firstName"`
	Events    []*calendar.Event `json:"events"`
}

func HandleUIReady(lifesplay *Lifesplay, payload communication.OutboundPayload) error {
	var err error

	s := uiReadyResponse{FirstName: viper.Get("me.firstName").(string)}

	// Retrieving the events to send them back straight after the UI has been booted.
	events, err := events.GetEventsOfTheDay(lifesplay.EventsClient, lifesplay.CalendarID)
	if err != nil {
		payload.SetError(fmt.Errorf("Could not retrieve the events: %s", err))
	} else {
		s.Events = events.Items
	}

	payload.SetBody(s)

	return err
}
