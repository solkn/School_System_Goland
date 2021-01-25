package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type AssessmentRepositoryImpl struct {
	conn *sql.DB
}

func NewAssessmentRepositoryImpl(Conn *sql.DB) *AssessmentRepositoryImpl {
	return &AssessmentRepositoryImpl{conn: Conn}
}

func (as *AssessmentRepositoryImpl) Assessments(grade string) ([]entity.Assessment, error) {

	rows, err := as.conn.Query("SELECT * FROM assessment WHERE grade=$1", grade)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	results := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.Grade, &result.Value, &result.SubjectID, &result.StudentID)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

//SingleStudentAssessment
func (as *AssessmentRepositoryImpl) SingleStudentAssessment(id int) ([]entity.Assessment, error) {

	rows, err := as.conn.Query("SELECT * FROM assessment WHERE studentid=$1", id)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	results := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.Value, &result.SubjectID, &result.StudentID, &result.Grade)
		if err != nil {
			return nil, err
		}
		results = append(results, result)

	}

	return results, nil
}

func (as *AssessmentRepositoryImpl) Assessment(assessment entity.Assessment) ([]entity.Assessment, error) {

	rows, err := as.conn.Query("SELECT * FROM assessment WHERE studentid=$1 AND subjectid=$2", assessment.StudentID, assessment.SubjectID)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	results := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.Grade, &result.Value, &result.SubjectID, &result.StudentID)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (as *AssessmentRepositoryImpl) UpdateGrade(assessment entity.Assessment) error {

	_, err := as.conn.Exec("UPDATE assessment SET value=$1,subjectid=$2, studentid=$3, grade=$4 WHERE id=$3", assessment.Value, assessment.SubjectID, assessment.StudentID, assessment.Grade)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

func (as *AssessmentRepositoryImpl) DeleteGrade(studentID int, subjectID int) error {

	_, err := as.conn.Exec("DELETE FROM assessment WHERE studentid=$1 AND subjectid=$2", studentID, subjectID)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (as *AssessmentRepositoryImpl) DeleteGrades(studentID int) error {

	_, err := as.conn.Exec("DELETE FROM assessment WHERE studentid=$1", studentID)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (as *AssessmentRepositoryImpl) StoreGrade(assessment entity.Assessment) error {

	_, err := as.conn.Exec("INSERT INTO assessment (value,studentid,subjectid,grade) values($1, $2, $3, $4)", assessment.Value, assessment.StudentID, assessment.SubjectID, assessment.Grade)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (as *AssessmentRepositoryImpl) IsQualified(studentID int) (bool, error) {

	rows, _ := as.SingleStudentAssessment(studentID)

	total := 0
	num := 0
	for _, element := range rows {
		total += element.Value
		num += 1
	}

	average := total / num

	if average >= 50 {
		return true, nil
	} else {
		return false, nil
	}

}
