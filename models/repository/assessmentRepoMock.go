package repository

import (
	"database/sql"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)


type AssessmentRepoMock struct {
	conn *sql.DB
}

func NewAssessmentMockRepository(db *sql.DB) models.AssessmentRepository {
	return &AssessmentRepoMock{conn: db}
}

func (aRepo *AssessmentRepoMock) SingleStudentAssessment(id int) ([]entity.Assessment, error) {
	aMock := []entity.Assessment{entity.AssessmentMock}

	if id == 0001 {
		return aMock, nil
	}

	return nil, nil
}

func (aRepo *AssessmentRepoMock) UpdateGrade(assessment entity.Assessment) error {
	assessment = entity.AssessmentMock

	return nil
}

func (aRepo *AssessmentRepoMock) DeleteGrade(studentID int, subjectID int) error {
	studentID = entity.AssessmentMock.StudentID
	subjectID = entity.AssessmentMock.SubjectID

	if studentID != 0001 {
		return nil
	}

	return nil
}

func (aRepo *AssessmentRepoMock) DeleteGrades(studentID int) error {
	panic("implement me")
}

func (aRepo *AssessmentRepoMock) StoreGrade(assessment entity.Assessment) error {
	assessment = entity.AssessmentMock

	return nil
}

func (aRepo *AssessmentRepoMock) IsQualified(studentID int) (bool, error) {
	panic("implement me")
}

func (aRepo *AssessmentRepoMock) Assessments(grade string) ([]entity.Assessment, error) {

	posts := []entity.Assessment{entity.AssessmentMock}

	return posts, nil
}

func (aRepo *AssessmentRepoMock) Assessment(assessment entity.Assessment) ([]entity.Assessment, error) {
	aMock := []entity.Assessment{entity.AssessmentMock}

	if assessment.StudentID == 0001 {
		return aMock, nil
	}

	return nil, nil

}




