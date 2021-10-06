IMAGE_TAG=aggregator:release
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
	gcloud compute instances delete aggregator-instance

deploy:
	echo "Does not initiate the container for unknown reasons; create manually"
	echo gcloud compute instances create-with-container\
		aggregator-instance\
		--container-image=$(IMAGE)\
		--container-env-file=cloud.env\
		--service-account=visitcounter@fiuba-taller3.iam.gserviceaccount.com\
		--network-tier=STANDARD\
		--custom-cpu=2\
		--custom-memory=1024MB\
		--no-address