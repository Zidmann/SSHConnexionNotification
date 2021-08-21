resource "google_secret_manager_secret" "jwt_signing_key" {
  secret_id = "jwt_signing_key"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "jwt_signing_key" {
  secret = google_secret_manager_secret.jwt_signing_key.id

  secret_data = var.jwt_signing_key
}

