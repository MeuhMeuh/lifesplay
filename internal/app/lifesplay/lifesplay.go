package lifesplay

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/meuhmeuh/lifesplay/internal/pkg/events"
	"github.com/spf13/viper"

	astilectron "github.com/asticode/go-astilectron"
	astilog "github.com/asticode/go-astilog"
)

// Lifesplay defines the app wrapper.
type Lifesplay struct {
	EventsClient *http.Client
	Debug        *bool
	NoUI         *bool
	Window       *astilectron.Window
	CalendarID   string
}

// IsDebug determines if the app is in debug mode.
func (lifesplay *Lifesplay) IsDebug() bool {
	return *lifesplay.Debug
}

// HasNoUI determines if the app has booted with no UI.
func (lifesplay *Lifesplay) HasNoUI() bool {
	return *lifesplay.NoUI
}

// Initialize sets up what's needed to run the application.
func (lifesplay *Lifesplay) Initialize() {
	astilog.FlagInit()
	flag.Parse()

	var err error
	lifesplay.EventsClient, err = events.GetClient()

	lifesplay.CalendarID = viper.Get("calendar.id").(string)

	if err != nil {
		panic(fmt.Errorf("Could not initialize the events client properly : %s", err))
	}
}

// Start starts the main application.
func (lifesplay *Lifesplay) Start() {
	if !lifesplay.HasNoUI() {
		startUI(lifesplay)
	} else {
		// Simulating a loop to not kill the app.
		log.Println("App booted with no UI.")

		events.GetNextEvents(lifesplay.EventsClient, lifesplay.CalendarID)
		for {
		}
	}
}

func (lifesplay *Lifesplay) handlePostUIBoot(w []*astilectron.Window) {
	// We just use one window.
	lifesplay.Window = w[0]
}
