output "outputs" {
  value = {
    network        = google_compute_network.vpc_network.self_link
    public_subnet  = google_compute_subnetwork.public_subnetwork.self_link
    database_ip    = google_sql_database_instance.db_instance.ip_address
    db_name        = google_sql_database.db.name
  }
}
