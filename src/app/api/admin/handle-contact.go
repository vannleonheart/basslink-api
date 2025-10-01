package admin

import "CRM/src/lib/basslink"

func (s *Service) getContacts() (*[]basslink.Contact, error) {
	var contacts []basslink.Contact

	if err := s.App.DB.Connection.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return &contacts, nil
}

func (s *Service) getContact(contactId string) (*basslink.Contact, error) {
	var contact basslink.Contact

	if err := s.App.DB.Connection.Preload("Documents").Preload("Accounts").Where("id = ?", contactId).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (s *Service) getContactAccounts(contactId string) (*[]basslink.ContactAccount, error) {
	var accounts []basslink.ContactAccount

	if err := s.App.DB.Connection.Where("contact_id = ?", contactId).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return &accounts, nil
}
