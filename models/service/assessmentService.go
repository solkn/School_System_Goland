package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type AssessmentServiceImpl struct {
	assRepo models.AssessmentRepository
}

func NewAssessmentServiceImpl(aRepo models.AssessmentRepository) *AssessmentServiceImpl {
	return &AssessmentServiceImpl{assRepo: aRepo}
}

func (as *AssessmentServiceImpl) Assessments(grade string) ([]entity.Assessment, error) {

	assessments, err := as.assRepo.Assessments(grade)

	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (as *AssessmentServiceImpl) SingleStudentAssessment(id int) ([]entity.Assessment, error) {

	assessments, err := as.assRepo.SingleStudentAssessment(id)

	if err != nil {
		return nil, err
	}

	return assessments, nil
}

func (as *AssessmentServiceImpl) Assessment(assessment entity.Assessment) ([]entity.Assessment, error) {

	assessments, err := as.assRepo.Assessment(assessment)

	if err != nil {
		return assessments, err
	}

	return assessments, nil
}

func (as *AssessmentServiceImpl) UpdateGrade(assessment entity.Assessment) error {

	err := as.assRepo.UpdateGrade(assessment)

	if err != nil {
		return err
	}

	return nil
}

func (as *AssessmentServiceImpl) DeleteGrade(studentID int, subjectID int) error {

	err := as.assRepo.DeleteGrade(studentID, subjectID)

	if err != nil {
		return err
	}

	return nil
}

func (as *AssessmentServiceImpl) DeleteGrades(studentID int) error {

	err := as.assRepo.DeleteGrades(studentID)

	if err != nil {
		return err
	}

	return nil
}

func (as *AssessmentServiceImpl) StoreGrade(assessment entity.Assessment) error {

	err := as.assRepo.StoreGrade(assessment)

	if err != nil {
		return err
	}

	return nil
}

func (as *AssessmentServiceImpl) IsQualified(studentID int) (bool, error) {

	status, err := as.assRepo.IsQualified(studentID)

	if err != nil {
		return status, err
	}

	return status, nil
}
