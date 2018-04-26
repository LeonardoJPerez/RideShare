package nhtsa

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/services"

	"github.com/juju/errors"
	"github.com/meshhq/gohttp"
)

var (
	baseURL                 = "https://vpic.nhtsa.dot.gov/api/"
	jsonFormat              = "format=json"
	modelsEndpointFormat    = "/vehicles/GetModelsForMakeIdYear/makeId/%s/vehicletype/moto?%s"
	makeEndpointFormat      = "/vehicles/GetMakesForVehicleType/moto?%s"
	decodeVINEndpointFormat = "/vehicles/decodevinvaluesextended/%s?%s&modelyear=%d"
)

// Service :
type Service struct {
	services.BaseService
}

func NewService() *Service {
	s := new(Service)
	s.Client = gohttp.NewClient(baseURL, http.Header{})

	return s
}

func (s *Service) GetMakes() ([]*Make, error) {
	request := &gohttp.Request{
		Method: gohttp.GET,
		URL:    fmt.Sprintf(makeEndpointFormat, jsonFormat),
	}

	response, err := s.Execute(request)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	data := &MakessResponse{}
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	return data.Results, nil
}

func (s *Service) GetModels(makeID string) ([]*Model, error) {
	request := &gohttp.Request{
		Method: gohttp.GET,
		URL:    fmt.Sprintf(modelsEndpointFormat, makeID, jsonFormat),
	}

	response, err := s.Execute(request)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	data := &ModelsResponse{}
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	return data.Results, nil
}

func (s *Service) DecodeVIN(vin string, year int) (*VINDetails, error) {
	request := &gohttp.Request{
		Method: gohttp.GET,
		URL:    fmt.Sprintf(decodeVINEndpointFormat, vin, year, jsonFormat),
	}

	response, err := s.Execute(request)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	data := &VINResponse{}
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
		return nil, errors.Trace(err)
	}

	return data.Results[0], nil
}
