# ecr-image-lister

A quick tool to list all docker images & tags present in ECR.

## How?

```sh
$ go get -v -d
$ go build
$  ./ecr-image-lister -help
Usage of ./ecr-image-lister:
  -out string
        The file to output to. Leave empty for stdout
  -region string
        The AWS region to list images from
```

## Samples:

It is possible to write to stdout:

```sh
$ ./ecr-image-lister
# Docker images

## Last versions
- plouf: 0.1.0-159

## All versions
- plouf: [0.1.0-159 0.1.0-116 0.1.0-103 0.1.0-80]
```

Or directly in a file:

```sh
$ ./ecr-image-lister -out index.md
$ cat index.md
# Docker images

## Last versions
- plouf: 0.1.0-159

## All versions
- plouf: [0.1.0-159 0.1.0-116 0.1.0-103 0.1.0-80]
$
```