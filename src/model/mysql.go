package model

type MySqlRequestData struct {
	Table        string                 `json:"table"`
	Where        map[string]interface{} `json:"where"`
	WhereGreater map[string]interface{} `json:"where_greater"`
	WhereLess    map[string]interface{} `json:"where_less"`
	WhereNot     map[string]interface{} `json:"where_not"`
	Data         map[string]interface{} `json:"data"`
	Limit        string                 `json:"limit"`
}
