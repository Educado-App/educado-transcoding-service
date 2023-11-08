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

func (s *GCPService) DownloadFile(fileName string) ([]byte, error) {
	var reader, err = s.client.Bucket(s.bucketName).Object(fileName).NewReader(s.context)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var content = make([]byte, reader.Attrs.Size)
	if _, err := reader.Read(content); err != nil {
		return nil, err
	}
	return content, nil
}

func (s *GCPService) UploadFile(fileName string, content []byte) error {
	var writer = s.client.Bucket(s.bucketName).Object(fileName).NewWriter(s.context)
	defer writer.Close()
	if _, err := writer.Write(content); err != nil {
		return err
	}
	return nil
}

func (s *GCPService) DeleteFile(fileName string) error {
	var err = s.client.Bucket(s.bucketName).Object(fileName).Delete(s.context)
	if err != nil {
		return err
	}
	return nil
}

func (s *GCPService) Reader(fileName string) (*storage.Reader, error) {
	var reader, err = s.client.Bucket(s.bucketName).Object(fileName).NewReader(s.context)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (s *GCPService) Attributes(fileName string) (*storage.ObjectAttrs, error) {
	var attrs, err = s.client.Bucket(s.bucketName).Object(fileName).Attrs(s.context)
	if err != nil {
		return nil, err
	}
	return attrs, nil
}
