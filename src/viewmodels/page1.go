package viewmodels

import (

)

type Page1 struct {
	Title string
	Active string
	}

func GetPage1() Page1 {
	result := Page1{
		Title: "Page1",
		Active: "home",
		}
	
	return result
	}