IMAGE_TAG=website:release
IMAGE=us-east1-docker.pkg.dev/fiuba-taller3/main/$(IMAGE_TAG)

build-image:
	docker build -t website:latest .

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
	docker tag website:latest $(IMAGE)

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
		--image=$(IMAGE)\
		--port=8080\
		--service-account=website@fiuba-taller3.iam.gserviceaccount.com\
		--allow-unauthenticated\
		--cpu=1\
		--memory=1Gi\
		--min-instances=0\
		--max-instances=2\
		--region=us-east1\
		--platform=managed\
		--timeout=60s\
		--set-env-vars=ENV=production\
		--set-env-vars=GCP_PROJECT_ID=fiuba-taller3\
		--set-env-vars=GCP_PUBSUB_TOPIC_ID=visit-counter\
		--set-env-vars=API_HOST=api-wl63molm2a-ue.a.run.app\
		--set-env-vars=API_PORT=443\
		--set-env-vars=API_ENDPOINT_GET_TOTAL_VISITS=/api/visits/total