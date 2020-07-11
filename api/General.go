package api

func (a API) GetServerTime(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"time": time.Now().Format(time.RFC3339),
	})
}
