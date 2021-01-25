package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type LoginHandler struct {
	tmpl         *template.Template
	loginService models.ProfileService
}

func NewLoginHandler(T *template.Template, PS models.ProfileService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginService: PS}
}

func (slh *LoginHandler) StudentLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		studentUser := entity.Student{}
		studentUser.UserName = r.FormValue("username")
		studentUser.Password = r.FormValue("password")

		student, err := slh.loginService.Students()

		if err != nil {

			panic(err)
		}

		for index := range student {

			uname := student[index]
			pass := student[index]

			if uname.UserName == studentUser.UserName && pass.Password == studentUser.Password {

				http.Redirect(w, r, "/student", http.StatusSeeOther)

			} else {
			}
		}

	} else {

		_ = slh.tmpl.ExecuteTemplate(w, "student.login.html", nil)

	}

}

func (slh *LoginHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering")
	if r.Method == http.MethodGet {
		slh.tmpl.ExecuteTemplate(w, "admin.login.html", nil)

	} else if r.Method == http.MethodPost {

		adminUser := entity.Admin{}
		adminUser.UserName = r.FormValue("username")
		adminUser.Password = r.FormValue("password")

		admin, err := slh.loginService.Admins()

		if err != nil {
			return
		}
		for a := range admin {

			uname := admin[a]
			pass := admin[a]

			if uname.UserName == adminUser.UserName && pass.Password == adminUser.Password {

				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}
		}

	} else {

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	}

}

func (slh *LoginHandler) TeacherLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		teacherUser := entity.Teacher{}
		teacherUser.UserName = r.FormValue("username")
		teacherUser.Password = r.FormValue("password")

		teacher, err := slh.loginService.Teachers()

		if err != nil {
			return
		}

		for t := range teacher {

			uname := teacher[t]
			pass := teacher[t]

			if uname.UserName == teacherUser.UserName && pass.Password == teacherUser.Password {

				http.Redirect(w, r, "/teacher", http.StatusSeeOther)
			}
		}

	} else {

		_ = slh.tmpl.ExecuteTemplate(w, "teacher.login.html", nil)

	}

}

func (slh *LoginHandler) FamilyLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		familyUser := entity.Family{}
		familyUser.Username = r.FormValue("username")
		familyUser.Password = r.FormValue("password")

		family, err := slh.loginService.Families()

		if err != nil {
			return
		}

		for f := range family {

			uname := family[f]
			pass := family[f]

			if uname.Username == familyUser.Username && pass.Password == familyUser.Password {

				http.Redirect(w, r, "/family", http.StatusSeeOther)
			}
		}

	} else {

		_ = slh.tmpl.ExecuteTemplate(w, "family.login.html", nil)

	}

}
