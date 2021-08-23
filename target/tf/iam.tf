resource "google_service_account" "sendmsg-svc" {
  account_id   = "sendmsg-svc"
  display_name = "Service account for the Cloud Run send message service"
}

resource "google_service_account" "api-gw-notification" {
  account_id   = "api-gw-notification"
  display_name = "Service account for the API Gateway to expose message service"
}

