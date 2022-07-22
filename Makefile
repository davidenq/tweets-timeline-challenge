.ONESHELL:
SHELL=/bin/bash
-include $(PWD)/$(ENV).env
export


create-mocks:
	@ mockery --dir=./app --all --recursive=true --disable-version-string

unit-tests:
	@ ENV=$(ENV) go test `go list ./... | grep -v ./tests` -cover

integration-tests:
	@  ENV=$(ENV) go test -tags=integration -v ./tests/integration 

end2end-tests:
	@ echo "not implemented yet!"

start-dynamo:
	@ docker-compose -f  $(shell pwd)/deployments/docker-compose/docker-compose.yml  run --rm -p 8000:8000 -d dynamodb
	@ sudo chmod 777 -R $(shell pwd)/deployments/docker-compose/storage

swagger:
	@ sh ./scripts/replace.sh
	@ echo swag init --output ./docs/api

prepare-app:
	@ ENV=$(ENV) go run ./app/ui/cli/cmd.go

build-app-dependencies-in-docker:
	@ docker build --compress --force-rm --rm --tag $(ECR_URI)-dependencies stageDepName=$(ECR_URI)-dependencies -f ./build/app/docker/Dockerfile --target dependencies .

build-app-in-docker:
	@ docker build --compress --force-rm --rm --tag $(ECR_URI) -f ./build/app/docker/Dockerfile --target tweets-timeline .

build-web-in-docker:
	@ docker build --compress --force-rm --rm --tag $(ECR_URI)/web -f ./build/web/docker/Dockerfile  .

run-app-in-docker:
	@ docker run -dti -p $(API_PORT):$(API_PORT) --env-file ./$(ENV).env --name tweets-timeline $(ECR_URI)

start-web:
	@ ENV=$(ENV) npm run --prefix ./web dev
start-app:
	@  ENV=$(ENV) go run ./main.go

aws-ecr-login:
	@ aws ecr get-login-password --region $(AWS_DEFAULT_REGION) --profile $(AWS_PROFILE) | docker login --username AWS --password-stdin $(ECR_URI)

publish-docker-app-on-ecr:
	@ docker push $(ECR_URI)

prepare-ecs-fargate:
	@ aws iam --region $(AWS_DEFAULT_REGION) create-role --role-name ecsTaskExecutionRole --assume-role-policy-document ./deployments/ecs/task-execution-assume-role.json
	@ aws iam --region $(AWS_DEFAULT_REGION) attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
	@ ecs-cli configure --cluster tweets-timeline-cluster-dev --default-launch-type FARGATE --config-name tweets-timeline --region $(AWS_DEFAULT_REGION)
	@ ecs-cli configure profile --access-key $(AWS_ACCESS_KEY_ID) --secret-key $(AWS_SECRET_ACCESS_KEY) --profile-name dynamo-ecr-wr
	@ ecs-cli up --cluster-config tweets-timeline-cluster-dev --ecs-profile dynamo-ecr-wr
	@ aws ec2 describe-security-groups --filters Name=vpc-id,Values=VPC_ID --region $(AWS_DEFAULT_REGION) --profile dynamo-ecr-wr