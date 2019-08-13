resource "google_service_networking_connection" "private_vpc_connection" {
  provider                = "google-beta"
  network                 = google_compute_network.vpc_network.self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "google_sql_database_instance" "db_instance" {
  provider         = "google-beta"
  project          = var.project_name
  name             = var.instance_name
  database_version = "MYSQL_5_7"
  region           = var.region
  settings {
    tier = var.machine_type
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.vpc_network.self_link
    }
    backup_configuration {
      enabled    = true
      start_time = "20:00"
    }
  }

  depends_on = [google_service_networking_connection.private_vpc_connection]
}

resource "google_sql_database" "db" {
  instance = google_sql_database_instance.db_instance.name
  name     = var.db_name
}

