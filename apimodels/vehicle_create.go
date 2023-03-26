package apimodels

type ReqVehicleCreate struct {
	VehicleID           string `json:"vehicleID"`
	Make                string `json:"make"`
	Model               string `json:"model"`
	Year                string `json:"year"`
	BodyType            string `json:"bodyType"`
	InitialRegistration string `json:"initialRegistration"`
	RegistrationNumber  string `json:"registrationNumber"`
	Fuel                int    `json:"fuel"`
	Mileage             string `json:"mileage"`
	Drive               int    `json:"drive"`
	Color               string `json:"color"`
	Transmission        int    `json:"transmission"`
	Price               int    `json:"price"`
	ImageURL            string `json:"imageURL"`
}
