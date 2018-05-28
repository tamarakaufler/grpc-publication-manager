package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	proto "github.com/tamarakaufler/grpc-publication-manager/author-service/proto"
	"golang.org/x/net/context"
)

var serviceHost string = "localhost:50051"

func init() {
	if os.Getenv("SERVICE_HOST") != "" {
		serviceHost = os.Getenv("SERVICE_HOST")
	}
}

func main() {
	var conn *grpc.ClientConn
	var err error

	fmt.Printf("Contacting author-service on %s ...\n", serviceHost)
	conn, err = grpc.Dial(serviceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := proto.NewAuthorServiceClient(conn)

	fmt.Println("1) Getting all authors ...\n")
	all, err := client.GetAllAuthors(context.Background(), &proto.GetAllRequest{})
	if err != nil {
		log.Printf("Could not list users: %v\n", err)
	} else {
		for _, v := range all.Authors {
			log.Println(v)
		}
	}

	firstName := "Lucien"
	lastName := "Kaufler"
	email := "lucien@gmail.com"
	password := "lucienpass"
	address := "101 Happy Street, Happy Town"
	country := "UK"

	fmt.Println("Creating a new author ...")
	_, err = client.CreateAuthor(context.Background(), &proto.Author{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Address:   address,
		Country:   country,
	})
	if err != nil {
		log.Printf("Could not create author: %v\n", err)
	} else {
		log.Printf("Created: %s\n", fmt.Sprintf("%s %s (%s)", firstName, lastName, email))
	}

	fmt.Println("2) Getting all authors ...\n")
	all, err = client.GetAllAuthors(context.Background(), &proto.GetAllRequest{})
	if err != nil {
		log.Printf("Could not list users: %v\n", err)
	} else {
		for _, v := range all.Authors {
			log.Println(v)
		}
	}

	fmt.Println("Authenticating author ...\n")
	ar, err := client.Authenticate(context.TODO(), &proto.Author{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Printf("Could not authenticate user: %s error: %v\n", email, err)
	} else {
		log.Printf("Author's access token is: %s \n", ar.Token)
	}

	os.Exit(0)
}
