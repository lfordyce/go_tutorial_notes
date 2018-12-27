package json_utils

type dimension struct {
	Height int
	Width  int
}

type Dog struct {
	Breed    string
	WeightKg int
	// Now `size` is a pointer to a `dimension` instance
	Size *dimension `json:",omitempty"`
}

type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}
