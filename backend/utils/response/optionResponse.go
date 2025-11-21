package response

import "latih.in-be/internal/model"

func OptionResponse(data model.Option) model.OptionResponse {
	optionRes := model.OptionResponse{
		Id:          data.Id,
		QuestionId:  data.QuestionId,
		OptionLabel: data.OptionLabel,
		OptionText:  data.OptionText,
		IsCorrect:   data.IsCorrect,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return optionRes
}

func OptionsResponse(data []model.Option) []model.OptionResponse {
	optionsRes := []model.OptionResponse{}

	for _, option := range data {
		optionRes := OptionResponse(option)
		optionsRes = append(optionsRes, optionRes)
	}

	return optionsRes
}