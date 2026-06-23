package service

import (
	models "alto_server/common/models"
	"alto_server/common/pkg/e"
	. "alto_server/common/utils"

	"github.com/gin-gonic/gin"
)

func getServices() ([]models.OltServiceInfo, error) {
	var serviceList []models.OltServiceInfo
	s1 := models.OltServiceInfo{
		Name:        "HSI_VLAN222",
		Version:     "1",
		Status:      "active",
		Description: "HSI_VLAN222",
		ID:          "11111",
	}

	s2 := models.OltServiceInfo{
		Name:        "HSI_123",
		Version:     "1",
		Status:      "inactive",
		Description: "HSI_123",
		ID:          "2222",
	}

	s3 := models.OltServiceInfo{
		Name:        "IGMP_2022",
		Version:     "1",
		Status:      "active",
		Description: "IGMP_2022",
		ID:          "3333",
	}

	s4 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s5 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s6 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s7 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s8 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s9 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	s10 := models.OltServiceInfo{
		Name:        "Alden_2023",
		Version:     "1",
		Status:      "inactive",
		Description: "Alden_2023",
		ID:          "4444",
	}
	serviceList = append(serviceList, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10)

	return serviceList, nil

}
func getAllServicesHandler(c *gin.Context) {
	services, err := getServices()
	if err != nil {
		AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
	}

	RES(c, e.SUCCESS, gin.H{
		"service_list": services,
		"status":       e.SUCCESS,
		"message":      "Services list  get success.",
	})
}

func createServicesHandler(c *gin.Context) {
	RES(c, e.SUCCESS, gin.H{
		"status":  e.SUCCESS,
		"message": "Services list  get success.",
	})

}
