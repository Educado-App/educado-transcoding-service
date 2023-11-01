package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
)

var Service = newGCPService()

type GCPService struct {
	context    context.Context
	client     *storage.Client
	bucketName string
}

func newGCPService() *GCPService {
	var ctx = context.Background()
	var client, err = storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	return &GCPService{
		context:    ctx,
		client:     client,
		bucketName: "educado-bucket",
	}
}

func (s *GCPService) ListFiles() ([]string, error) {
	var files []string
	it := s.client.Bucket(s.bucketName).Objects(s.context, nil)
	for {
		var attrs, err = it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, attrs.Name)
	}
	return files, nil
}
