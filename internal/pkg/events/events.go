package events

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
)

func GetNextEvents(eventsClient *http.Client, calendarID string) (*calendar.Events, error) {
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
	// if len(events.Items) == 0 {
	// 	fmt.Println("You're free.")
	// } else {
	// 	for _, item := range events.Items {
	// 		date := item.Start.DateTime
	// 		if date == "" {
	// 			date = item.Start.Date
	// 		}
	// 		fmt.Printf("%v (%v)\n", item.Summary, date)
	// 	}
	// }
}
