package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pedramaghasian/go-blog/blogpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Pedram",
		Title:    "My First Blog",
		Content:  "Content Of The First Blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("error while calling CreateBlog RPC: %v", err)
	}
	fmt.Printf("Blog has been created: %v", res)
}
