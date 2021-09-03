package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func (c echo.Context) error {
    log.Println(">>> start action >>>")
    if err := next(c); err != nil {
      c.Error(err)
    }
    log.Println("<<< finished action <<<")
    return nil
  }
}
