package requests

type Country struct {
	Country				string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
