package database

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
}

// type ProjectNames struct {

// }

// type ProjectLinks struct {
// 	AlertTriggers 
// }

// type AlertTrigger struct {
// 	SiteName string `json:"sitename"`
// 	SiteLink string `json:"sitelink"` 
// 	Alerts []string `json:"alerts"`
// }