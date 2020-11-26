package model

type (
	// TodoModel struct for each todo
	TodoModel struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
