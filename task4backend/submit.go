package main

type Submit struct {
	FirstName  string   `form:"fname"`
	LastName   string   `form:"lname"`
	Email      string   `form:"email"`
	DevSide    string   `form:"fav_lang"`
	DrinkSide  []string `form:"fav_drink"`
	DrinkAbout string   `form:"about"`
	DrinkUrl   string   `form:"beer-url"`
	DrinkColor string   `form:"beer-color"`
	DrinkTime  string   `form:"beer-time"`
	FormId     string   `form:"form_id"`

	DrinkPic         string
	DrinkPicCheckSum string
}

type SubmitStorage []*Submit

var storage SubmitStorage
