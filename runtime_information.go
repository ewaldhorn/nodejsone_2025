package main

import (
	"fmt"
	"runtime"

	webview "github.com/webview/webview_go"
)

// ----------------------------------------------------------------------------
// The JSON we want to return to the UI when called
type RuntimeImformationResult struct {
	Allocated      uint `json:"allocated"`
	TotalAllocated uint `json:"totalAllocated"`
	Reserved       uint `json:"reserved"`
	GC             uint `json:"gc"`
}

// ----------------------------------------------------------------------------
// setupRuntimeInformationCallback binds the getMemoryStats callback to the webview
// Returns runtime memory information as RuntimeImformationResult including:
// - Currently allocated memory in MB
// - Total allocated memory since start in MB
// - Reserved memory from OS in MB
// - Number of garbage collections performed
func setupRuntimeInformationCallback(w webview.WebView) {
	w.Bind("getMemoryStats", func() RuntimeImformationResult {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)

		if IS_DEBUGGING {
			// display the stats we are interested in when debugging purposes
			fmt.Printf("Allocated heap        = %v MiB\n", byteToMegabyte(stats.Alloc))
			fmt.Printf("Total Allocated       = %v MiB\n", byteToMegabyte(stats.TotalAlloc))
			fmt.Printf("Total memory reserved = %v MiB\n", byteToMegabyte(stats.Sys))
			fmt.Printf("Number of GC cycles   = %v\n", stats.NumGC)
		}

		return RuntimeImformationResult{
			Allocated:      uint(byteToMegabyte(stats.Alloc)),
			TotalAllocated: uint(byteToMegabyte(stats.TotalAlloc)),
			Reserved:       uint(byteToMegabyte(stats.Sys)),
			GC:             uint(stats.NumGC),
		}
	})
}
