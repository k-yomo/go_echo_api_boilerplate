output "outputs" {
  value = {
    cluster_username = google_container_cluster.cluster.master_auth.0.username
    cluster_password = google_container_cluster.cluster.master_auth.0.password
    endpoint         = google_container_cluster.cluster.endpoint
  }
}
