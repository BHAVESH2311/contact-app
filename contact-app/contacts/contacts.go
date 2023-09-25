package contact

import contactdetails "contactapp/contactdetails"


type Contact struct {

    ContactID int

    FirstName string

    LastName  string

    Email     string

    Phone     string

    Address   string

	ContactDetails []*contactdetails.ContactDetails

}

 


func NewContact(contactID int, firstName, lastName, email, phone, address string) *Contact {

    return &Contact{

        ContactID: contactID,

        FirstName: firstName,

        LastName:  lastName,

        Email:     email,

        Phone:     phone,

        Address:   address,

		ContactDetails: []*contactdetails.ContactDetails{},

    }

}

 


type ContactRepository struct {

    contacts []*Contact

}

 


func (cr *ContactRepository) CreateContact(firstName, lastName, email, phone, address string) *Contact {

    contactID := len(cr.contacts) + 1

    newContact := NewContact(contactID, firstName, lastName, email, phone, address)

    cr.contacts = append(cr.contacts, newContact)

    return newContact

}

 


func (cr *ContactRepository) UpdateContact(contactID int, newFirstName, newLastName, newEmail, newPhone, newAddress string) {

    for _, contact := range cr.contacts {

        if contact.ContactID == contactID {

            contact.FirstName = newFirstName

            contact.LastName = newLastName

            contact.Email = newEmail

            contact.Phone = newPhone

            contact.Address = newAddress

            return

        }

    }

}

 


func (cr *ContactRepository) DeleteContact(contactID int) {

    for i, contact := range cr.contacts {

        if contact.ContactID == contactID {

            cr.contacts = append(cr.contacts[:i], cr.contacts[i+1:]...)

            return

        }

    }

}

 


func (cr *ContactRepository) GetContactByID(contactID int) *Contact {

    for _, contact := range cr.contacts {

        if contact.ContactID == contactID {

            return contact

        }

    }

    return nil 

}