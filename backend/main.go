package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "backend/migrations"
	_ "backend/tests"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	// Do not automigrate if we are not using the default dataDir (i.e. testing)
	// enableAutomigrate := app.DataDir() == "pb_data"
	enableAutomigrate := false

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun && enableAutomigrate,
	})

	app.OnRecordBeforeCreateRequest("permissions").Add(func(e *core.RecordCreateEvent) error {
		org := e.Record.Get("org")
		group := e.Record.Get("group")

		if org == "" && group == "" {
			return fmt.Errorf("An organization or group must be provided")
		}

		if org != "" && group != "" {
			return fmt.Errorf("Either an organization or a group must be provided, but not both")
		}

		return nil
	})

	app.OnRecordBeforeUpdateRequest("permissions").Add(func(e *core.RecordUpdateEvent) error {
		org := e.Record.Get("org")
		group := e.Record.Get("group")

		if org == "" && group == "" {
			return fmt.Errorf("An organization or group must be provided")
		}

		if org != "" && group != "" {
			return fmt.Errorf("Either an organization or a group must be provided, but not both")
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
