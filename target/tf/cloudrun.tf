provider "google" {
  project     = var.gcp_project_id
  region      = var.gcp_region
}


resource "google_project_service" "run" {
  service = "run.googleapis.com"
}

# users svc
resource "google_cloud_run_service" "sendmsg-svc" {
  name     = "sendmsg-svc"
  location = var.gcp_region

  template {
    spec {
      containers {
        image = var.sendmsg_app_image
        env {
          name = "PROJECT_NAME"
          value = var.gcp_project_id
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

    depends_on = [google_project_service.run]
}

