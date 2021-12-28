package models

type Directory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	Files []File `json:"files"`
}

type File struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
