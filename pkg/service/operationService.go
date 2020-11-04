package service

import "lottery_back/pkg/model"

func WebSocketSender(prizeId int, operation string) error {
	result, err := model.GetResult(prizeId)
	if err != nil {
		return err
	}

	if operation == "lottery" {
		err := result.Lottery()
		if err != nil {
			return err
		}
	}

	switch operation {
	case "init":
		result = model.GetEmptyResult()
	case "show id":
		result = result.GetPrizeMaskedResult()
	case "show prize":
		result = result.GetWinnerMaskedResult()
	}

	WebsocketApp.sender <- result

	return nil
}
