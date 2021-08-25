package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // interface = any type
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
