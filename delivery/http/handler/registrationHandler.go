package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/form"
	"github.com/Rob-a21/Cassiopeia/models"
	"github.com/Rob-a21/Cassiopeia/token"
)

//type RegistrationHandler struct {
//	tmpl           *template.Template
//	profileService    models.ProfileService
//	sessionService models.SessionService
//	userSess       *entity.Session
//	csrfSignKey    []byte
//	regService models.RegistrationService
//}

type RegistrationHandler struct {
	tmpl        *template.Template
	regService  models.RegistrationService
	csrfSignKey []byte
}

func NewRegistrationHandler(T *template.Template, RS models.RegistrationService, csKey []byte) *RegistrationHandler {
	return &RegistrationHandler{tmpl: T, regService: RS, csrfSignKey: csKey}
}

//func NewRegistrationHandler(t *template.Template, usrServ models.ProfileService,
//	sessServ models.SessionService,
//	usrSess *entity.Session, csKey []byte,RS models.RegistrationService) *RegistrationHandler{
//	return &RegistrationHandler{tmpl: t, profileService: usrServ, sessionService: sessServ,
//		userSess: usrSess, csrfSignKey: csKey,regService:RS}
//}

func (srh *RegistrationHandler) StudentRegistration(w http.ResponseWriter, r *http.Request) {

	//hashedPassword ,err:= bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
	//
	//if err != nil{
	//	panic(err)
	//}

	token, err := token.CSRFToken(srh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		srh.tmpl.ExecuteTemplate(w, "admin.register.student.layout", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		//v := url.Values{}
		//v.Add("name", r.FormValue("name"))

		//singnUpForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		//err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("username", "password", "fname", "lname", "id", "email", "grade")
		singnUpForm.MinLength("password", 8)
		singnUpForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			srh.tmpl.ExecuteTemplate(w, "admin.register.student.layout", singnUpForm)
			return
		}

		// eExists := srh.profileService.EmailExists(r.FormValue("email"))
		// if eExists {
		// 	singnUpForm.VErrors.Add("email", "Email Already Exists")
		// 	srh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			srh.tmpl.ExecuteTemplate(w, "admin.register.student.layout", singnUpForm)
			return
		}

		// role, errs := srh.userRole.RoleByName("USER")

		// if len(errs) > 0 {
		// 	singnUpForm.VErrors.Add("role", "could not assign role to the user")
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

		student := entity.Student{}
		student.UserName = r.FormValue("username")
		student.Password = string(hashedPassword)
		student.FirstName = r.FormValue("fname")
		student.LastName = r.FormValue("lname")
		student.ID, _ = strconv.Atoi(r.FormValue("id"))
		student.Email = r.FormValue("email")
		student.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		student.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		_ = srh.regService.RegisterStudent(student)

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {

		srh.tmpl.ExecuteTemplate(w, "admin.register.student.layout", nil)

	}
}

// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
// if err != nil {

// 	return
// }
// if r.Method == http.MethodPost {

// 	student := entity.Student{}
// 	student.UserName = r.FormValue("username")
// 	student.Password = string(hashedPassword)
// 	student.FirstName = r.FormValue("fname")
// 	student.LastName = r.FormValue("lname")
// 	student.ID, _ = strconv.Atoi(r.FormValue("id"))
// 	student.Grade, _ = strconv.Atoi(r.FormValue("grade"))
// 	student.Email = r.FormValue("email")

// 	mf, fh, err := r.FormFile("catimg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mf.Close()

// 	student.Image = fh.Filename

// 	writeFile(&mf, fh.Filename)

// 	srh.regService.RegisterStudent(student)

// 	if err != nil {
// 		panic(err)
// 	}

// 	http.Redirect(w, r, "/admin", http.StatusSeeOther)

// } else {

// 	srh.tmpl.ExecuteTemplate(w, "admin.register.student.layout", nil)

// }

