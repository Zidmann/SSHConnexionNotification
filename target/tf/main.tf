provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}

data "google_project" "current_project" {
}
