package requests

type Text struct {
	Country             string `json:"Country"`
	Language            string `json:"Language"`
	CountryName         string `json:"CountryName"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
