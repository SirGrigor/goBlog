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
	fmt.Printf("Blog has been created: %v", crBlog)

	return crBlog
}

func main() {
	log.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
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
		&blogpb.ReadBlogRequest{BlogId: "608a9b465fc072108d2273dc"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}

	readBlogResponse, readBlogError := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogError != nil {
		fmt.Printf("Error happened while: %v", readBlogError)
	}

	log.Printf("Blog was Read: %v", readBlogResponse.GetBlog())
}
