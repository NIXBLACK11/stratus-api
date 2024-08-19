package models

type AlertTrigger struct {
	SiteName string `json:"sitename"`
	SiteUrl string `json:"siteurl"`
	AlertType string `json:"alerttype"`
}

type Project struct {
	ProjectName string `bson:"projectname"`
	AlertTriggers []AlertTrigger `bson:"alerttriggers"` 
}