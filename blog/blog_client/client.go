package main

import (
	"context"
	"fmt"
	"goBlog/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func createBlog(err error, c blogpb.BlogServiceClient, blog *blogpb.Blog) *blogpb.CreateBlogResponse {
	createBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpercted Error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlog)

	return createBlog
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
	createblog := createBlog(err, c, blog)
	// ReadBlog
	blogID := createblog.GetBlog().GetId()

	_, err = c.ReadBlog(
		context.Background(),
		&blogpb.ReadBlogRequest{BlogId: "21414513"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogResponse, readBlogError := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogError != nil {
		fmt.Printf("Error happened while: %v", readBlogError)
	}

	log.Printf("Blog was Read: %v", readBlogResponse)
}
