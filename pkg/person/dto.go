package person

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
)

type ServiceI interface {
	Create(ctx context.Context, person DTO) error
	SendSurvey(ctx context.Context, people []DTO)
}

type DTO struct {
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
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
