package handler

import (
	"html/template"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/models"
)

type HomeHandler struct {
	tmpl        *template.Template
	HomeService models.ProfileService
}

func NewHomeHandler(T *template.Template, PS models.ProfileService) *HomeHandler {
	return &HomeHandler{tmpl: T, HomeService: PS}
}

func (srh *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {

	srh.tmpl.ExecuteTemplate(w, "mainpage.layout", nil)

}

func (srh *HomeHandler) Admin(w http.ResponseWriter, r *http.Request) {

	srh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
func (srh *HomeHandler) Student(w http.ResponseWriter, r *http.Request) {

	srh.tmpl.ExecuteTemplate(w, "student.index.layout", nil)
}
func (srh *HomeHandler) Teacher(w http.ResponseWriter, r *http.Request) {

	srh.tmpl.ExecuteTemplate(w, "teacher.index.layout", nil)
}
func (srh *HomeHandler) Family(w http.ResponseWriter, r *http.Request) {

	srh.tmpl.ExecuteTemplate(w, "family.index.layout", nil)
}
