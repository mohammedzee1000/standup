package config

func ViewModeToString(mode uint) string {
	if mode == ViewInPanels {
		return "View in Panels"
	} else if mode == ViewSimple {
		return "View Simple"
	}
	return "Unknown or Unsupported"
}

func IsValidViewMode(mode uint) bool {
	return mode == ViewInPanels || mode == ViewSimple
}
