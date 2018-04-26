package nhtsa

type (
	BaseResponse struct {
		Count          int64  `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"searchCriteria"`
	}

	MakessResponse struct {
		BaseResponse
		Results []*Make `json:"results"`
	}

	ModelsResponse struct {
		BaseResponse
		Results []*Model `json:"results"`
	}

	VINResponse struct {
		BaseResponse
		Results []*VINDetails `json:"results"`
	}

	Model struct {
		MakeID    int64  `json:"make_id"`
		MakeName  string `json:"make_name"`
		ModelID   int64  `json:"model_id"`
		ModelName string `json:"model_name"`
	}

	Make struct {
		MakeID   int64  `json:"make_id"`
		MakeName string `json:"make_name"`
	}

	VINDetails struct {
		ABS                      string `json:"ABS"`
		AEB                      string `json:"AEB"`
		AdaptiveCruiseControl    string `json:"AdaptiveCruiseControl"`
		AdaptiveHeadlights       string `json:"AdaptiveHeadlights"`
		AdditionalErrorText      string `json:"AdditionalErrorText"`
		AirBagLocCurtain         string `json:"AirBagLocCurtain"`
		AirBagLocFront           string `json:"AirBagLocFront"`
		AirBagLocKnee            string `json:"AirBagLocKnee"`
		AirBagLocSeatCushion     string `json:"AirBagLocSeatCushion"`
		AirBagLocSide            string `json:"AirBagLocSide"`
		Artemis                  string `json:"Artemis"`
		AxleConfiguration        string `json:"AxleConfiguration"`
		Axles                    string `json:"Axles"`
		BasePrice                string `json:"BasePrice"`
		BatteryA                 string `json:"BatteryA"`
		BatteryATo               string `json:"BatteryA_to"`
		BatteryCells             string `json:"BatteryCells"`
		BatteryInfo              string `json:"BatteryInfo"`
		BatteryKWh               string `json:"BatteryKWh"`
		BatteryKWhTo             string `json:"BatteryKWh_to"`
		BatteryModules           string `json:"BatteryModules"`
		BatteryPacks             string `json:"BatteryPacks"`
		BatteryType              string `json:"BatteryType"`
		BatteryV                 string `json:"BatteryV"`
		BatteryVTo               string `json:"BatteryV_to"`
		BedLengthIN              string `json:"BedLengthIN"`
		BedType                  string `json:"BedType"`
		BlindSpotMon             string `json:"BlindSpotMon"`
		BodyCabType              string `json:"BodyCabType"`
		BodyClass                string `json:"BodyClass"`
		BrakeSystemDesc          string `json:"BrakeSystemDesc"`
		BrakeSystemType          string `json:"BrakeSystemType"`
		BusFloorConfigType       string `json:"BusFloorConfigType"`
		BusLength                string `json:"BusLength"`
		BusType                  string `json:"BusType"`
		CAFEBodyType             string `json:"CAFEBodyType"`
		CAFEMake                 string `json:"CAFEMake"`
		CAFEModel                string `json:"CAFEModel"`
		CashForClunkers          string `json:"CashForClunkers"`
		ChargerLevel             string `json:"ChargerLevel"`
		ChargerPowerKW           string `json:"ChargerPowerKW"`
		CoolingType              string `json:"CoolingType"`
		Country                  string `json:"Country"`
		CurbWeightLB             string `json:"CurbWeightLB"`
		CustomMotorcycleType     string `json:"CustomMotorcycleType"`
		DestinationMarket        string `json:"DestinationMarket"`
		DisplacementCC           string `json:"DisplacementCC"`
		DisplacementCI           string `json:"DisplacementCI"`
		DisplacementL            string `json:"DisplacementL"`
		Doors                    string `json:"Doors"`
		DriveType                string `json:"DriveType"`
		DriverAssist             string `json:"DriverAssist"`
		ESC                      string `json:"ESC"`
		EVDriveUnit              string `json:"EVDriveUnit"`
		ElectrificationLevel     string `json:"ElectrificationLevel"`
		EngineConfiguration      string `json:"EngineConfiguration"`
		EngineCycles             string `json:"EngineCycles"`
		EngineCylinders          string `json:"EngineCylinders"`
		EngineHP                 string `json:"EngineHP"`
		EngineHPTo               string `json:"EngineHP_to"`
		EngineKW                 string `json:"EngineKW"`
		EngineManufacturer       string `json:"EngineManufacturer"`
		EngineModel              string `json:"EngineModel"`
		EntertainmentSystem      string `json:"EntertainmentSystem"`
		EquipmentType            string `json:"EquipmentType"`
		ErrorCode                string `json:"ErrorCode"`
		ForwardCollisionWarning  string `json:"ForwardCollisionWarning"`
		FuelInjectionType        string `json:"FuelInjectionType"`
		FuelTypePrimary          string `json:"FuelTypePrimary"`
		FuelTypeSecondary        string `json:"FuelTypeSecondary"`
		GVWR                     string `json:"GVWR"`
		LaneDepartureWarning     string `json:"LaneDepartureWarning"`
		LaneKeepSystem           string `json:"LaneKeepSystem"`
		Make                     string `json:"Make"`
		Manufacturer             string `json:"Manufacturer"`
		ManufacturerID           string `json:"ManufacturerId"`
		ManufacturerType         string `json:"ManufacturerType"`
		Model                    string `json:"Model"`
		ModelYear                string `json:"ModelYear"`
		MotorcycleChassisType    string `json:"MotorcycleChassisType"`
		MotorcycleSuspensionType string `json:"MotorcycleSuspensionType"`
		NCAPBodyType             string `json:"NCAPBodyType"`
		NCAPMake                 string `json:"NCAPMake"`
		NCAPModel                string `json:"NCAPModel"`
		NCICCode                 string `json:"NCICCode"`
		NCSABodyType             string `json:"NCSABodyType"`
		NCSAMake                 string `json:"NCSAMake"`
		NCSAModel                string `json:"NCSAModel"`
		Note                     string `json:"Note"`
		OtherBusInfo             string `json:"OtherBusInfo"`
		OtherEngineInfo          string `json:"OtherEngineInfo"`
		OtherMotorcycleInfo      string `json:"OtherMotorcycleInfo"`
		OtherRestraintSystemInfo string `json:"OtherRestraintSystemInfo"`
		OtherTrailerInfo         string `json:"OtherTrailerInfo"`
		ParkAssist               string `json:"ParkAssist"`
		PlantCity                string `json:"PlantCity"`
		PlantCompanyName         string `json:"PlantCompanyName"`
		PlantCountry             string `json:"PlantCountry"`
		PlantState               string `json:"PlantState"`
		PossibleValues           string `json:"PossibleValues"`
		Pretensioner             string `json:"Pretensioner"`
		RearVisibilityCamera     string `json:"RearVisibilityCamera"`
		SeatBeltsAll             string `json:"SeatBeltsAll"`
		SeatRows                 string `json:"SeatRows"`
		Seats                    string `json:"Seats"`
		Series                   string `json:"Series"`
		Series2                  string `json:"Series2"`
		SteeringLocation         string `json:"SteeringLocation"`
		SuggestedVIN             string `json:"SuggestedVIN"`
		TPMS                     string `json:"TPMS"`
		TopSpeedMPH              string `json:"TopSpeedMPH"`
		TrackWidth               string `json:"TrackWidth"`
		TractionControl          string `json:"TractionControl"`
		TrailerBodyType          string `json:"TrailerBodyType"`
		TrailerLength            string `json:"TrailerLength"`
		TrailerType              string `json:"TrailerType"`
		TransmissionSpeeds       string `json:"TransmissionSpeeds"`
		TransmissionStyle        string `json:"TransmissionStyle"`
		Trim                     string `json:"Trim"`
		Trim2                    string `json:"Trim2"`
		Turbo                    string `json:"Turbo"`
		VIN                      string `json:"VIN"`
		ValveTrainDesign         string `json:"ValveTrainDesign"`
		VehicleType              string `json:"VehicleType"`
		WheelBaseLong            string `json:"WheelBaseLong"`
		WheelBaseShort           string `json:"WheelBaseShort"`
		WheelBaseType            string `json:"WheelBaseType"`
		WheelSizeFront           string `json:"WheelSizeFront"`
		WheelSizeRear            string `json:"WheelSizeRear"`
		Wheels                   string `json:"Wheels"`
		Windows                  string `json:"Windows"`
	}
)
