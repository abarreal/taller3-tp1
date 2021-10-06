IMAGE_TAG=api:release
IMAGE=us-east1-docker.pkg.dev/fiuba-taller3/main/$(IMAGE_TAG)

build-image:
	docker build -t app:latest .

test-build:
	docker-compose up -d --build

test-logs:
	docker-compose logs -f

test-up:
	docker-compose up -d

test-stop:
	docker-compose stop

test-down:
	docker-compose down

tag-image:
	docker tag app:latest $(IMAGE)

gcloud-login:
	gcloud auth login

gcloud-logout:
	gcloud auth revoke

build-release:
	gcloud builds submit --ignore-file=.gcloudignore --tag $(IMAGE)

undeploy:
	gcloud run services delete api

deploy:
	gcloud run deploy api \
		--image=$(IMAGE)\
		--port=8080\
		--service-account=webapi@fiuba-taller3.iam.gserviceaccount.com\
		--allow-unauthenticated\
		--cpu=1\
		--memory=1Gi\
		--min-instances=0\
		--max-instances=2\
		--region=us-east1\
		--platform=managed\
		--timeout=60s\
		--set-env-vars=GCP_PROJECT_ID=fiuba-taller3,GIN_MODE=release