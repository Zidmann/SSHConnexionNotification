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

resource "google_cloud_run_service_iam_member" "api-gw-notification" {
  service  = google_cloud_run_service.sendmsg-svc.name
  role     = "roles/viewer"
  member   = "serviceAccount:${google_service_account.api-gw-notification.email}"
}
