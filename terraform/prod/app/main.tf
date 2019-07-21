module "vpc_network" {
  source       = "../../modules/network"
  network_name = "vpc-network-${var.env}"
}

module "database" {
  source          = "../../modules/database"
  machine_type    = "db-f1-micro"
  db_name         = "fast_ticket_${var.env}"
  region          = var.gcp_region
  network = lookup(module.vpc_network.outputs, "network")
}

module "kubernetes-engine" {
  source             = "../../modules/kubernetes-engine"
  env                = var.env
  location           = var.gcp_region
  network = lookup(module.vpc_network, "network", null)
  initial_node_count = 1
  primary_node = {
    node_count   = 1
    machine_type = "f1-small"
  }

  preemptible_node = {
    node_count   = 2
    machine_type = "f1-micro"
  }
}

