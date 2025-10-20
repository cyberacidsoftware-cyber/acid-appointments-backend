package usecase

import (
	"backend/internal/infra"

	"github.com/google/uuid"
)

type SpecialistWithServicesInfoArgs struct {
	Specialist SpecialistInfoArgs
	Services   []SpecialistServiceInfoArgs
}

func (args *SpecialistWithServicesInfoArgs) Validate() *UsecaseError {
	if err := args.Specialist.Validate(); err != nil {
		return err
	}
	for idx := range args.Services {
		s := &args.Services[idx]
		if err := s.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// OBS: Should be in a transaction
func CreateSpecialistWithServices(state State, args SpecialistWithServicesInfoArgs) (uuid.UUID, *UsecaseError) {
	var specialistID uuid.UUID
	{
		id, err := CreateSpecialist(state, args.Specialist)
		if err != nil {
			return specialistID, err
		}
		specialistID = id
	}
	{
		for _, s := range args.Services {
			s.SpecialistID = specialistID
			_, err := CreateSpecialistService(state, s)
			if err != nil {
				return specialistID, err
			}
		}
	}
	return specialistID, nil
}

// OBS: Should be in a transaction
func UpdateSpecialistWithServices(state State, specialistID uuid.UUID, args SpecialistWithServicesInfoArgs) (infra.Specialist, *UsecaseError) {
	var none infra.Specialist
	specialist, err := UpdateSpecialist(state, specialistID, args.Specialist)
	if err != nil {
		return none, err
	}
	{
		services, err := GetServicesBySpecialistID(state, specialistID)
		if err != nil {
			return none, err
		}
		for _, serviceArg := range args.Services {
			found := false
			for _, service := range services {
				if service.ServiceNameID == serviceArg.ServiceNameID {
					// Update
					// TODO: Should I run a diff check and only update when there are differences?
					_, err := UpdateSpecialistService(state, service.ID, serviceArg)
					if err != nil {
						return none, err
					}
					found = true
					break
				}
			}
			if !found {
				// Create
				_, err := CreateSpecialistService(state, SpecialistServiceInfoArgs{
					SpecialistID:  specialist.ID,
					ServiceNameID: serviceArg.ServiceNameID,
					Price:         serviceArg.Price,
					DurationMin:   serviceArg.DurationMin,
					Duration:      serviceArg.Duration,
				})
				if err != nil {
					return none, err
				}
			}
		}
		for _, service := range services {
			found := false
			for _, serviceArg := range args.Services {
				if service.ServiceNameID == serviceArg.ServiceNameID {
					// Update (Already done)
					found = true
					break
				}
			}
			if !found {
				// Delete
				err := DeleteSpecialistService(state, service.ID)
				if err != nil {
					return none, err
				}
			}
		}
	}
	return specialist, nil
}
