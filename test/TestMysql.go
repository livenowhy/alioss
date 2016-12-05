package main

import "github.com/liuzhangpei/alioss/utils"

func main() {
	const dataSourceName  = "root:root123admin@tcp(192.168.1.6:3306)/registry?autocommit=true"
	utils.UpdateLogo("ac0b5a11-96aa-37a5-992f-e5a43fe5c55d", "sdsdsd", dataSourceName)
}

