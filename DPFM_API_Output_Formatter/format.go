package dpfm_api_output_formatter

import (
	"data-platform-api-country-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToCountry(rows *sql.Rows) (*[]Country, error) {
	defer rows.Close()
	country := make([]Country, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Country{}

		err := rows.Scan(
			&pm.Country,
			&pm.GlobalRegion,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &country, nil
		}

		data := pm
		country = append(country, Country{
			Country:             data.Country,
			GlobalRegion:        data.GlobalRegion,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &country, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	countryText := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.Country,
			&pm.Language,
			&pm.CountryName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &countryText, err
		}

		data := pm
		countryText = append(countryText, Text{
			Country:             data.Country,
			Language:            data.Language,
			CountryName:         data.CountryName,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &countryText, nil
}
