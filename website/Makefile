IMAGE_TAG=website:release
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
	gcloud run services delete website

deploy:
	gcloud run deploy website \
		--port=5000\
		--image=$(IMAGE)\
		--service-account=webapi@fiuba-taller3.iam.gserviceaccount.com\
		--allow-unauthenticated\
		--cpu=1\
		--memory=1Gi\
		--min-instances=0\
		--max-instances=2\
		--region=us-east1\
		--platform=managed\
		--timeout=60s\
		--set-env-vars=ASPNETCORE_ENVIRONMENT=Release\
		--set-env-vars=ASPNETCORE_URLS=http://0.0.0.0:5000\
		--set-env-vars=PubSubVisitCounterService__ProjectId=fiuba-taller3\
		--set-env-vars=PubSubVisitCounterService__TopicId=visit-counter\
		--set-env-vars=PubSubVisitCounterService__IsPublishingEnabled=true\
		--set-env-vars=VisitCounterService__VisitCounterApiBase=https://api-wl63molm2a-ue.a.run.app\
		--set-env-vars=VisitCounterService__GetCountEndpoint=/api/visits/total