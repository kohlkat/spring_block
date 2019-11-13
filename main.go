package main

import (
	"fmt"
	s "github.com/gaspardpeduzzi/spring_block/server"
)

func main(){
	display()

	s.LaunchServer()





}



func display(){
		fmt.Println(`
____  ____  ____  __  __ _   ___    ____  __     __    ___  __ _    __     __   ____  ____ 
/ ___)(  _ \(  _ \(  )(  ( \ / __)  (  _ \(  )   /  \  / __)(  / )  (  )   / _\ (  _ \/ ___)
\___ \ ) __/ )   / )( /    /( (_ \   ) _ (/ (_/\(  O )( (__  )  (   / (_/\/    \ ) _ (\___ \
(____/(__)  (__\_)(__)\_)__) \___/  (____/\____/ \__/  \___)(__\_)  \____/\_/\_/(____/(____/


`)

}