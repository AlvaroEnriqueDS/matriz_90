package matriz_90

type model struct {
	Matriz [][]int `json:"Matriz"`
}

type service interface {
	processService(data *model) ([][]int, error)
}