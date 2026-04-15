package schema

type ListProjectsRequest struct {
	Path		string 	`json:"path"`
}

type GetLatestCommitRequest struct {
	ProjectName		string 	`json:"projectName"`
}

type SearchTodoRequest struct {
	ProjectName		string 	`json:"projectName"`
	Keyword			string 	`json:"keyword"`
}