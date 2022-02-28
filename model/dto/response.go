package dto

type ResponseHttp struct {
	Message string `json:"message"`
}

type ResponseDropdown struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type ResponseList struct {
	Count        int         `json:"count"`
	NextPage     string      `json:"next_page,omitempty"`
	PreviousPage string      `json:"previous_page,omitempty"`
	Results      interface{} `json:"results"`
}
