 package main

import (
	"net/http"
	"strconv"
     
	"github.com/gin-gonic/gin"

	"bahire-hasab-engine/internal/bahirehasab"
)

func main() {

	router := gin.Default()

	router.GET("/bahire/:year/metqi", func(c *gin.Context) {

		yearParam := c.Param("year")

		year, err := strconv.Atoi(yearParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid year",
			})
			return
		}

		result := gin.H{
			"year":   year,
			"wenber": bahirehasab.Wenber(year),
			"abekte": bahirehasab.Abekte(year),
			"metqi":  bahirehasab.Metqi(year),
		}

		c.JSON(http.StatusOK, result)
	})
	router.GET("/bahire/:year/nenewe", func(c *gin.Context) {
    yearStr := c.Param("year")
    year, err := strconv.Atoi(yearStr)

    if err != nil {
        c.JSON(400, gin.H{
            "error": "invalid year",
        })
        return
    }

    nenewe := bahirehasab.Nenewe(year)

    c.JSON(200, gin.H{
        "year": year,
        "nenewe": nenewe,
    })
})
router.GET("/bahire/:year/fasika", func(c *gin.Context) {
    yearStr := c.Param("year")
    year, err := strconv.Atoi(yearStr)

    if err != nil {
        c.JSON(400, gin.H{
            "error": "invalid year",
        })
        return
    }

    fasika := bahirehasab.Fasika(year)

    c.JSON(200, gin.H{
        "year": year,
        "fasika": fasika,
    })
})

router.GET("/bahire/:year/all", func(c *gin.Context) {
	yearStr := c.Param("year")
	year, err := strconv.Atoi(yearStr)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid year",
		})
		return
	}

	result := bahirehasab.CalculateBahireHasab(year)

	c.JSON(200, result)
})
	router.Run(":8080")
}
