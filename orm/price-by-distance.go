package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PriceByDistance struct { // poi ของ รถแต่ละประเภท
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FleetClientID string `json:"fleet_client_id" query:"fleet_client_id" gorm:"type:varchar(10);"`

	Distance float64 `json:"distance" query:"distance" gorm:"type:decimal(16,4)"`

	ByVehicleModelID01 float64 `json:"by_vehicle_model_id_01" query:"by_vehicle_model_id_01" gorm:"default:null; column:by_vehicle_model_id_01; type:decimal(10,4); comment:'BMW' "`
	ByVehicleModelID02 float64 `json:"by_vehicle_model_id_02" query:"by_vehicle_model_id_02" gorm:"default:null; column:by_vehicle_model_id_02; type:decimal(10,4); comment:'BENZ' "`
	ByVehicleModelID03 float64 `json:"by_vehicle_model_id_03" query:"by_vehicle_model_id_03" gorm:"default:null; column:by_vehicle_model_id_03; type:decimal(10,4); comment:'Toyota Camry' "`
	ByVehicleModelID04 float64 `json:"by_vehicle_model_id_04" query:"by_vehicle_model_id_04" gorm:"default:null; column:by_vehicle_model_id_04; type:decimal(10,4); comment:'Toyota Commuter' "`
	ByVehicleModelID05 float64 `json:"by_vehicle_model_id_05" query:"by_vehicle_model_id_05" gorm:"default:null; column:by_vehicle_model_id_05; type:decimal(10,4); comment:'Isuzu Mu-x' "`
	ByVehicleModelID06 float64 `json:"by_vehicle_model_id_06" query:"by_vehicle_model_id_06" gorm:"default:null; column:by_vehicle_model_id_06; type:decimal(10,4); comment:'' "`
	ByVehicleModelID07 float64 `json:"by_vehicle_model_id_07" query:"by_vehicle_model_id_07" gorm:"default:null; column:by_vehicle_model_id_07; type:decimal(10,4); comment:'' "`
	ByVehicleModelID08 float64 `json:"by_vehicle_model_id_08" query:"by_vehicle_model_id_08" gorm:"default:null; column:by_vehicle_model_id_08; type:decimal(10,4); comment:'' "`
	ByVehicleModelID09 float64 `json:"by_vehicle_model_id_09" query:"by_vehicle_model_id_09" gorm:"default:null; column:by_vehicle_model_id_09; type:decimal(10,4); comment:'' "`
	ByVehicleModelID10 float64 `json:"by_vehicle_model_id_10" query:"by_vehicle_model_id_10" gorm:"default:null; column:by_vehicle_model_id_10; type:decimal(10,4); comment:'' "`
}
