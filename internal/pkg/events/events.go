package events

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
)

// GetEventsOfTheDay retrieves the list of events for a given calendar ID
// for the whole day ([0:00 -> 23:59:59.99999999] interval).
func GetEventsOfTheDay(eventsClient *http.Client, calendarID string) (*calendar.Events, error) {
	srv, err := calendar.New(eventsClient)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	now := time.Now()
	minDate := beginningOfDay(now)
	maxDate := endOfDay(now)

	return srv.
		Events.
		List(calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(minDate.Format(time.RFC3339)).
		TimeMax(maxDate.Format(time.RFC3339)).
		MaxResults(10).
		OrderBy("startTime").
		Do()
}
