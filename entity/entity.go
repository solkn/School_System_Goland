package entity

import "time"

type Admin struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Image     string `json:"img"`
}

//Student struct for data caching
type Student struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Grade     int    `json:"grade"`
	Image     string `json:"img"`
}

//Teacher struct for data caching
type Teacher struct {
	UserName  string `json:"username"`
	Password  string `json:"Password"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	TeacherID int    `json:"TeacherID"`
	Image     string `json:"Image"`
}

//Family struct for data caching
type Family struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Username  string `json:"username"`
	Password  string `json:"Password"`
	FamilyID  int    `json:"FamilyID"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Image     string `json:"Image"`
}

type Assessment struct {
	Value     int `json:"value"`
	SubjectID int `json:"subjectid"`
	StudentID int `json:"studentid"`
	Grade     int `json:"grade"`
}

type Attendance struct {
	Date      time.Time `json:"date"`
	StudentID int       `json:"studentid"`
}

type Course struct {
	CourseID   int    `json:"courseid"`
	CourseName string `json:"coursename"`
	Grade      int    `json:"grade"`
}

type Notification struct {
	Message          string    `json:"message"`
	NotifyName       string    `json:"name"`
	NotificationDate time.Time `json:"date"`
}

type Session struct {
	ID         uint   `json:"ID"`
	UUID       string `json:"uuid"`
	Expires    int64  `json:"expires"`
	SigningKey []byte `json:"signinkey"`
}
