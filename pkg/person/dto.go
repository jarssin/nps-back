package person

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
)

type ServiceI interface {
	SendSurvey(people []DTO)
}

type DTO struct {
	VisitorId string `json:"visitorId"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}

func ToDTOFromCSV(csvReader *csv.Reader) ([]DTO, error) {
	headers, err := csvReader.Read()
	if err != nil {
		return []DTO{}, err
	}

	for i := range headers {
		headers[i] = strings.TrimSpace(headers[i])
	}

	var people []DTO
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []DTO{}, err
		}

		rowMap := make(map[string]string)
		for i, value := range record {
			if i < len(headers) {
				rowMap[headers[i]] = strings.TrimSpace(value)
			}
		}

		rowJson, err := json.Marshal(rowMap)
		if err != nil {
			return []DTO{}, err
		}

		var p DTO
		if err := json.Unmarshal(rowJson, &p); err != nil {
			return []DTO{}, err
		}
		people = append(people, p)
	}

	return people, nil
}
