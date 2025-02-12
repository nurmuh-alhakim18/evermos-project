package external

import (
	"encoding/json"
	"fmt"
	"net/http"

	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type External struct{}

func (ext *External) GetProvince(provinceId string) (usermodel.Provinsi, error) {
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(url)
	if err != nil {
		return usermodel.Provinsi{}, err
	}

	defer resp.Body.Close()

	var provinces []usermodel.Provinsi
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return usermodel.Provinsi{}, err
	}

	var province usermodel.Provinsi
	for _, p := range provinces {
		if p.ID == provinceId {
			province = p
		}
	}

	return province, nil
}

func (ext *External) GetCity(provinceId, cityId string) (usermodel.Kota, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceId)
	resp, err := http.Get(url)
	if err != nil {
		return usermodel.Kota{}, err
	}

	defer resp.Body.Close()

	var cities []usermodel.Kota
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return usermodel.Kota{}, err
	}

	var city usermodel.Kota
	for _, c := range cities {
		if c.ID == cityId {
			city = c
		}
	}

	return city, nil
}
