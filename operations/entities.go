package operations

type Book struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	Origin     string `json:"origin"`
	TotalPages int    `json:"totalPages"`
	IsRead     bool   `json:"isRead"`
}
