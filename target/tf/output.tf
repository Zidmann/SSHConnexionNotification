output "sendmsg-svc-url" {
  value = google_cloud_run_service.sendmsg-svc.status[0].url
}

output "notification-gw-url" {
  value = google_api_gateway_gateway.api_gw_notification.default_hostname
}
