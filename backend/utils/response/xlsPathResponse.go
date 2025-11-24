package response

import "latih.in-be/internal/model"

func XlsPathResponse(data model.XlsPath) model.XlsPathResponse {
	xlsPathRes := model.XlsPathResponse{
		Id:        data.Id,
		FilePath:  data.FilePath,
		CreatedAt: data.CreatedAt,
	}

	return xlsPathRes
}

func XlsPathsResponse(data []model.XlsPath) []model.XlsPathResponse {
	xlsPathsRes := []model.XlsPathResponse{}

	for _, xlsPath := range data {
		xlsPathRes := XlsPathResponse(xlsPath)
		xlsPathsRes = append(xlsPathsRes, xlsPathRes)
	}

	return xlsPathsRes
}