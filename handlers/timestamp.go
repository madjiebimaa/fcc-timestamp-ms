package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-timestamp-ms/models"
)

func UnixToUTCHandler(c *gin.Context) {
	u := c.Param("unix")
	if strings.Contains(u, "-") {
		s := strings.Split(u, "-")

		var date []int
		for _, v := range s {
			val, err := strconv.Atoi(v)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "invalid url param value",
				})
				return
			}

			date = append(date, val)
		}

		mount := time.Month(date[1])
		t := time.Date(date[0], mount, date[2], 0, 0, 0, 0, time.UTC)
		u = strconv.Itoa(int(t.Unix()))
	}

	unix, err := strconv.ParseInt(u, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid url param value",
		})
		return
	}

	utc := time.Unix(unix, 0).UTC()

	res := models.TimeStamp{
		Unix: u,
		UTC:  utc.Format(time.RFC822),
	}

	c.JSON(http.StatusOK, res)
}
