package model

type MySqlReqArgs struct {
	Table        string                 `json:"table"`
	Where        map[string]interface{} `json:"where"`
	WhereGreater map[string]interface{} `json:"where_greater"`
	WhereLess    map[string]interface{} `json:"where_less"`
	WhereNot     map[string]interface{} `json:"where_not"`
	Data         map[string]interface{} `json:"data"`
	Limit        string                 `json:"limit"`
}

type Book struct {
	Title       string `json:"titles"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
