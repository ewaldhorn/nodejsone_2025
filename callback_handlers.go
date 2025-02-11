package main

import (
	"hello_there/components"
	"runtime"

	webview "github.com/webview/webview_go"
)

// ----------------------------------------------------------------------------
// configure the various specific handlers uses to communicate with the UI
func setupCallbackHandlers(view webview.WebView) {
	setupQuitHandler(view)
	setupGCHandler(view)
	setupJSEvalHandler(view)
	setupIncrementor(view)
	setupRuntimeInformationCallback(view)
	setupPeopleCountHandler(view)
	setupPopulateTableHandler(view)
}

// ----------------------------------------------------------------------------
// allows us to quit the app from the ui side if needed
func setupQuitHandler(w webview.WebView) {
	w.Bind("quit", func() {
		w.Terminate()
	})
}

// ----------------------------------------------------------------------------
// allows us to request a garbage collection event from the ui
func setupGCHandler(w webview.WebView) {
	w.Bind("gc", func() {
		runtime.GC()
	})
}

// ----------------------------------------------------------------------------
// allows for the execution of JS code from the UI side
// this feature allows us to sanitise code before running eval on it if we wish
// in this demo, we just execute the code to show it is possible
func setupJSEvalHandler(w webview.WebView) {
	w.Bind("runJS", func(args ...string) {
		w.Eval(args[0])
		return
	})
}

// ----------------------------------------------------------------------------
// enables the ui to call for the people table to be constructed from the Go
// side of things
func setupPopulateTableHandler(w webview.WebView) {
	w.Bind("populateTable", func() string {
		return components.BuildPeopleTable(&dataRepo.People, PEOPLE_TEMPLATE)
	})
}

// ----------------------------------------------------------------------------
// retrieve the number of elements in the people list
func setupPeopleCountHandler(w webview.WebView) {
	w.Bind("getPeopleCount", func() PeopleCountResult {
		return PeopleCountResult{Count: uint(dataRepo.PeopleCount())}
	})
}
