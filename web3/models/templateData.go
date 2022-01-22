package models

type TemplateData struct {
	StrMap      map[string]string
	TntMap      map[string]int
	FlotMap     map[string]float32
	Data        map[string]interface{}
	CRFToken    string
	FlasMessage string
	Warning     string
	Error       string
}
