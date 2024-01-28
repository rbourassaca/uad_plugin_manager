package models

type (
	Config struct {
		Files files
		Path  map[string]path
	}

	files struct {
		UADSystemProfile string
	}

	path struct {
		VST2 string
		VST3 string
		AAX  string
		AU   string
	}
)