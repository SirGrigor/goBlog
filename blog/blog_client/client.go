package main

import (
	"context"
	"fmt"
	"goBlog/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func createBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog) *blogpb.CreateBlogResponse {
	crBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpercted Error: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", crBlog)

	return crBlog
}

func main() {
	log.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Cannot connect: %v\n", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	blog := &blogpb.Blog{
		AuthorId: "Stephane",
		Title:    "My First Blog",
		Content:  "Content of the first blog"}

	log.Println("Creating the Blog")
	//Create Blog
	createBlogResponse := createBlog(c, blog)
	// ReadBlog
	blogID := createBlogResponse.GetBlog().GetId()

	_, err = c.ReadBlog(
		context.Background(),
		&blogpb.ReadBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}

	readBlogResponse, readBlogError := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogError != nil {
		fmt.Printf("Error happened while: %v\n", readBlogError)
	}

	log.Printf("Blog was Read: %v\n", readBlogResponse.GetBlog())

	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog(with additions)",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})

	if updateErr != nil {
		fmt.Printf("Error happened while update: %v\n", updateErr)
	}

	log.Printf("Blog was updated: %v\n", updateRes)

}
