package entity

type Page struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Resource    string `json:"resource"`
	Method      string `json:"method"`
}
