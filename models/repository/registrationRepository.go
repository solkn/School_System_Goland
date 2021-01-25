package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type PsqlRegistrationRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlRegistrationRepositoryImpl(Conn *sql.DB) *PsqlRegistrationRepositoryImpl {
	return &PsqlRegistrationRepositoryImpl{conn: Conn}
}

func (pr *PsqlRegistrationRepositoryImpl) RegisterStudent(student entity.Student) error {

	_, err := pr.conn.Exec("insert into student (username,password,fname,lname,id,email,grade,img) values($1, $2, $3,$4, $5, $6,$7,$8)", student.UserName, student.Password, student.FirstName, student.LastName, student.ID, student.Email, student.Grade, student.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (pr *PsqlRegistrationRepositoryImpl) RegisterFamily(family entity.Family) error {

	_, err := pr.conn.Exec("insert into family (fname,lname,username, password, id, phone,email,img) values($1, $2, $3,$4, $5, $6,$7,$8)", family.FirstName, family.LastName, family.Username, family.Password, family.FamilyID, family.Phone, family.Email, family.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (pr *PsqlRegistrationRepositoryImpl) RegisterTeacher(teacher entity.Teacher) error {

	_, err := pr.conn.Exec("insert into teacher (username,password,phone,email,fname,lname,id,img) values($1, $2, $3,$4, $5, $6,$7,$8)", teacher.UserName, teacher.Password, teacher.Phone, teacher.Email, teacher.FirstName, teacher.LastName, teacher.TeacherID, teacher.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
func (pr *PsqlRegistrationRepositoryImpl) RegisterAdmin(admin entity.Admin) error {

	_, err := pr.conn.Exec("insert into admin (username,password,fname,lname,id,email,img) values($1, $2, $3,$4, $5, $6,$7)", admin.UserName, admin.Password, admin.FirstName, admin.LastName, admin.ID, admin.Email, admin.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
