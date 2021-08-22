resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "api-gw"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "config"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = templatefile("${path.module}/openapi.yaml.template", { sendmsg_svc_adress = "{google_cloud_run_service.sendmsg-svc.status[0].url}" })
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "api-gw"
}
