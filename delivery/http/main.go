package main

import (
	"database/sql"
	"time"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"github.com/Rob-a21/Cassiopeia/token"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("c:/Users/solki/go/src/github.com/Rob-a21/Cassiopeia/delivery/web/templates/*"))

func main() {

	dbconn, err := sql.Open("postgres",
		"postgres://postgres:aait@localhost/school?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	csrfSignKey := []byte(token.GenerateRandomID(32))
	//sessionRepo := repository.NewSessionRepo(dbconn)
	//sessionSrv := service.New(sessionRepo)

	registrationRepository := repository.NewPsqlRegistrationRepositoryImpl(dbconn)
	registrationService := service.NewRegistrationServiceImpl(registrationRepository)
	registrationHandler := handler.NewRegistrationHandler(tmpl, registrationService, csrfSignKey)

	profileRepository := repository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := service.NewProfileServiceImpl(profileRepository)
	profileHandler := handler.NewProfileHandler(tmpl, profileService)

	notificationRepository := repository.NewPsqlNotificationRepositoryImpl(dbconn)
	notificationService := service.NewNotificationServiceImpl(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(tmpl, notificationService)

	courseRepository := repository.NewPsqlCourseRepositoryImpl(dbconn)
	courseService := service.NewCourseServiceImpl(courseRepository)
	courseHandler := handler.NewCourseHandler(tmpl, courseService)

	attendanceRepository := repository.NewStudentAttendanceRepositoryImpl(dbconn)
	attendanceService := service.NewStudentAttendanceServiceImpl(attendanceRepository)
	attendanceHandler := handler.NewAttendanceHandler(tmpl, attendanceService)

	assessmentRepository := repository.NewAssessmentRepositoryImpl(dbconn)
	assessmentService := service.NewAssessmentServiceImpl(assessmentRepository)
	assessmentHandler := handler.NewAssessmentHandler(tmpl, assessmentService)

	homeHandler := handler.NewHomeHandler(tmpl, profileService)
	loginHandler := handler.NewLoginHandler(tmpl, profileService)
	logoutHandler := handler.NewLogoutHandler(tmpl, profileService)

	//sess := configSess()
	//uh := handler.RegistrationHandler(tmpl, registrationService, sessionSrv, sess, csrfSignKey)

	fs := http.FileServer(http.Dir("c:/Users/solki/go/src/github.com/Rob-a21/Cassiopeia/delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/home", homeHandler.Home)

	http.HandleFunc("/admin", homeHandler.Admin)
	http.HandleFunc("/student", homeHandler.Student)
	http.HandleFunc("/teacher", homeHandler.Teacher)
	http.HandleFunc("/family", homeHandler.Family)

	http.HandleFunc("/admin/login", loginHandler.AdminLogin)
	http.HandleFunc("/student/login", loginHandler.StudentLogin)
	http.HandleFunc("/teacher/login", loginHandler.TeacherLogin)
	http.HandleFunc("/family/login", loginHandler.FamilyLogin)

	http.HandleFunc("/logout", logoutHandler.Logout)

	// student handler
	http.HandleFunc("/student/profile", profileHandler.StudentProfile)
	http.HandleFunc("/student/year/registration", profileHandler.NewYearRegistration)
	http.HandleFunc("/student/assessment", assessmentHandler.StudentAssessment)
	http.HandleFunc("/student/course", courseHandler.StudentCourse)
	http.HandleFunc("/student/notification", notificationHandler.StudentGetNotification)
	http.HandleFunc("/student/attendance/new", attendanceHandler.FillStudentAttendance)
	http.HandleFunc("/student/attendance/check", attendanceHandler.CheckStudentAttendance)

	// teacher handlers

	http.HandleFunc("/teacher/profile", profileHandler.TeacherProfile)
	http.HandleFunc("/teacher/notification", notificationHandler.TeacherAddNotification)
	http.HandleFunc("/teacher/assessment/new", assessmentHandler.StoreGrade)
	http.HandleFunc("/teacher/assessment/update", assessmentHandler.UpdateGrade)
	http.HandleFunc("/teacher/assessment/delete", assessmentHandler.DeleteGrade)
	http.HandleFunc("/teacher/assessments/delete", assessmentHandler.DeleteGrades)

	// family handlers
	http.HandleFunc("/family/profile", profileHandler.FamilyProfile)
	http.HandleFunc("/family/assessment", assessmentHandler.StudentAssessment)
	http.HandleFunc("/family/attendance/check", attendanceHandler.CheckStudentAttendance)

	// admin handlers

	http.HandleFunc("/admin/register/admin", registrationHandler.AdminRegistration)
	http.HandleFunc("/admin/register/student", registrationHandler.StudentRegistration)
	http.HandleFunc("/admin/register/teacher", registrationHandler.TeacherRegistration)
	http.HandleFunc("/admin/register/family", registrationHandler.FamilyRegistration)
	http.HandleFunc("/admin/profile", profileHandler.AdminProfile)
	http.HandleFunc("/admin/student", profileHandler.AdminGetStudent)
	http.HandleFunc("/admin/student/delete", profileHandler.AdminDeleteStudent)
	http.HandleFunc("/admin/teacher", profileHandler.AdminGetTeacher)
	http.HandleFunc("/admin/teacher/delete", profileHandler.AdminDeleteTeacher)
	http.HandleFunc("/admin/course", courseHandler.AdminGetCourse)
	http.HandleFunc("/admin/course/new", courseHandler.AdminAddCourse)

	http.ListenAndServe(":8181", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := token.GenerateRandomID(32)
	signingString, err := token.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