//token, err := token.CSRFToken(srh.csrfSignKey)
//if err != nil {
//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//}
//if r.Method == http.MethodGet {
//	signUpForm := struct {
//		Values  url.Values
//		VErrors form.ValidationErrors
//		CSRF    string
//	}{
//		Values:  nil,
//		VErrors: nil,
//		CSRF:    token,
//	}
//	srh.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
//	return
//}
//
//if r.Method == http.MethodPost {
//	// Parse the form data
//	err := r.ParseForm()
//	if err != nil {
//		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//		return
//	}
//
//	// Validate the form contents
//	singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
//	singnUpForm.Required("fname", "email", "password", "confirmpassword")
//	singnUpForm.MatchesPattern("email", form.EmailRX)
//	singnUpForm.MinLength("password", 10)
//	singnUpForm.PasswordMatches("password", "confirmpassword")
//	singnUpForm.CSRF = token
//	// If there are any errors, redisplay the signup form.
//	if !singnUpForm.Valid() {
//		srh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
//		return
//	}
//
//
//	eExists := srh.profileService.EmailExists(r.FormValue("email"))
//	if eExists {
//		singnUpForm.VErrors.Add("email", "Email Already Exists")
//		srh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
//		return
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
//	if err != nil {
//		singnUpForm.VErrors.Add("password", "Password Could not be stored")
//		srh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
//		return
//	}

//role, errs := srh.S.RoleByName("USER")

//if len(errs) > 0 {
//	singnUpForm.VErrors.Add("role", "could not assign role to the user")
//	srh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
//	return
//}

//student := &entity.Student{
//	UserName: r.FormValue("fullname"),
//	Email:    r.FormValue("email"),
//	FirstName:    r.FormValue("phone"),
//	LastName:    r.FormValue("phone"),
//	Password: string(hashedPassword),
//}
//srh.regService.RegisterStudent(student)
//if len(errs) > 0 {
//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//	return
//}

//if r.Method == http.MethodPost {

//	student := entity.Student{}
//	student.UserName = r.FormValue("username")
//	student.Password = string(hashedPassword)//r.FormValue("password")
//	student.FirstName = r.FormValue("fname")
//	student.LastName = r.FormValue("lname")
//	student.ID, _ = strconv.Atoi(r.FormValue("id"))
//	student.Email = r.FormValue("email")
//
//	mf, fh, err := r.FormFile("catimg")
//	if err != nil {
//		panic(err)
//	}
//	defer mf.Close()
//
//	student.Image = fh.Filename
//
//	writeFile(&mf, fh.Filename)
//
//	srh.regService.RegisterStudent(student)
//
//	if err != nil {
//		panic(err)
//	}
//
//	http.Redirect(w, r, "/student/register", http.StatusSeeOther)
//
//} else {
//
//	srh.tmpl.ExecuteTemplate(w, "admin.register.student.html", nil)
//
//}

func (srh *RegistrationHandler) FamilyRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		family := entity.Family{}
		family.FirstName = r.FormValue("fname")
		family.LastName = r.FormValue("lname")
		family.Username = r.FormValue("username")
		family.Password = r.FormValue("password")
		family.FamilyID, _ = strconv.Atoi(r.FormValue("familyid"))
		family.Phone = r.FormValue("phone")
		family.Email = r.FormValue("email")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		family.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterFamily(family)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "admin.register.family.layout", nil)

	}
}

func (srh *RegistrationHandler) TeacherRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		teacher := entity.Teacher{}
		teacher.UserName = r.FormValue("username")
		teacher.Password = r.FormValue("password")
		teacher.Phone = r.FormValue("phone")
		teacher.Email = r.FormValue("email")
		teacher.FirstName = r.FormValue("fname")
		teacher.LastName = r.FormValue("lname")
		teacher.TeacherID, _ = strconv.Atoi(r.FormValue("teacherid"))

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		teacher.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterTeacher(teacher)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "admin.register.teacher.layout", nil)

	}
}

func (srh *RegistrationHandler) AdminRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		admin := entity.Admin{}
		admin.UserName = r.FormValue("username")
		admin.Password = r.FormValue("password")
		admin.FirstName = r.FormValue("fname")
		admin.LastName = r.FormValue("lname")
		admin.ID, _ = strconv.Atoi(r.FormValue("id"))
		admin.Email = r.FormValue("email")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		admin.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		_ = srh.regService.RegisterAdmin(admin)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {

		_ = srh.tmpl.ExecuteTemplate(w, "admin.register.admin.layout", nil)

	}
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../", "web", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
