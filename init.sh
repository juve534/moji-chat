#!/bin/bash -ue

aws cloudformation deploy \
--template-file infrastructure/vpc.yaml \
--stack-name moji-chat-vpc \
--s3-bucket moji-chat-stack \
--parameter-overrides $(cat parameters/.env) \
--region=ap-northeast-1