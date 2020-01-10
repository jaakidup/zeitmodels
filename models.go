package zeitmodels

// Person ...
type Person struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Age     int    `json:"age,omitempty"`
	Email   string `json:"email,omitempty"`
}

// Response is an http response wrapper
type Response struct {
	Payload interface{} `json:"payload,omitempty"`
}
