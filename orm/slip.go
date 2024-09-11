package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Slip struct {
	ID string `json:"id" query:"id" firestore:"ID" gorm:"type:varchar(36);unique;"`
	orm.GormModel
	RunningNo int `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`

	AppDriverStatusID string `json:"app_driver_status_id" query:"app_driver_status_id" firestore:"app_driver_status_id" gorm:"type:varchar(2);"`
	AppUserStatusID   string `json:"app_user_status_id" query:"app_user_status_id" firestore:"app_user_status_id" gorm:"type:varchar(2);"`

	SlipNo             string    `json:"slip_no" query:"slip_no" firestore:"slip_no" gorm:"type:varchar(36);index"`
	ReceiptNo          string    `json:"receipt_no" query:"receipt_no" firestore:"receipt_no" gorm:"type:varchar(50);index"`
	RrNo               string    `json:"rr_no" query:"rr_no" firestore:"rr_no" gorm:"type:varchar(50); index"`
	SlipDate           time.Time `json:"slip_date" query:"slip_date" firestore:"slip_date" gorm:"default:null;comment:'วันที่ลูกค้าจ่ายตัง' "`
	TravelDate         time.Time `json:"travel_date" query:"travel_date" firestore:"travel_date" gorm:"default:null;comment:'วันทีวิ่งจริง' "`
	ReconcileDate      time.Time `json:"reconcile_date" query:"reconcile_date" firestore:"reconcile_date" gorm:"default:null;comment:'วันที่คีย์ตั๋วเทียบ' "`
	SlipTypeID         string    `json:"slip_type_id" query:"slip_type_id" firestore:"slip_type_id" gorm:"type:varchar(2);"`
	SlipSubTypeID      string    `json:"slip_sub_type_id" query:"slip_sub_type_id" firestore:"slip_sub_type_id" gorm:"type:varchar(2);"`
	SlipVehicleModelID string    `json:"slip_vehicle_model_id" query:"slip_vehicle_model_id" firestore:"slip_vehicle_model_id" gorm:"type:varchar(2);comment:'ประเภทรถตามหน้าตั๋วที่ซื้อ'"`
	// SlipVehicleSubModelID string    `json:"slip_vehicle_submodel_id" query:"slip_vehicle_submodel_id" firestore:"slip_vehicle_submodel_id" gorm:"type:varchar(2);comment:'ประเภทรถตามหน้าตั๋วที่ซื้อ'"`
	IsPickup             bool      `json:"is_pickup" query:"is_pickup" firestore:"is_pickup" gorm:"default:0; type:bool;comment:'รับลูกค้าหรือยัง' "`
	PickupDate           time.Time `json:"pickup_date" query:"pickup_date" firestore:"pickup_date"`
	CounterID            string    `json:"counter_id" query:"counter_id" firestore:"counter_id" gorm:"type:varchar(36)"`
	OriginPoiID          string    `json:"origin_poi_id" query:"origin_poi_id" firestore:"origin_poi_id" gorm:"type:varchar(36);"`
	OriginName           string    `json:"origin_name" query:"origin_name" firestore:"origin_name" gorm:"type:varchar(150);"`
	ForceOriginName      string    `json:"force_origin_name" query:"force_origin_name" firestore:"force_origin_name" gorm:"type:varchar(150);comment:'กรณีที่ไม่มี POI' "`
	OriginLatitude       float64   `json:"origin_latitude" query:"origin_latitude" firestore:"origin_latitude" gorm:"type:decimal(10,6)"`
	OriginLongitude      float64   `json:"origin_longitude" query:"origin_longitude" firestore:"origin_longitude" gorm:"type:decimal(10,6)"`
	OriginRemark         string    `json:"origin_remark" query:"origin_remark" firestore:"origin_remark" gorm:"type:text"`
	DestinationPoiID     string    `json:"destination_poi_id" query:"destination_poi_id" firestore:"destination_poi_id" gorm:"type:varchar(36);"`
	DestinationName      string    `json:"destination_name" query:"destination_name" firestore:"destination_name" gorm:"type:varchar(150);"`
	ForceDestinationName string    `json:"force_destination_name" query:"force_destination_name" firestore:"force_destination_name" gorm:"type:varchar(150);comment:'กรณีที่ไม่มี POI' "`
	DestinationLatitude  float64   `json:"destination_latitude" query:"destination_latitude" firestore:"destination_latitude" gorm:"type:decimal(10,6)"`
	DestinationLongitude float64   `json:"destination_longitude" query:"destination_longitude" firestore:"destination_longitude" gorm:"type:decimal(10,6)"`
	IsDeliver            bool      `json:"is_deliver" query:"is_deliver" firestore:"is_deliver" gorm:"type:bool"`
	DeliverDate          time.Time `json:"deliver_date" query:"deliver_date" firestore:"deliver_date"`
	DestinationRemark    string    `json:"destination_remark" query:"destination_remark" firestore:"destination_remark" gorm:"type:text"`
	Distance             float64   `json:"distance" query:"distance" firestore:"distance" gorm:"type:decimal(16,4)"`
	Duration             int       `json:"duration" query:"duration" firestore:"duration" gorm:"type:int"`

	PriceRateID  string  `json:"price_rate_id" query:"price_rate_id" firestore:"price_rate_id" gorm:"type:varchar(36);"`
	PromotionID  string  `json:"promotion_id" query:"promotion_id" firestore:"promotion_id" gorm:"type:varchar(36);"`
	PromotionRef string  `json:"promotion_ref" query:"promotion_ref" firestore:"promotion_ref" gorm:"type:varchar(36);comment:'กรณี duoslip เก็บเลขที่ slip ถ้าkansaiเก็บ เลขของ kansai' "`
	Price        float64 `json:"price" query:"price" firestore:"price" gorm:"default:null; type:decimal(16,4);comment:'ราคาค่าบริการ' "`
	Discount     float64 `json:"discount" query:"discount" firestore:"discount" gorm:"default:null; type:decimal(16,4);comment:'ส่วนลด' "`
	Wht          float64 `json:"wht" query:"wht" firestore:"wht" gorm:"default:null; type:decimal(16,4);comment:'ภาษีหัก ที่จ่าย' "`
	Vat          float64 `json:"vat" query:"vat" firestore:"vat" gorm:"default:null; type:decimal(16,4);comment:'ภาษี' "`
	NetPrice     float64 `json:"net_price" query:"net_price" firestore:"net_price" gorm:"default:null; type:decimal(16,4);comment:'รวมราคาค่าบริการทั้งหมด' "`
	IsPaid       bool    `json:"is_paid" query:"is_paid" firestore:"is_paid" gorm:"default:0; type:bool;comment:'ลูกค้าจ่ายหรือยัง' "`
	PaidBy       string  `json:"paid_by" query:"paid_by" firestore:"paid_by" gorm:"type:varchar(36);"`

	IsConfirm   bool      `json:"is_confirm" query:"is_confirm" firestore:"is_confirm" gorm:"default:0; type:bool;comment:'ลูกค้าคอนเฟิร์มหรือยัง' "`
	ConfirmBy   string    `json:"confirm_by" query:"confirm_by" firestore:"confirm_by" gorm:"type:varchar(36);"`
	ConfirmDate time.Time `json:"confirm_date" query:"confirm_date" firestore:"confirm_date"`

	RefCode          string    `json:"ref_code" query:"ref_code" firestore:"ref_code" gorm:"type:varchar(36);comment:'refในหน้าstatusfleet'"`
	PaymentTypeID    string    `json:"payment_type_id" query:"payment_type_id" firestore:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentDate      time.Time `json:"payment_date" query:"payment_date" firestore:"payment_date" gorm:"default:0;comment:'วันเวลาที่จ่าย'"`
	CreditCardNo     string    `json:"credit_card_number" query:"credit_card_number" firestore:"credit_card_number" gorm:"default:null; type:varchar(36);"`
	CreditCardTypeID string    `json:"credit_card_type_id" query:"credit_card_type_id" firestore:"credit_card_type_id" gorm:"type:varchar(2);"`
	BankRefNo        string    `json:"bank_ref_number" query:"bank_ref_number" firestore:"bank_ref_number" gorm:"default:null; type:varchar(20);"`
	IsVoid           bool      `json:"is_void" query:"is_void" firestore:"is_void" gorm:"default:0; type:bool;comment:'ยกเลิก?' "`
	IsVoidTypeID     string    `json:"is_void_type_id" query:"is_void_type_id" firestore:"is_void_type_id" gorm:"type:varchar(36);"`
	VoidRemark       string    `json:"void_remark" query:"void_remark" firestore:"void_remark" gorm:"type:varchar(500);"`
	VoidBy           string    `json:"void_by" query:"void_by" firestore:"void_by" gorm:"type:varchar(36);"`
	VoidDate         time.Time `json:"void_date" query:"void_date" firestore:"void_date"`
	IsCancel         bool      `json:"is_cancel" query:"is_cancel" firestore:"is_cancel" gorm:"default:0;type:bool;comment:'ยกเลิก?' "`
	IsCancelTypeID   string    `json:"is_cancel_type_id" query:"is_cancel_type_id" firestore:"is_cancel_type_id" gorm:"type:varchar(36);"`
	CancelRemark     string    `json:"cancel_remark" query:"cancel_remark" firestore:"cancel_remark" gorm:"type:varchar(500);"`
	CancelBy         string    `json:"cancel_by" query:"cancel_by" firestore:"cancel_by" gorm:"type:varchar(36);"`
	CancelDate       time.Time `json:"cancel_date" query:"cancel_date" firestore:"cancel_date"`
	IsNewCustomer    bool      `json:"is_new_customer" query:"is_new_customer" firestore:"is_new_customer" gorm:"type:bool"`
	CustomerID       string    `json:"customer_id" query:"customer_id" firestore:"customer_id" gorm:"type:varchar(36);"`
	// Code             string    `json:"code" query:"code" firestore:"code" gorm:"type:varchar(20) ; dafault:null ; index"`
	CompanyName      string    `json:"company_name" query:"company_name" firestore:"company_name" gorm:"type:varchar(150)"`
	TaxNo            string    `json:"tax_no" query:"tax_no" firestore:"tax_no" gorm:"type:varchar(20)"`
	IsTax            bool      `json:"is_tax" query:"is_tax" firestore:"is_tax" gorm:"type:bool"`
	IsHQ             bool      `json:"is_hq" query:"is_hq" firestore:"is_hq" gorm:"type:bool"`
	Address          string    `json:"address" query:"address" firestore:"address" gorm:"type:text"`
	Postcode         string    `json:"postcode" query:"postcode" firestore:"postcode" gorm:"type:varchar(5)"`
	PhoneNo          string    `json:"phone_no" query:"phone_no" firestore:"phone_no" gorm:"type:varchar(10)"`
	FlightNo         string    `json:"flight_no" query:"flight_no" firestore:"flight_no" gorm:"type:varchar(20)"`
	FlightDate       time.Time `json:"flight_date" query:"flight_date" firestore:"flight_date"`
	RentalRateID     string    `json:"rental_rate_id" query:"rental_rate_id" firestore:"rental_rate_id" gorm:"type:varchar(36)"`
	RentalPrice      float64   `json:"rental_price" query:"rental_price" firestore:"rental_price" gorm:"type:decimal(16,4)"`
	RentalFuelRateID string    `json:"rental_fuel_rate_id" query:"rental_fuel_rate_id" firestore:"rental_fuel_rate_id" gorm:"type:varchar(20)"`
	RentalFuelLitre  float64   `json:"rental_fuel_litre" query:"rental_fuel_litre" firestore:"rental_fuel_litre" gorm:"type:decimal(16,4)"`
	FuelAverage      float64   `json:"fuel_average" query:"fuel_average" firestore:"fuel_average" gorm:"type:decimal(16,4);comment:'เก็บราคาของค่าน้ำมันเฉลี่ย '"`
	AssignVehicleID  string    `json:"assign_vehicle_id" query:"assign_vehicle_id" firestore:"assign_vehicle_id" gorm:"type:varchar(36);comment:'รถที่วิ่งงานจริง'"`
	AssignVehicleBy  string    `json:"assign_vehicle_by" query:"assign_vehicle_by" firestore:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID   string    `json:"assign_driver_id" query:"assign_driver_id" firestore:"assign_driver_id" gorm:"default:null; type:varchar(36); "`
	DriverRegisterID string    `json:"driver_register_id" query:"driver_register_id" firestore:"driver_register_id" gorm:"default:null; type:varchar(36); "`

	IsArrive        bool      `json:"is_arrive" query:"is_arrive" firestore:"is_arrive" gorm:"default:null;"`
	ArriveDate      time.Time `json:"arrive_date" query:"arrive_date" firestore:"arrive_date" gorm:"default:null;comment:'วันเวลาที่มาถึง'"`
	ArriveLatitude  float64   `json:"arrive_latitude" query:"arrive_latitude" firestore:"arrive_latitude" gorm:"type:decimal(10,6);comment:'จุดรับลูกค้าระยะพิกัดต้นทาง'"`
	ArriveLongitude float64   `json:"arrive_longitude" query:"arrive_longitude" firestore:"arrive_longitude" gorm:"type:decimal(10,6);comment:'จุดรับลูกค้าระยะพิกัดปลายทาง'"`
	IsComplete      bool      `json:"is_complete" query:"is_complete" firestore:"is_complete" gorm:"default:0; type:bool;comment:'เสร็จสมบูรณ์'"`
	CompleteDate    time.Time `json:"complete_date" query:"complete_date" firestore:"complete_date"`
	Remark          string    `json:"remark" query:"remark" firestore:"remark" gorm:"type:text"`
	BookingNo       string    `json:"booking_no" query:"booking_no" firestore:"booking_no" gorm:"type:varchar(50);index"`
	BookingBy       string    `json:"booking_by" query:"booking_by" firestore:"booking_by" gorm:"type:varchar(36)"`
	BookingDate     time.Time `json:"booking_date" query:"booking_date" firestore:"booking_date"`
	JobNo           string    `json:"job_no" query:"job_no" firestore:"job_no" gorm:"type:varchar(20);"`
	AcceptJobDate   time.Time `json:"accept_job_date" query:"accept_job_date" firestore:"accept_job_date"`
	AirlineName     string    `json:"airline_name" query:"airline_name" firestore:"airline_name" gorm:"type:varchar(100)"`

	//---FLEET--
	FleetClientID   string `json:"fleet_client_id" query:"fleet_client_id" firestore:"fleet_client_id" gorm:"type:varchar(10) ; default:null ; "`
	BookingHour     string `json:"booking_hour" query:"booking_hour" firestore:"booking_hour" gorm:"type:varchar(2);comment:'ระยะเวลาเช่าของระบบ Fleet '"`
	Passenger       int    `json:"passenger" query:"passenger" firestore:"passenger" gorm:"type:int"`
	NumberOfLuggage int    `json:"number_of_luggage" query:"number_of_luggage" firestore:"number_of_luggage" gorm:"type:int"`

	PassengerPrefixID  string `json:"passenger_prefix_id" query:"passenger_prefix_id" firestore:"passenger_prefix_id" gorm:"type:varchar(2);"`
	PassengerFirstname string `json:"passenger_firstname" query:"passenger_firstname" firestore:"passenger_firstname" gorm:"type:varchar(50);"`
	PassengerLastname  string `json:"passenger_lastname" query:"passenger_lastname" firestore:"passenger_lastname" gorm:"type:varchar(200);"`
	// PassengerCountryID     string `json:"passenger_country_id" query:"passenger_country_id" firestore:"passenger_country_id" gorm:"type:varchar(10);"`
	PassengerGenderID      string `json:"passenger_gender_id" query:"passenger_gender_id" firestore:"passenger_gender_id" gorm:"type:varchar(10);"`
	PassengerNationalityID string `json:"passenger_nationality_id" query:"passenger_nationality_id" firestore:"passenger_nationality_id" gorm:"type:varchar(10);"`
	PassengerEmail         string `json:"passenger_email" query:"passenger_email" firestore:"passenger_email" gorm:"type:varchar(255);"`
	PassengerPhone         string `json:"passenger_phone" query:"passenger_phone" firestore:"passenger_phone" gorm:"type:varchar(20);"`

	ContactPersonPrefixID  string `json:"contact_person_prefix_id" query:"contact_person_prefix_id" firestore:"contact_person_prefix_id" gorm:"type:varchar(10);"`
	ContactPersonFirstName string `json:"contact_person_firstname" query:"contact_person_name" firestore:"contact_person_name" gorm:"type:varchar(200);"`
	ContactPersonLastname  string `json:"contact_person_lastname" query:"contact_person_lastname" firestore:"contact_person_lastname" gorm:"type:varchar(200);"`
	ContactPersonPhone     string `json:" contact_person_phone " query:"contact_person_phone" firestore:"contact_person_phone" gorm:"type:varchar(20)"`

	SubPassengerFirstname1   string `json:"sub_passenger_firstname_1" query:"sub_passenger_firstname_1" firestore:"sub_passenger_firstname_1" gorm:"type:varchar(255);column:sub_passenger_name_1;"`
	SubPassengerLastname1    string `json:"sub_passenger_lastname_1" query:"sub_passenger_lastname_1" firestore:"sub_passenger_lastname_1" gorm:"type:varchar(255);column:sub_passenger_lastname_1;"`
	SubPassengerDestination1 string `json:"sub_passenger_destination_1" query:"sub_passenger_destination_1" firestore:"sub_passenger_destination_1" gorm:"type:varchar(255);column:sub_passenger_destination_1;"`

	SubPassengerFirstname2   string `json:"sub_passenger_firstname_2" query:"sub_passenger_firstname_2" firestore:"sub_passenger_firstname_2" gorm:"type:varchar(255);column:sub_passenger_firstname_2;"`
	SubPassengerLastname2    string `json:"sub_passenger_lastname_2" query:"sub_passenger_lastname_2" firestore:"sub_passenger_lastname_2" gorm:"type:varchar(255);column:sub_passenger_lastname_2;"`
	SubPassengerDestination2 string `json:"sub_passenger_destination_2" query:"sub_passenger_destination_2" firestore:"sub_passenger_destination_2" gorm:"type:varchar(255);column:sub_passenger_destination_2;"`

	SubPassengerFirstname3   string `json:"sub_passenger_firstname_3" query:"sub_passenger_firstname_3" firestore:"sub_passenger_firstname_3" gorm:"type:varchar(255);column:sub_passenger_name_3;"`
	SubPassengerLastname3    string `json:"sub_passenger_lastname_3" query:"sub_passenger_lastname_3" firestore:"sub_passenger_lastname_3" gorm:"type:varchar(255);column:sub_passenger_lastname_3;"`
	SubPassengerDestination3 string `json:"sub_passenger_destination_3" query:"sub_passenger_destination_3" firestore:"sub_passenger_destination_3" gorm:"type:varchar(255);column:sub_passenger_destination_3;"`

	DropOff1          string    `json:"drop_off_1" query:"drop_off_1" firestore:"drop_off_1" gorm:"varchar(255);column:drop_off_1"`
	DropOffLatitude1  float64   `json:"drop_off_latitude_1" query:"drop_off_latitude_1" firestore:"drop_off_latitude_1" gorm:"type:decimal(10,6);column:drop_off_latitude_1"`
	DropOffLongitude1 float64   `json:"drop_off_longitude_1" query:"drop_off_longitude_1" firestore:"drop_off_longitude_1" gorm:"type:decimal(10,6);column:drop_off_longitude_1"`
	IsDropOff1        bool      `json:"is_drop_off_1" query:"is_drop_off_1" firestore:"is_drop_off_1" gorm:"type:bool;column:is_drop_off_1"`
	IsCharge1         bool      `json:"is_charge_1" query:"is_charge_1" firestore:"is_charge_1" gorm:"default:0; type:bool;column:is_charge_1;"`
	DropOffDate1      time.Time `json:"drop_off_date_1" query:"drop_off_date_1" firestore:"drop_off_date_1" gorm:"column:drop_off_date_1"`
	DropOffRemark1    string    `json:"drop_off_remark_1" query:"drop_off_remark_1" firestore:"drop_off_remark_1" gorm:"type:text;column:drop_off_remark_1"`

	DropOff2          string    `json:"drop_off_2" query:"drop_off_2" firestore:"drop_off_2" gorm:"varchar(255);column:drop_off_2"`
	DropOffLatitude2  float64   `json:"drop_off_latitude_2" query:"drop_off_latitude_2" firestore:"drop_off_latitude_2" gorm:"type:decimal(10,6);column:drop_off_latitude_2"`
	DropOffLongitude2 float64   `json:"drop_off_longitude_2" query:"drop_off_longitude_2" firestore:"drop_off_longitude_2" gorm:"type:decimal(10,6);column:drop_off_longitude_2"`
	IsDropOff2        bool      `json:"is_drop_off_2" query:"is_drop_off_2" firestore:"is_drop_off_2" gorm:"type:bool;column:is_drop_off_2"`
	IsCharge2         bool      `json:"is_charge_2" query:"is_charge_2" firestore:"is_charge_2" gorm:"default:0; type:bool;column:is_charge_2;"`
	DropOffDate2      time.Time `json:"drop_off_date_2" query:"drop_off_date_2" firestore:"drop_off_date_2" gorm:"column:drop_off_date_2"`
	DropOffRemark2    string    `json:"drop_off_remark_2" query:"drop_off_remark_2" firestore:"drop_off_remark_2" gorm:"type:text;column:drop_off_remark_2"`

	DropOff3          string    `json:"drop_off_3" query:"drop_off_3" firestore:"drop_off_3" gorm:"varchar(255);column:drop_off_3"`
	DropOffLatitude3  float64   `json:"drop_off_latitude_3" query:"drop_off_latitude_3" firestore:"drop_off_latitude_3" gorm:"type:decimal(10,6);column:drop_off_latitude_3"`
	DropOffLongitude3 float64   `json:"drop_off_longitude_3" query:"drop_off_longitude_3" firestore:"drop_off_longitude_3" gorm:"type:decimal(10,6);column:drop_off_longitude_3"`
	IsDropOff3        bool      `json:"is_drop_off_3" query:"is_drop_off_3" firestore:"is_drop_off_3" gorm:"type:bool;column:is_drop_off_3"`
	IsCharge3         bool      `json:"is_charge_3" query:"is_charge_3" firestore:"is_charge_3" gorm:"default:0;type:bool;column:is_charge_3;"`
	DropOffDate3      time.Time `json:"drop_off_date_3" query:"drop_off_date_3" firestore:"drop_off_date_3" gorm:"column:drop_off_date_3"`
	DropOffRemark3    string    `json:"drop_off_remark_3" query:"drop_off_remark_3" firestore:"drop_off_remark_3" gorm:"type:text;column:drop_off_remark_3"`

	IsDuo     bool   `json:"is_duo" query:"is_duo" firestore:"is_duo" gorm:"default:0;type:bool;"`            //  ตั๋วไปกลับ
	DuoSlipID string `json:"duo_slip_id" query:"duo_slip_id" firestore:"duo_slip_id" gorm:"type:varchar(36)"` // ตั๋วขาแรก

	IsCarSeat    bool   `json:"is_car_seat" query:"is_car_seat" firestore:"is_car_seat" gorm:"default:0; type:bool;column:is_car_seat;"`
	RemarkDriver string `json:"remark_driver" query:"remark_driver" firestore:"remark_driver" gorm:"type:text;"`

	IsAssign bool `json:"is_assign" query:"is_assign" firestore:"is_assign" gorm:"default:0; type:bool;column:is_assign;"`

	IsOperationApprove   bool      `json:"is_operation_approve" query:"is_operation_approve" firestore:"is_operation_approve" gorm:"type:bool; default:false;"`
	OperationApproveBy   string    `json:"operation_approve_by" query:"operation_approve_by" firestore:"operation_approve_by" gorm:"type:varchar(36)"`
	OperationApproveDate time.Time `json:"operation_approve_date" query:"operation_approve_date" firestore:"operation_approve_date"`
}
