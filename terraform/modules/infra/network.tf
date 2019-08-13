locals {
  public  = "public"
  private = "private"
}

resource "google_compute_network" "vpc_network" {
  name                    = var.network_name
  auto_create_subnetworks = "false"
  routing_mode            = "REGIONAL"
}

resource "google_compute_global_address" "private_ip_address" {
  project       = var.project_name
  provider      = "google-beta"
  name          = "private-ip-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.vpc_network.self_link
}

resource "google_compute_router" "vpc_router" {
  name    = "${var.network_name}-router"
  region  = var.region
  network = google_compute_network.vpc_network.self_link
}

resource "google_compute_subnetwork" "public_subnetwork" {
  name                     = "${var.network_name}-public"
  ip_cidr_range            = var.public_ip_cidr_range
  private_ip_google_access = true
  network                  = google_compute_network.vpc_network.self_link
  region                   = var.region
}

resource "google_compute_router_nat" "vpc_nat" {
  name                               = "${var.network_name}-nat"
  region                             = var.region
  router                             = google_compute_router.vpc_router.name
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"

  subnetwork {
    name                    = google_compute_subnetwork.public_subnetwork.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}

resource "google_compute_subnetwork" "private_subnetwork" {
  name                     = "${var.network_name}-private"
  ip_cidr_range            = var.private_ip_cidr_range
  private_ip_google_access = true
  network                  = google_compute_network.vpc_network.self_link
  region                   = var.region
}

resource "google_compute_firewall" "public_allow_all_inbound" {
  name    = "${var.network_name}-public-allow-ingress"
  network = google_compute_network.vpc_network.self_link

  target_tags   = [local.public]
  direction     = "INGRESS"
  source_ranges = ["0.0.0.0/0"]

  priority = "1000"

  allow {
    protocol = "all"
  }
}

resource "google_compute_firewall" "private_allow_all_network_inbound" {
  name    = "${var.network_name}-private-allow-ingress"
  network = google_compute_network.vpc_network.self_link

  target_tags = [local.private]
  direction   = "INGRESS"

  source_ranges = [
    google_compute_subnetwork.public_subnetwork.ip_cidr_range,
    google_compute_subnetwork.private_subnetwork.ip_cidr_range,
  ]

  priority = "1000"

  allow {
    protocol = "all"
  }
}