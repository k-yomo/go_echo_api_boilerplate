resource "google_container_cluster" "cluster" {
  name       = "gke-cluster-${var.env}"
  location   = var.location
  network    = var.network
  subnetwork = var.subnetwork

  // cluster can't be created with no node pool, so we need at least 1 node here.
  remove_default_node_pool = true
  initial_node_count       = 1

  ip_allocation_policy {
    use_ip_aliases = true
  }
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "gke-primary-node-pool-${var.env}"
  location   = var.location
  cluster    = google_container_cluster.cluster.name
  node_count = var.primary_node.node_count

  management {
    auto_repair = true
  }

  autoscaling {
    max_node_count = 1
    min_node_count = 1
  }

  node_config {
    machine_type = var.primary_node.machine_type

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}

resource "google_container_node_pool" "preemptible_nodes" {
  name       = "gke-preemptible-node-pool-${var.env}"
  location   = var.location
  cluster    = google_container_cluster.cluster.name
  node_count = var.preemptible_node.node_count

  management {
    auto_repair = true
  }

  autoscaling {
    max_node_count = 3
    min_node_count = 1
  }

  node_config {
    preemptible  = true
    machine_type = var.preemptible_node.machine_type

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}
