package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func RecoveryResponse(data entities.Recovery) *Responses {
	recovery := Recovery{
		ID:     data.ID,
		UserId: int(data.UserID),
		OTP:    data.OTP,
	}

	return &Responses{
		Status: true,
		Data:   recovery,
		Error:  nil,
	}
}

func RecoveriesResponse(data []entities.Recovery) *Responses {
	recoveries := []Recovery{}

	for _, recovery := range data {
		recoveries = append(recoveries, Recovery{
			ID:     recovery.ID,
			UserId: int(recovery.UserID),
			OTP:    recovery.OTP,
		})
	}
	return &Responses{
		Status: true,
		Data:   recoveries,
		Error:  nil,
	}
}
