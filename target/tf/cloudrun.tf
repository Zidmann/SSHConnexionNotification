resource "google_project_service" "run" {
  service = "run.googleapis.com"
}

# Send message service
resource "google_cloud_run_service" "sendmsg-svc" {
  name     = "sendmsg-svc"
  location = var.gcp_region

  template {
    spec {
      containers {
        image = "gcr.io/${var.gcp_project_id}/sendmessage:${var.sendmsg_app_version}"
        env {
          name  = "PROJECT_ID"
          value = data.google_project.current_project.number
        }
      }
      service_account_name = google_service_account.sendmsg-svc.email
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on = [google_project_service.run]
}

resource "google_cloud_run_service_iam" "api-gw-notification" {
  project   = data.google_project.current_project.project_id
  secret_id = google_secret_manager_secret.jwt_signing_key.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${api-gw-notification.sendmsg-svc.email}"
}
