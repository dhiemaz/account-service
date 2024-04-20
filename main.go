package main

import "github.com/dhiemaz/AccountService/cmd"

func main() {
	AccountSvc := cmd.NewCommand()
	AccountSvc.Run()
}
