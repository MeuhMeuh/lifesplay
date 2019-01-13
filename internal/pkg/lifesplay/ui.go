package lifesplay

import (
	"log"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// StartUI boots the Electron app.
func StartUI(lifesplay *Lifesplay) {
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            "Lifesplay",
			AppIconDarwinPath:  "resources/logo.icns",
			AppIconDefaultPath: "resources/logo.png",
		},
		Debug: lifesplay.IsDebug(),
		Windows: []*bootstrap.Window{&bootstrap.Window{
			Homepage: "index.html",
			Options: &astilectron.WindowOptions{
				TitleBarStyle: astilectron.TitleBarStyleHidden,
				// Fullscreen:      astilectron.PtrBool(true), /* Will be great for the Rasp later. */
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Width:           astilectron.PtrInt(800),
				Height:          astilectron.PtrInt(480),
			},

			MessageHandler: lifesplay.HandleMessages,
		}},
		OnWait: func(_ *astilectron.Astilectron, w []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			if *lifesplay.Debug {
				w[0].OpenDevTools()
			}

			lifesplay.handlePostUIBoot(w)
			return nil
		},
	}); err != nil {
		log.Println("Failed bootstrapping:", err)
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
