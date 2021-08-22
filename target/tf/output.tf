output "sendmsg-svc-url" {
  value = google_cloud_run_service.sendmsg-svc.status[0].url
}
