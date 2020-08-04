.PHONY: build

build:
	sam build
deploy:
	sam build
	cd .aws-sam/build
	sam package --output-template-file packaged.yaml --s3-bucket sam-simple-chat
	sam deploy --template-file packaged.yaml --stack-name sam-simple-chat --capabilities CAPABILITY_IAM