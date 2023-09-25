package main

import (

    "fmt"

    "contactapp/user"

    "contactapp/contacts"

   "contactapp/contactdetails"

)

 

func main() {

    

    userRepo := &user.UserRepository{}

 


    adminUser := userRepo.CreateUser("admin", "admin", "Admin")

    staffUser := userRepo.CreateUser("user", "userpass", "Staff")

 


    contactRepo := &contact.ContactRepository{}

 

    contact1 := contactRepo.CreateContact("bhavesh", "Mishra", "bhavesh@example.com", "123-456-7890", "123 ghansoli St")

    contact2 := contactRepo.CreateContact("swastik", "chaudhary", "swastik@example.com", "987-654-3210", "456 ulwe St")

 


    contactDetailsRepo := &contactdetails.ContactDetailsRepository{}

 

	detail1 := contactdetails.NewContactDetail(1, "Email", "john@example.com")     
	detail2 := contactdetails.NewContactDetail(2, "Phone", "123-456-7890")

    detail3 := contactDetailsRepo.CreateContactDetail("Email", "alice@example.com")

    detail4 := contactDetailsRepo.CreateContactDetail("Phone", "987-654-3210")

    contact1.ContactDetails = append(contact1.ContactDetails, detail1, detail2)

    contact2.ContactDetails = append(contact2.ContactDetails, detail3, detail4)


    fmt.Println("Admin User:")

    fmt.Println(adminUser)

 

    fmt.Println("\nStaff User:")

    fmt.Println(staffUser)

 

    fmt.Println("\nContacts:")

    fmt.Println(contact1)

    fmt.Println(contact2)

 
    fmt.Println("\nContact 1 Details:")

    for _, detail := range contact1.ContactDetails {

        fmt.Println(detail)

    }

 

    fmt.Println("\nContact 2 Details:")

    for _, detail := range contact2.ContactDetails {

        fmt.Println(detail)

    }

}