package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"github.com/Rob-a21/Cassiopeia/token"
)



//////////-------------------Profile----------------////////////////
//////////////////////////////////
////////////////////////////////
////////////////////////////////
///////////////////////////////////////
func TestAdminGetTeacher(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	getteacherrepo := repository.NewprofileRepoMock(nil)
	getteacherserv := service.NewProfileServiceImpl(getteacherrepo)

	adminGetTeacherHandler := handler.NewProfileHandler(tmpl, getteacherserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/teacher", adminGetTeacherHandler.AdminGetTeacher)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/admin/teacher")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("teacher")) {
		t.Errorf("want body to contain %q", body)
	}
}


func TestAdminGetStudent(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	profrepo := repository.NewprofileRepoMock(nil)
	profServ := service.NewProfileServiceImpl(profrepo)

	adminGetStudentHandler := handler.NewProfileHandler(tmpl, profServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/student", adminGetStudentHandler.AdminGetStudent)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/admin/student")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("username")) {
		t.Errorf("want body to contain %q", body)
	}
}


////////-------------------registration----------------////////////////

func TestStudentRegistration(t *testing.T) {

	csrfSignKey := []byte(token.GenerateRandomID(32))

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	regRepo := repository.NewregistrationMockRepo(nil)
	regServ := service.NewRegistrationServiceImpl(regRepo)

	adminRegistrationHandler := handler.NewRegistrationHandler(tmpl, regServ, csrfSignKey)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/register/student", adminRegistrationHandler.StudentRegistration)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("username", entity.StudentMock.UserName)
	form.Add("studentID", string(entity.StudentMock.ID))
	form.Add("password", entity.StudentMock.Password)
	form.Add("firstname", entity.StudentMock.FirstName)
	form.Add("lastname", entity.StudentMock.LastName)
	form.Add("grade", string(entity.StudentMock.Grade))
	form.Add("email", entity.StudentMock.Email)

	resp, err := tc.PostForm(sURL+"/admin/register/student", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("username")) {
		t.Errorf("want body to contain %q", body)
	}
}



func TestTeacherPostAssessment(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	assRepo := repository.NewAssessmentMockRepository(nil)
	assServ := service.NewAssessmentServiceImpl(assRepo)

	adminAssessmentHandler := handler.NewAssessmentHandler(tmpl, assServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/teacher/assessment/new", adminAssessmentHandler.StoreGrade)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("value", string(entity.AssessmentMock.Value))
	form.Add("studentID", string(entity.AssessmentMock.StudentID))
	form.Add("subjectID", string(entity.AssessmentMock.SubjectID))
	form.Add("grade", string(entity.AssessmentMock.Grade))

	resp, err := tc.PostForm(sURL+"/teacher/assessment/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("")) {
		t.Errorf("want body to contain %q", body)
	}

}

// func TestTeacherPostNotification(t *testing.T) {

// 	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

// 	attRepo := repository.NewNotificationRepoMock(nil)
// 	attServ := service.NewNotificationServiceImpl(attRepo)

// 	teacherNotificationHandler := handler.NewNotificationHandler(tmpl, attServ)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/teacher/notification", teacherNotificationHandler.TeacherAddNotification)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	sURL := ts.URL

// 	form := url.Values{}
// 	form.Add("message", entity.NotificationMock.Message)
// 	form.Add("notifyname", entity.NotificationMock.NotifyName)
// 	form.Add("date", entity.NotificationMock.NotificationDate.String())

// 	resp, err := tc.PostForm(sURL+"/teacher/notification", form)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !bytes.Contains(body, []byte("")) {
// 		t.Errorf("want body to contain %q", body)
// 	}

// }

///////////////////////////

// func TestStudentAssessment(t *testing.T) {

// 	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

// 	assRepo := repository.NewAssessmentMockRepository(nil)
// 	assServ := service.NewAssessmentServiceImpl(assRepo)

// 	adminAssessmentHandler := handler.NewAssessmentHandler(tmpl, assServ)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/student/assessment", adminAssessmentHandler.StudentAssessment)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	sURL := ts.URL

// 	resp, err := tc.Get(sURL + "/student/assessment")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	//////////////////////////////////////////////////////
// 	if !bytes.Contains(body, []byte("")) {
// 		t.Errorf("want body to contain %q", body)
// 	}
// }

//////////-----------AttendanceHandlerTest-------------///////////////////

// func TestStudentPostAttendance(t *testing.T) {

// 	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

// 	attRepo := repository.NewAttendanceRepoMock(nil)
// 	attServ := service.NewStudentAttendanceServiceImpl(attRepo)

// 	studentAttendanceHandler := handler.NewAttendanceHandler(tmpl, attServ)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/student/attendance/new", studentAttendanceHandler.FillStudentAttendance)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	sURL := ts.URL

// 	form := url.Values{}
// 	form.Add("Date", entity.AttendanceMock.Date.String())
// 	form.Add("studentID", string(entity.AttendanceMock.StudentID))

// 	resp, err := tc.PostForm(sURL+"/student/attendance/new", form)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !bytes.Contains(body, []byte("")) {
// 		t.Errorf("want body to contain %q", body)
// 	}

// }

///////// Notification Testing  ////////


