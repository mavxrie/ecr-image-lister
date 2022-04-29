package main

import (
	"flag"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type Image struct {
	Name     string
	Versions Versions
	RawTags  []string
}

func fetchTags(ecrClient *ecr.ECR, repositoryName string) ([]string, error) {
	names := []string{}

	// log.Printf("Retrieving image tags for %s...", repositoryName)

	images, err := ecrClient.ListImages(&ecr.ListImagesInput{
		RepositoryName: &repositoryName,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, imageId := range images.ImageIds {
		names = append(names, *imageId.ImageTag)
	}

	// log.Printf("Retrieved tags: %v", names)

	return names, nil
}

func main() {
	regionPtr := flag.String("region", "", "The AWS region to list images from")
	flag.Parse()

	awsRegion := *regionPtr

	if awsRegion == "" {
		awsRegion = os.Getenv("AWS_DEFAULT_REGION")
		if awsRegion == "" {
			awsRegion = "us-east-1"
		}
	}

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.Fatal(err)
	}

	ecrClient := ecr.New(awsSession)

	// log.Printf("Retrieving repositories list...")
	repositories, err := ecrClient.DescribeRepositories(&ecr.DescribeRepositoriesInput{})
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(repositories.Repositories, func(i, j int) bool {
		return *repositories.Repositories[i].RepositoryName < *repositories.Repositories[j].RepositoryName
	})

	images := []Image{}

	for _, repository := range repositories.Repositories {
		tags, err := fetchTags(ecrClient, *repository.RepositoryName)
		if err != nil {
			log.Fatal(err)
		}

		versionTags := []*Version{}
		for _, tag := range tags {
			parsedTag, err := parseVersion(tag)
			if err == nil {
				versionTags = append(versionTags, &parsedTag)
			}
		}

		versionSort(versionTags)

		images = append(images, Image{
			Name:     *repository.RepositoryName,
			Versions: versionTags,
			RawTags:  tags,
		})
	}

	imageListToMarkdown(images)
}
