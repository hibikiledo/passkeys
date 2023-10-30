package response

import "github.com/teamhanko/passkey-server/persistence/models"

type GetConfigResponse struct {
	Cors     GetCorsResponse     `json:"cors"`
	Webauthn GetWebauthnResponse `json:"webauthn"`
}

func ToGetConfigResponse(config *models.Config) GetConfigResponse {
	return GetConfigResponse{
		Cors:     ToGetCorsResponse(&config.Cors),
		Webauthn: ToGetWebauthnResponse(&config.WebauthnConfig),
	}
}
