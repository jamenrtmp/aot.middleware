package utils

import (
	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GetPriceOld(systemName string, distance float64) orm.PriceByDistance {
	var getprice orm.PriceByDistance

	var sql = `SELECT * FROM price_by_distances
	    WHERE fleet_client_id = ? AND distance >= ? ORDER BY distance LIMIT 1`
	common.Database.Raw(sql, systemName, int(distance)).Debug().Scan(&getprice)
	return getprice
}

type GetPriceModel struct {
	VehicleModelID string  `json:"vehicle_model_id" query:"vehicle_model_id"`
	Price          float64 `json:"price" query:"price" `
}

func GetPrice(fleetName string, distance float64) []GetPriceModel {
	var byDistance orm.PriceByDistance
	sql := `SELECT * FROM price_by_distances
	    WHERE fleet_client_id = ? AND distance >= ? ORDER BY distance LIMIT 1`
	// /log.Println(sql)
	common.Database.Raw(sql, fleetName, distance).Debug().Scan(&byDistance)

	var prices []GetPriceModel
	prices = append(prices, GetPriceModel{VehicleModelID: "01", Price: byDistance.ByVehicleModelID01})
	prices = append(prices, GetPriceModel{VehicleModelID: "02", Price: byDistance.ByVehicleModelID02})
	prices = append(prices, GetPriceModel{VehicleModelID: "03", Price: byDistance.ByVehicleModelID03})
	prices = append(prices, GetPriceModel{VehicleModelID: "04", Price: byDistance.ByVehicleModelID04})
	prices = append(prices, GetPriceModel{VehicleModelID: "05", Price: byDistance.ByVehicleModelID05})
	prices = append(prices, GetPriceModel{VehicleModelID: "06", Price: byDistance.ByVehicleModelID06})
	prices = append(prices, GetPriceModel{VehicleModelID: "07", Price: byDistance.ByVehicleModelID07})
	prices = append(prices, GetPriceModel{VehicleModelID: "08", Price: byDistance.ByVehicleModelID08})
	prices = append(prices, GetPriceModel{VehicleModelID: "09", Price: byDistance.ByVehicleModelID09})
	prices = append(prices, GetPriceModel{VehicleModelID: "10", Price: byDistance.ByVehicleModelID10})

	return prices
}

// func getSlipVehicle() {
// 	var getVehicle orm.PriceByDistance

// 	var sql = `SELECT * FROM price_by_distances
// 	    WHERE fleet_client_id = ? AND distance >= ? ORDER BY distance LIMIT 1`
// 	common.Database.Raw(sql, fleetClientID, distance).Scan(&getprice)
// 	common.Print(sql, "error")
// 	return getprice
// }
