package main

import (
	"crm/utils"
	"fmt"
)

func main() {
	AccessToken := utils.GetTokenV2()
	fmt.Println(AccessToken)
	data := utils.DataQuery(AccessToken)

	fmt.Println(data)

}
