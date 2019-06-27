package datadog

// Attributes is series metadata
type Attributes struct {
	Top struct {
		Method []string  `json:"method,omitempty"`
		Value  []float64 `json:"value,omitempty"`
		Order  []string  `json:"order,omitempty"`
	} `json:"top"`
}
