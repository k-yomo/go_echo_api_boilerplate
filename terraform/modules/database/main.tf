resource "google_sql_database_instance" "db_instance" {
  name             = var.db_name
  database_version = "MYSQL_5_7"
  region           = var.region
  settings {
    tier = var.machine_type
    ip_configuration {
      private_network = var.network
    }
  }
}

resource "google_sql_database" "db" {
  instance = google_sql_database_instance.db_instance.name
  name     = var.db_name
}
