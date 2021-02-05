SHELL := /bin/sh

.PHONY: go
go:
	docker build -t nfinstana/s3tester:latest .
	docker push nfinstana/s3tester:latest

.PHONY: java
java:
	mvn compile com.google.cloud.tools:jib-maven-plugin:2.7.1:dockerBuild
	docker push nfinstana/s3tester:latest

.PHONY: job
job:
	kubectl delete -f job.yaml
	kubectl apply -f job.yaml

.PHONY: log
log:
	kubectl logs -n namespace-core job/s3tester
