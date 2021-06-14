package repository

import (
	"com.snap.co/sample/models"
	"encoding/json"
	"github.com/imroc/req"
)

const (
	url = "https://api.divar.ir/v8/search/1/transport"
)

func GetDivarByCategory(page int) *models.Divar {
	divar := new(models.Divar)
var data =SetPage(page)
//	req.Debug = true
	r, _ := req.Post(url, data)
	if r != nil {
	}
	r.ToJSON(divar)
	return divar
}
func SetPage(page int) string {
	var input = new(models.InputDivar)
	input.Page = page
	input.JSONSchema.Category.Value="transport"
	var jn, _ =json.Marshal(input)
	return string(jn)

}
