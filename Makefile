SHELL := /bin/sh

.PHONY: all
all:
	docker build -t nfinstana/s3tester:latest .
	docker push nfinstana/s3tester:latest
