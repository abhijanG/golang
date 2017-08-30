package main

import "fmt"

type userprofile interface {
	getname() string
	getaddress() string
	getcontact() map[string]string
}

func getuserinfo(u userprofile) {

	fmt.Println("Name : ", u.getname())
	fmt.Println("Address : ", u.getaddress())
	fmt.Println("Conatct Details ", u.getcontact())
	fmt.Println("Email Extraction ", u.getcontact()["email"])
}
func interfacedemo() {

	s := profile{
		name:       "Sherlock Holmes",
		address:    "221B Baker Street",
		email:      "Sherlock@holmes.com",
		contactnum: "18811904",
	}

	getuserinfo(s)

}
