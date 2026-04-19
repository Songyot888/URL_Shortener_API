package controller

import (
	"my-shortener/controller/dto"
	"my-shortener/model"
	"my-shortener/service"
	"my-shortener/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func UrlController(route *gin.Engine) {
	routes := route.Group("/api/urls")
	{
		routes.GET("/", getAllUrls)
		routes.POST("/create", createUrl)
	}
	route.GET("/:shortCode", getNewUrlByShortCode)
}

func getAllUrls(c *gin.Context) {

	urlService := service.NewURLService()

	url, err := urlService.GetAllUrls()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	var urldto []dto.GetAllUrlsResponse
	copier.Copy(&urldto, &url)

	c.JSON(200, gin.H{"data": urldto})

}

func createUrl(c *gin.Context) {

	var createDTO dto.CreateUrlRequest
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var createModel model.URL
	createModel.ShortCode = utils.GenerateShortURL(6)

	copier.Copy(&createModel, &createDTO)

	urlService := service.NewURLService()
	createdUrl, err := urlService.Create(&createModel)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create URL"})
		return
	}

	var createDTOResponse dto.CreateUrlResponse
	copier.Copy(&createDTOResponse, &createdUrl)

	NewUrl := c.Request.Host + "/" + createDTOResponse.ShortCode

	c.JSON(201, gin.H{"message": "URL created successfully",
		"id":           createDTOResponse.Id,
		"original_url": createDTOResponse.OriginalURL,
		"short_url":    NewUrl})
}

func getNewUrlByShortCode(c *gin.Context) {
	shortCode := c.Param("shortCode")

	urlService := service.NewURLService()
	url, err := urlService.GetNewUrlByShortCode(shortCode)
	if err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}

	var urlRes model.URL
	copier.Copy(&urlRes, &url)

	c.Redirect(302, urlRes.OriginalURL)
}