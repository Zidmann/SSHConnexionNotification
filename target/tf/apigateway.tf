resource "google_api_gateway_api" "api_gw_notification" {
  provider = google-beta
  api_id   = "gateway-notification"
}

resource "google_api_gateway_api_config" "api_gw_notification" {
  provider      = google-beta
  api           = google_api_gateway_api.api_gw_notification.api_id
  api_config_id = "config"

  openapi_documents {
    document {
      path     = "spec.yaml"
      contents = base64encode(templatefile("${path.module}/openapi.yaml.template", { sendmsg_svc_adress = google_cloud_run_service.sendmsg-svc.status[0].url }))
    }
  }

  gateway_config {
    backend_config {
      google_service_account = google_service_account.api-gw-notification.email
    }
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw_notification" {
  provider   = google-beta
  region     = var.gcp_region
  api_config = google_api_gateway_api_config.api_gw_notification.id
  gateway_id = "gateway-notification"
}
