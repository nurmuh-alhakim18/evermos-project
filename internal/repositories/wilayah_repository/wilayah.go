package wilayahrepository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type WilayahRepository struct{}

func (r *WilayahRepository) GetProvinces(ctx context.Context, search string, offset, limit int) ([]usermodel.Provinsi, error) {
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var provinces []usermodel.Provinsi
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	if search != "" {
		for _, p := range provinces {
			if p.Name == strings.ToUpper(search) {
				return []usermodel.Provinsi{p}, nil
			}
		}
	}

	if limit < len(provinces) {
		return provinces[offset : offset+limit], nil
	}

	return provinces, nil
}

func (r *WilayahRepository) GetCities(ctx context.Context, provinceID string) ([]usermodel.Kota, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cities []usermodel.Kota
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *WilayahRepository) GetProvinceDetail(ctx context.Context, provinceID string) (*usermodel.Provinsi, error) {
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var provinces []usermodel.Provinsi
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	var province usermodel.Provinsi
	for _, p := range provinces {
		if p.ID == provinceID {
			province = p
		}
	}

	return &province, nil
}

func (r *WilayahRepository) GetCityDetail(ctx context.Context, cityID string) (*usermodel.Kota, error) {
	provincies, err := r.getProvince()
	if err != nil {
		return nil, err
	}

	for _, p := range provincies {
		url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", p.ID)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		var cities []usermodel.Kota
		if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
			return nil, err
		}

		resp.Body.Close()

		for _, c := range cities {
			if c.ID == cityID {
				return &c, nil
			}
		}
	}

	return nil, errors.New("city not exists")
}

func (r *WilayahRepository) getProvince() ([]usermodel.Provinsi, error) {
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var provinces []usermodel.Provinsi
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	return provinces, nil
}
