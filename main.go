package main

func main() {
	//fmt.Println("hell buddy")

	server := NewAPIServer(":3000")

	server.Run()
}
