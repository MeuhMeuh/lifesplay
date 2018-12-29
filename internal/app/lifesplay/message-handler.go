package lifesplay

import (
	"log"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// HandleMessages handles messages coming from the electron window.
func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "event.name":
		// Unmarshal payload
		log.Println(m.Payload)
		payload = "Hello from Go!"
	}
	return
}
