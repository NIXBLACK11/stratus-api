package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlertTrigger struct {
	SiteName string `json:"sitename"`
	SiteUrl string `json:"siteurl"`
	AlertType string `json:"alerttype"`
}

type Project struct {
	UserId primitive.ObjectID `json:"userid"`
	ProjectName string `json:"projectname"`
	AlertTriggers []AlertTrigger `bson:"alerttriggers"` 
}