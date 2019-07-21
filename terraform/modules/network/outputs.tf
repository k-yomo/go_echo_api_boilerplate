output "outputs" {
  value = {
    network = google_compute_network.vpc_network.self_link
  }
}
