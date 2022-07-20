
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
	@ swag init --output ./docs/api

prepare-app:
	@ ENV=$(ENV) go run ./app/ui/cli/cmd.go

build-app-dependencies-in-docker:
	@ docker build --compress --force-rm --rm --tag $(ECR_URI)-dependencies stageDepName=$(ECR_URI)-dependencies -f ./build/app/docker/Dockerfile --target dependencies .

build-app-in-docker:
	@ docker build --compress --force-rm --rm --tag $(ECR_URI) -f ./build/app/docker/Dockerfile --target tweets-timeline .

run-app-in-docker:
	@ docker run -dti -p $(API_PORT):$(API_PORT) --env-file ./$(ENV).env --name tweets-timeline $(ECR_URI)

start:
	@  ENV=$(ENV) go run ./main.go