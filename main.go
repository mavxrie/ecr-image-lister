package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

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
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
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

	for _, repository := range repositories.Repositories {
		tags, err := fetchTags(ecrClient, *repository.RepositoryName)
		if err != nil {
			log.Fatal(err)
		}

		allTagsValid := true

		versionTags := []*Version{}
		for _, tag := range tags {
			parsedTag, err := parseVersion(tag)
			if err != nil {
				allTagsValid = false
			} else {
				versionTags = append(versionTags, &parsedTag)
			}
		}

		versionSort(versionTags)

		if allTagsValid {
			fmt.Printf("%s: %v\n", *repository.RepositoryName, versionString(versionTags))
		} else {
			fmt.Printf("%s: %v\n", *repository.RepositoryName, tags)
		}
	}
}
