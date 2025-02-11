package main

import webview "github.com/webview/webview_go"

// ----------------------------------------------------------------------------
type IncrementResult struct {
	Count uint `json:"count"`
}

// ----------------------------------------------------------------------------
var currentCount = 0

// ----------------------------------------------------------------------------
func setupIncrementor(w webview.WebView) {
	// A binding that increments a value and immediately returns the new value.
	w.Bind("increment", func() IncrementResult {
		currentCount++
		return IncrementResult{Count: uint(currentCount)}
	})
}
