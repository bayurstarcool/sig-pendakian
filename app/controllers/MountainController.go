package controllers

import (
	"html/template"
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/helpers"
	"github.com/bayurstarcool/BayurGo/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type M map[string]interface{}

func MountainCreate(w http.ResponseWriter, r *http.Request) {
	tpls := []string{"views/layouts/backend.html", "views/backend/mountains/create.html"}
	rnd.Template(w, http.StatusOK, tpls, nil)
}
func (c *AppContext) MountainStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	lat := r.FormValue("lat")
	lng := r.FormValue("lng")
	// note := r.FormValue("note")
	Mountain := models.Mountain{Name: name, Lat: lat, Lng: lng}
	c.DB.Create(&Mountain)
	http.Redirect(w, r, "/mountains", http.StatusSeeOther)
}
func (c *AppContext) TrackStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	mid := r.FormValue("mid")
	lat := r.FormValue("lat")
	lng := r.FormValue("lng")
	// note := r.FormValue("note")
	track := models.Track{Name: name, Latitude: lat, Longitude: lng, MountainID: mid}
	c.DB.Create(&track)
	http.Redirect(w, r, "/mountains", http.StatusSeeOther)
}
func (c *AppContext) MountainUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	// note := r.FormValue("note")
	// password := r.FormValue("password")
	Mountain := models.Mountain{Name: name}
	c.DB.Updates(&Mountain)
	http.Redirect(w, r, "/mountains", http.StatusSeeOther)
}
func (c *AppContext) MountainEdit(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName("id")
	db := c.DB
	Mountain := models.Mountain{}
	db.Where("id = ?", id).First(&Mountain)
	tpls := []string{"views/layouts/backend.html", "views/backend/mountains/edit.html"}
	rnd.Template(w, http.StatusOK, tpls, Mountain)
}
func (c *AppContext) TrackCreate(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName("id")
	db := c.DB
	mountain := models.Mountain{}
	db.Where("id = ?", id).First(&mountain)
	tpls := []string{"views/layouts/backend.html", "views/backend/mountains/tracks/create.html"}
	rnd.FuncMap(template.FuncMap{
		"inc": helpers.GetIncrement,
	})
	rnd.Template(w, http.StatusOK, tpls, mountain)
}
func (c *AppContext) TrackIndex(w http.ResponseWriter, r *http.Request) {

	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName("id")
	db := c.DB
	mountain := models.Mountain{}
	db.Where("id = ?", id).First(&mountain)
	tpls := []string{"views/layouts/backend.html", "views/backend/mountains/tracks/index.html"}
	rnd.FuncMap(template.FuncMap{
		"inc": helpers.GetIncrement,
	})
	tracks := []models.Track{}
	db.Find(&tracks)
	rnd.Template(w, http.StatusOK, tpls, M{
		// We can pass as many things as we like
		"mountain": mountain,
		"tracks":   tracks, // Just an example
	})
}
func (c *AppContext) MountainIndex(w http.ResponseWriter, r *http.Request) {
	db := c.DB
	mountains := []models.Mountain{}
	db.Find(&mountains)
	tpls := []string{"views/layouts/backend.html", "views/backend/mountains/index.html"}
	rnd.FuncMap(template.FuncMap{
		"inc": helpers.GetIncrement,
	})
	rnd.Template(w, http.StatusOK, tpls, mountains)
}
