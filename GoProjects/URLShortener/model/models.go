package model

type UrlObject struct {
	LongUrl   string `json:"longurl" bson:"longurl"`
	ShortCode string `json:"shortcode" bson:"shortcode"`
}
