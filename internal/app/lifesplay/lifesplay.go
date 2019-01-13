package lifesplay

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/meuhmeuh/lifesplay/internal/pkg/communication"
	"github.com/meuhmeuh/lifesplay/internal/pkg/events"
	"github.com/spf13/viper"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
)

// Lifesplay defines the app wrapper.
type Lifesplay struct {
	EventsClient *http.Client
	Debug        *bool
	Logger       *logrus.Logger
	Window       *astilectron.Window
	CalendarID   string
}

// IsDebug determines if the app is in debug mode.
func (lifesplay *Lifesplay) IsDebug() bool {
	return *lifesplay.Debug
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
	startUI(lifesplay)
}

func (lifesplay *Lifesplay) handlePostUIBoot(w []*astilectron.Window) {
	// We just use one window.
	lifesplay.Window = w[0]
}

// HandleMessages handles messages coming from the electron window.
func (lifesplay *Lifesplay) HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	payload := &communication.OutboundPayloadImpl{}
	var err error

	switch m.Name {
	case "ui.ready":
		err = handleUIReady(lifesplay, payload)
	}

	return payload, err
}
