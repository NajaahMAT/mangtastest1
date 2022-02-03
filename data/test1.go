package data

type MyApplicationRegistryResponse struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result"`
}

type Test1Request struct {
	InputText string `json:"InputText"`
}

type Test1Response struct {
	Response []Pair `json:"response"`
}

type Pair struct {
	Word  string `json: "Word"`
	Count int    `json: "Count"`
}
