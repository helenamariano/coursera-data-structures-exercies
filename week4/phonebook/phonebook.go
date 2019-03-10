package main

// for 7 digit phone numbers max numbers is 10^7
const maxPhoneBookSize = 10000000

type PhoneBook struct {
	numToName []string
}

func (pb *PhoneBook) Add(number int, name string) {
	pb.numToName[number] = name
}

func (pb *PhoneBook) Del(number int) {
	// set the value to empty string to mark as deleted
	pb.numToName[number] = ""
}

func (pb *PhoneBook) Find(number int) string {
	if pb.numToName[number] == "" {
		return "not found"
	}
	return pb.numToName[number]
}

func NewPhoneBook() *PhoneBook {

	numToName := make([]string, maxPhoneBookSize)
	return &PhoneBook{numToName: numToName}
}
