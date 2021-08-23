variable "gcp_project_id" {
  type = string
}

variable "gcp_region" {
  type = string
}

variable "sendmsg_app_image" {
  type = string
}

variable "jwt_signing_key" {
  type = string
}

variable "channel_token_list" {
  type = list(object({
    channel = string,
    token   = string
  }))
}

