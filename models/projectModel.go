package models

type AlertTrigger struct {
	SiteName string `json:"sitename"`
	SiteUrl string `json:"siteurl"`
	AlertType []string `json:"alerttype"`
}

type Project struct {
	UserName string `json:"username"`
	ProjectName string `json:"projectname"`
	Tries int `json:"tries"`
	AlertTriggers []AlertTrigger `bson:"alerttriggers"` 
}