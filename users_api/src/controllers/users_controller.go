package controllers

type userServer struct {
	//pb.UnimplementedUsersServer
}

func (us userServer) GetUser() {

}

func (us userServer) CreateUser() {

}

func (us userServer) ModifyUser() {

}

func (us userServer) EliminateUser() {

}

func getNewUserServer() userServer {
	return userServer{}
}
