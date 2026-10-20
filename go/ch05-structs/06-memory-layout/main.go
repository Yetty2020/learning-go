package main

type contact struct {
	userID       string //16bytes
	sendingLimit int32  //4bytes
	age          int32  //4bytes - that 32 represents that it is 32 bit machine
}

type perms struct {
	permissionLevel int  //8 byte
	canSend         bool // 1 byte
	canReceive      bool //1 byte
	canManage bool //1 byte
}
