package data

import "go-api-gin-swagger/model"

// albums slice to seed record album data.
var Testalbums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var TestserviceCodes = []model.ServiceCode{
	{ID: 1, Name: "SVC01", Description: "super svc01", BusinessUnit: "BFI"},
	{ID: 2, Name: "SVC02", Description: "super svc02", BusinessUnit: "BSX"},
	{ID: 3, Name: "SVC03", Description: "super svc03", BusinessUnit: "BSS"},
}
