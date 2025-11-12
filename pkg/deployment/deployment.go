package deployment

import (
	"github.com/francescofiorentinoTIM/opg-ewbi-api-main/api/federation/models"

)

type InstallDeployment struct {
	*models.InstallAppJSONBody
	FederationContextID string
}
