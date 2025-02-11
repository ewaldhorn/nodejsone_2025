package main

import (
	"flag"
	"fmt"
	"hello_there/repo"

	webview "github.com/webview/webview_go"
)

// ----------------------------------------------------------------------------
const VERSION = "0.0.1"

// ----------------------------------------------------------------------------
var dataRepo *repo.DataRepo

// ----------------------------------------------------------------------------
func main() {
	if hadNoFlags() {
		// get the app window configured and ready
		// for the demo, we run in debug mode so we have dev tools at runtime
		mainView := webview.New(true)
		defer mainView.Destroy()

		// get our main view configured
		setupMainView(mainView)

		// initialise the data repo and populate it
		dataRepo = repo.NewDataRepository()
		loadDataFromCSV(DATA_File, dataRepo)

		// configure any handlers for JS actions
		setupCallbackHandlers(mainView)

		// run the app
		mainView.Run()
	}
}

// ----------------------------------------------------------------------------
// sets up the main view properties
func setupMainView(view webview.WebView) {
	view.SetTitle("Go with JavaScript Desktop Demo")
	view.SetSize(1024, 768, webview.HintNone)
	view.Init(INJECTED_JS)
	view.SetHtml(HTML_File)
}

// ----------------------------------------------------------------------------
// the caller can ask for help of version information instead of launching the app
func hadNoFlags() bool {
	help := flag.Bool("help", false, "Show help menu")
	version := flag.Bool("version", false, "Show application version")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return false
	}

	if *version {
		showVersion()
		return false
	}

	return true
}

// ----------------------------------------------------------------------------
func showVersion() {
	fmt.Println(fmt.Sprintf("Go and JavaScript Demo v%s\n", VERSION))
}
