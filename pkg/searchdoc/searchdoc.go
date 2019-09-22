package searchdoc

type SearchDoc struct {
	ObjectID string `json:"objectID"`
	Title    string `json:"title"`
	File     string `json:"file"`
	Content  string `json:"content"`
}
