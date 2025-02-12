package externalinterface

import usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"

type ExternalInterface interface {
	GetProvince(provinceId string) (usermodel.Provinsi, error)
	GetCity(provinceId, cityId string) (usermodel.Kota, error)
}
