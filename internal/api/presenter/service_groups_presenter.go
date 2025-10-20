package presenter

import (
	"backend/internal/api/dtos"
	"backend/internal/infra"
)

func GetServiceGroups(specializations []infra.Specialization, serviceNames []infra.ServiceName) []dtos.ServiceGroup {
	response := make([]dtos.ServiceGroup, 0, len(specializations))

	for _, s := range specializations {
		count := 0
		for _, sn := range serviceNames {
			if sn.SpecializationID == s.ID {
				count++
			}
		}

		items := make([]dtos.ServiceGroupItem, 0, count)
		if count > 0 {
			for _, sn := range serviceNames {
				if sn.SpecializationID == s.ID {
					items = append(items, dtos.ServiceGroupItem{
						ID:   sn.ID.String(),
						Name: sn.Name,
					})
				}
			}
		}

		response = append(response, dtos.ServiceGroup{
			ID:    s.ID.String(),
			Name:  s.Name,
			Items: items,
		})
	}
	return response
}
