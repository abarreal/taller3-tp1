gcloud-login:
	gcloud auth login

gcloud-logout:
	gcloud auth revoke

undeploy:
	gcloud functions delete visitcount --region=us-east1

deploy:
	gcloud functions deploy visitcount\
		--entry-point=CloudFunctionsRun\
		--runtime=go116\
		--trigger-topic=visit-counter\
		--timeout=60s\
		--memory=128MB\
		--max-instances=2\
		--region=us-east1\
		--retry\
		--service-account=visitcounter@fiuba-taller3.iam.gserviceaccount.com\
		--env-vars-file=env.yaml