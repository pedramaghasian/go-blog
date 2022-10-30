package main

import (
	"context"
	"fmt"
	"io"
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
	fmt.Printf("Blog has been created: %v\n", res)
	blogId := res.GetBlog().GetId()

	//read Blog
	fmt.Println("Reading the blog")
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		log.Fatalf("error happened while reading: %v", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog
	fmt.Println("updating the blog")
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content Of The First Blog (edited)",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		log.Fatalf("error happened while updating: %v", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

	// delete a blog
	fmt.Println("deleting the blog")
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})
	if deleteErr != nil {
		log.Fatalf("error happened while deleting: %v", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)

	// list Blogs
	fmt.Println("list the blogs")
	resStream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
