package models

type AlertTrigger struct {
	SiteName string `json:"sitename"`
	SiteUrl string `json:"siteurl"`
	AlertType string `json:"alerttype"`
}

type Project struct {
	UserName string `json:"username"`
	ProjectName string `json:"projectname"`
	AlertTriggers []AlertTrigger `bson:"alerttriggers"` 
}