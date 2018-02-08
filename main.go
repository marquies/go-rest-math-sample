// main.go

package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

type result struct {
	Result int `json:"result"`
}

type errorResponse struct {
	Reason string `json:"reason"`
	Hint   string `json:"hint"`
}

// The main function defines a single route, its handler
// and starts listening on port 8080 (default port for Beego)
func main() {
	/* This would match routes like the following:
	/service/sum/3/5
	/service/product/6/23
	...
	*/
	beego.Router("/service/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

// This is the controller that this application uses
type mainController struct {
	beego.Controller
}

// Get() handles all requests to the route defined above
func (c *mainController) Get() {
	//Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		opResult := add(num1, num2)
		r := result{Result: opResult}
		c.Data["json"] = &r
	case "product":
		opResult := multiply(num1, num2)
		r := result{Result: opResult}
		c.Data["json"] = &r
	default:
		//c.TplName = "invalid-route.html"
		r := errorResponse{Reason: "unknown operation",
			Hint: "Known operations are 'product' and 'sum'"}
		c.Data["json"] = &r
	}
	c.ServeJSON()
}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
