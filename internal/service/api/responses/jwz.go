package responses

import "github.com/rarimo/verify-proxy/resources"

func NewCallbackData(callback_data string) *resources.CallbackDataResponse {
	return &resources.CallbackDataResponse{
		Data: resources.CallbackData{
			Key: resources.Key{
				Type: resources.CALLBACK_DATA,
			},
			Attributes: resources.CallbackDataAttributes{
				CallbackData: callback_data,
			},
		},
	}
}
