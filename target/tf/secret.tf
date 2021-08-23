resource "google_secret_manager_secret" "jwt_signing_key" {
  secret_id = "jwt_signing_key"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret" "channel_tokens" {
  for_each = var.channel_token_list
  secret_id = "discord_${each.value.channel}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "jwt_signing_key" {
  secret      = google_secret_manager_secret.jwt_signing_key.id
  secret_data = var.jwt_signing_key
}

resource "google_secret_manager_secret_version" "channel_tokens" {
  for_each    = var.channel_token_list
  secret      = google_secret_manager_secret.discord_${each.value.channel}.id
  secret_data = each.value.token
}

resource "google_secret_manager_secret_iam_member" "jwt_signing_key" {
  project   = data.google_project.current_project.project_id
  secret_id = google_secret_manager_secret.jwt_signing_key.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.sendmsg-svc.email}"
}

resource "google_secret_manager_secret_iam_member" "channel_tokens" {
  for_each  = var.channel_token_list
  project   = data.google_project.current_project.project_id
  secret_id = google_secret_manager_secret.discord_${each.value.channel}.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.sendmsg-svc.email}"
}

