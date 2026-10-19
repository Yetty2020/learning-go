package main

//Structs can be nested to represent more complex entities


type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}
// The fields of a struct can be accessed with the dot operator
// myMessage := messageToSend{}
//myMessage.sender.name

func canSendMessage(mToSend messageToSend) bool {

	
	if mToSend.sender.name == "" || mToSend.sender.number == 0{
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0{
		return false
	}
	return true
}
