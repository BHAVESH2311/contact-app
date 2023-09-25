package contactdetails

type ContactDetails struct {

    DetailID    int

    ContactType string

    DetailValue string

}

 


func NewContactDetail(detailID int, contactType, detailValue string) *ContactDetails {

    return &ContactDetails{

        DetailID:    detailID,

        ContactType: contactType,

        DetailValue: detailValue,

    }

}

 


type ContactDetailsRepository struct {

    contactDetails []*ContactDetails

}

 


func (cdr *ContactDetailsRepository) CreateContactDetail(contactType, detailValue string) *ContactDetails {

    detailID := len(cdr.contactDetails) + 1

    newContactDetail := NewContactDetail(detailID, contactType, detailValue)

    cdr.contactDetails = append(cdr.contactDetails, newContactDetail)

    return newContactDetail

}

 


func (cdr *ContactDetailsRepository) UpdateContactDetail(detailID int, newContactType, newDetailValue string) {

    for _, detail := range cdr.contactDetails {

        if detail.DetailID == detailID {

            detail.ContactType = newContactType

            detail.DetailValue = newDetailValue

            return

        }

    }

}

 


func (cdr *ContactDetailsRepository) DeleteContactDetail(detailID int) {

    for i, detail := range cdr.contactDetails {

        if detail.DetailID == detailID {

            cdr.contactDetails = append(cdr.contactDetails[:i], cdr.contactDetails[i+1:]...)

            return

        }

    }

}

 


func (cdr *ContactDetailsRepository) GetContactDetailByID(detailID int) *ContactDetails {

    for _, detail := range cdr.contactDetails {

        if detail.DetailID == detailID {

            return detail

        }

    }

    return nil 

}

 