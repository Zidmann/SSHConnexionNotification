SENDMSG_APP_VERSION=$(shell cat services/send_message/VERSION)
SENDMSG_APP_IMAGE=gcr.io/$(PROJECT_ID)/sendmessage:$(SENDMSG_APP_VERSION)

gcp-enable-apis:
	gcloud services enable cloudresourcemanager.googleapis.com --project $(PROJECT_ID)
	gcloud services enable cloudbuild.googleapis.com --project $(PROJECT_ID)
	gcloud services enable apigateway.googleapis.com --project $(PROJECT_ID)
	gcloud services enable servicemanagement.googleapis.com --project $(PROJECT_ID)
	gcloud services enable servicecontrol.googleapis.com --project $(PROJECT_ID)

gcr-push-services:
	(cd services/send_message &&\
	gcloud builds submit --tag $(SENDMSG_APP_IMAGE) --project $(PROJECT_ID))

tf-init:
	(cd tf &&\
	terraform init)

tf-plan:
	(cd tf &&\
	terraform plan -out=out.tfplan)

tf-apply:
	(cd tf &&\
	terraform apply out.tfplan)

tf-destroy:
	(cd tf &&\
	terraform destroy)

get-gateway-url:	
	$(eval GATEWAY_URL=$(shell gcloud beta api-gateway gateways describe $(GATEWAY_DEMO_ID) \
	--location=$(REGION) --project=$(PROJECT_ID) \
	--format=json|jq ".defaultHostname"))
	@echo https://$(GATEWAY_URL)

