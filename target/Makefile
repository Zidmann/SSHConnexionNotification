SENDMSG_APP_VERSION=$(shell cat services/send_message/VERSION)
SENDMSG_APP_IMAGE=gcr.io/$(PROJECT_ID)/sendmessage:$(SENDMSG_APP_VERSION)

gcp-enable-apis:
	gcloud services enable cloudresourcemanager.googleapis.com --project $(PROJECT_ID)
	gcloud services enable cloudbuild.googleapis.com --project $(PROJECT_ID)
	gcloud services enable apigateway.googleapis.com --project $(PROJECT_ID)
	gcloud services enable servicemanagement.googleapis.com --project $(PROJECT_ID)
	gcloud services enable servicecontrol.googleapis.com --project $(PROJECT_ID)
	gcloud services enable secretmanager.googleapis.com --project $(PROJECT_ID)

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

get-sendmsg-url:
	$(eval SENDMSG_SVC_URL=$(shell cd tf && \
					terraform output sendmsg-svc-url))
	@echo SENDMSG_SVC_URL: $(SENDMSG_SVC_URL)

get-notification-url:
	$(eval NOTIFICATION_GW_URL=$(shell cd tf && \
					    terraform output notification-gw-url))
	@echo NOTIFICATION_GW_URL: $(NOTIFICATION_GW_URL)

