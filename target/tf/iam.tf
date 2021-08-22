resource "google_service_account" "sendmsg-svc" {
  account_id   = "sendmsg-svc"
  display_name = "Service account for the Cloud Run send message service"
}

resource "google_project_iam_member" "sendmsg-svc-secret" {
  project = data.google_project.current_project.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${google_service_account.sendmsg-svc.email}"
}
