package gendata

// AmmoRequest представляет структуру для Pandora ammo
type AmmoRequest struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Tag     string            `json:"tag"`
	URI     string            `json:"uri"`
}
