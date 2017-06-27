package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"net/http"
)


/**
Evolucionar el primer ejercicio del curso de forma tal que ahora se ejecute
en un webserver que escuche en el puerto 8080 y sea invocado al llamar al servicio
GET /calculate/:number?divs=1,2,3 El servicio debe devolver la respuesta en formato JSON
 */
func main() {
	r := gin.Default()
	r.GET("/calculate/:number", func(c *gin.Context) {
		var num int
		var err error
		var divsStr []string
		var divs []int
		numberParam := c.Param("number")
		divsParams := c.Query("divs")
		num, err = strconv.Atoi(numberParam);
		if (err != nil) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "number es obligatorio y debe ser numerico"})
			return
		}
		if (len(divsParams)==0) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "divs son obligatorios y numericos"})
			return
		}

		divsStr = strings.Split(divsParams, ",")
		divs = make([]int, len(divsStr))
		for i,v := range divsStr {
			divs[i], err = strconv.Atoi(v)
			if (err != nil) {
				c.JSON(http.StatusBadRequest, gin.H{"status": "divs deben ser numericos"})
				return
			}
		}

		var sum = Sum(num, divs)
		c.JSON(200, gin.H{
			"result": sum,
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func Sum(number int, divs []int) int {
	sum := 0
	for  n := 1;n <= number ;n++  {
		for _,d := range divs {
			if (n % d == 0) {
				sum += n
				break
			}
		}
	}
	return sum
}
