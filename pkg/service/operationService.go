package service

import (
	"errors"
	"lottery_back/pkg/model"
)

func WebSocketSender(prizeId int, operation string) error {
	result, err := model.GetResult(prizeId)
	if err != nil {
		return err
	}

	switch operation {
	case "init":
		result = model.GetEmptyResult()
	case "show id":
		result = result.GetPrizeMaskedResult()
	case "show prize":
		result = result.GetWinnerMaskedResult()
	case "show winner":
	case "lottery":
		err := result.Lottery()
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid operation")
	}

	WebsocketApp.sender <- result

	return nil
}
