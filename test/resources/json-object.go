package resources

type Data struct {
	Names []string `json:"names"`
}

type Expected struct {
	Count int            `json:"count"`
	Valid map[string]int `json:"valid"`
}
