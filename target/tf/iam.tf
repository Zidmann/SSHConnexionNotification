resource "google_service_account" "sendmsg-svc" {
  account_id   = "sendmsg-svc"
  display_name = "Service account for the Cloud Run send message service"
}

