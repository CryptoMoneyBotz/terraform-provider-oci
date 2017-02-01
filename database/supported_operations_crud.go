package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type SupportedOperationDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListSupportedOperations
}

func (s *SupportedOperationDatasourceCrud) Get() (e error) {
	var list *baremetal.ListSupportedOperations
	if list, e = s.Client.ListSupportedOperations(); e != nil {
		return
	}

	s.Res = list
	return
}

func (s *SupportedOperationDatasourceCrud) SetData() {
	if s.Res != nil {
		resources := []map[string]interface{}{}
		s.D.SetId(time.Now().UTC().String())
		for _, v := range s.Res.SupportedOperations {
			res := map[string]interface{}{
				"id": v.ID,
			}
			resources = append(resources, res)
		}
		s.D.Set("supported_operations", resources)
	}
	return
}
