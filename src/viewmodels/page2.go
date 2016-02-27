package viewmodels

import (

)

type Page2 struct {
	Title string
	Active string
	}

func GetPage2() Page2 {
	result := Page2{
		Title: "Page2",
		Active: "home",
		}
	
	return result
	}