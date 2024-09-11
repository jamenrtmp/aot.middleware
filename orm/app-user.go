package orm

import (
	"github.com/FourWD/middleware/model"
)

type AppUser struct {
	ID string `json:"id" query:"id" firestore:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	AppUserStatusID  string `json:"app_user_status_id" query:"app_user_status_id" firestore:"app_user_status_id" gorm:"type:varchar(2)"`
	AppDriverID      string `json:"app_driver_id" query:"app_driver_id" firestore:"app_driver_id" gorm:"type:varchar(36)"`
	DriverName       string `json:"driver_name" query:"driver_name" firestore:"driver_name" gorm:"type:varchar(100)"`
	Remark           string `json:"remark" query:"remark" firestore:"remark" gorm:"type:text"`
	DriverImagePath  string `json:"driver_image_path" query:"driver_image_path" firestore:"driver_image_path" gorm:"type:varchar(255)"`
	VehicleImage     string `json:"vehicle_image" query:"vehicle_image" firestore:"vehicle_image" gorm:"type:varchar(255)"`
	VehicleColorName string `json:"vehicle_color_name" query:"vehicle_color_name" firestore:"vehicle_color_name" gorm:"type:varchar(50)"`
	VehicleCode      string `json:"vehicle_code" query:"vehicle_code" firestore:"vehicle_code" gorm:"type:varchar(10)"`

	VehicleName     string  `json:"vehicle_name" query:"vehicle_name" firestore:"vehicle_name" gorm:"type:varchar(100)"`
	LicensePlate    string  `json:"license_plate" query:"license_plate" firestore:"license_plate" gorm:"type:varchar(100)"`
	OriginName      string  `json:"origin_name" query:"origin_name" firestore:"origin_name" gorm:"type:varchar(100)"`
	DestinationName string  `json:"destination_name" query:"destination_name" firestore:"destination_name" gorm:"type:varchar(100)"`
	CustomerID      string  `json:"customer_id" query:"customer_id" firestore:"customer_id" gorm:"type:varchar(100)"`
	VehiclePrice    float64 `json:"vehicle_price" query:"vehicle_price" firestore:"vehicle_price" gorm:"default:null; type:decimal(16,4);comment:'ราคาค่าบริการ'"`

	SlipID string `json:"slip_id" query:"slip_id" firestore:"slip_id" gorm:"type:varchar(36);"`
}
