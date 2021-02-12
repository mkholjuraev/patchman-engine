package controllers

import (
	"app/base/database"
	"app/base/models"
	"app/base/utils"
	"app/manager/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type AdvisorySystemsIdsResponse struct {
	Data  []SystemItem `json:"data"`
	Links Links        `json:"links"`
	Meta  ListMeta     `json:"meta"`
}

// nolint: lll
// @Summary Show me IDs of systems applicable to which current advisory is applicable
// @Description Show me the list of IDs related to particular advisory systems
// @ID listAdvisorySystems
// @Security RhIdentity
// @Accept json
// @Produce json
// @Param    advisory_id    path    string  true    "Advisory ID"
// @Param    limit          query   int     false   "Limit for paging, set -1 to return all"
// @Param    offset         query   int     false   "Offset for paging"
// @Param    sort    query   string  false   "Sort field" Enums(id,display_name,last_evaluation,last_upload,rhsa_count,rhba_count,rhea_count,stale)
// @Param    search         query   string  false   "Find matching text"
// @Param    filter[id]              query   string  false "Filter"
// @Param    filter[display_name]    query   string  false "Filter"
// @Param    filter[last_evaluation] query   string  false "Filter"
// @Param    filter[last_upload]     query   string  false "Filter"
// @Param    filter[rhsa_count]      query   string  false "Filter"
// @Param    filter[rhba_count]      query   string  false "Filter"
// @Param    filter[rhea_count]      query   string  false "Filter"
// @Param    filter[stale]           query   string    false "Filter"
// @Param    tags                    query   []string  false "Tag filter"
// @Param    filter[system_profile][sap_system] query  string  false "Filter only SAP systems"
// @Param    filter[system_profile][sap_sids][in] query []string  false "Filter systems by their SAP SIDs"
// @Success 200 {object} AdvisorySystemsIdsResponse
// @Router /api/patch/v1/advisories/{advisory_id}/ids [get]

func AdvisorySystemsIdsHandler(c *gin.Context){
	var dbItems = buildHandler(c)

	data := buildAdvisorySystemsIdsData(dbItems)
	var resp = AdvisorySystemsIdsResponse{
		Data:  data,
		Links: *links,
		Meta:  *meta,
	}
	c.JSON(http.StatusOK, &resp)
}

// func buildQuery(account int, advisoryName string) *gorm.DB {
// 	query := database.SystemAdvisories(database.Db, account).
// 		Select(SystemsSelect).
// 		Joins("JOIN advisory_metadata am ON am.id = sa.advisory_id").
// 		Joins("JOIN inventory.hosts ih ON ih.id = sp.inventory_id").
// 		Where("am.name = ?", advisoryName).
// 		Where("sp.stale = false")

// 	return query
// }

func buildAdvisorySystemsIdsData(dbItems []SystemDBLookup) []SystemItem {
	data := make([]SystemItem, len(dbItems))
	for i, model := range dbItems {
		item := SystemItem{
			ID:         model.ID,
			Type:       "system",
		}
		data[i] = item
	}
	return data
}
