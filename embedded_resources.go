package main

// We embed all the required resources directly into the binary, making it easier
// to distribute the application.

import (
	_ "embed"
)

// ----------------------------------------------------------------------------
//
//go:embed ui/dist/index.html
var HTML_File string

// ----------------------------------------------------------------------------
//
//go:embed data/dataload.csv
var DATA_File string

// ----------------------------------------------------------------------------
//
//go:embed templates/peopleTableTemplate.html
var PEOPLE_TEMPLATE string

// ----------------------------------------------------------------------------
//
//go:embed data/injected.js
var INJECTED_JS string
