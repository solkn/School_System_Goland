package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type RegistrationServiceImpl struct {
	registrationRepo models.RegistrationRepository
}

func NewRegistrationServiceImpl(regRepo models.RegistrationRepository) *RegistrationServiceImpl {

	return &RegistrationServiceImpl{registrationRepo: regRepo}
}

func (ss *RegistrationServiceImpl) RegisterStudent(student entity.Student) error {

	err := ss.registrationRepo.RegisterStudent(student)

	if err != nil {
		return err
	}

	return nil
}

func (ss *RegistrationServiceImpl) RegisterFamily(family entity.Family) error {

	err := ss.registrationRepo.RegisterFamily(family)

	if err != nil {
		return err
	}

	return nil
}

func (ss *RegistrationServiceImpl) RegisterTeacher(teacher entity.Teacher) error {

	err := ss.registrationRepo.RegisterTeacher(teacher)

	if err != nil {
		return err
	}

	return nil
}

func (ss *RegistrationServiceImpl) RegisterAdmin(admin entity.Admin) error {

	err := ss.registrationRepo.RegisterAdmin(admin)

	if err != nil {
		return err
	}

	return nil
}
