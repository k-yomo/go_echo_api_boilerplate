provider "google" {
  version = "~> 2.8.0"
  project = var.project_name
  region  = var.gcp_region
  zone    = var.availability_zones[0]
}

