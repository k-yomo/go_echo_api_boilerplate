data "terraform_remote_state" "infra_state" {
  backend = "gcs"

  config = {
    bucket = var.infra_bucket_name
  }
}

module "kubernetes-engine" {
  source           = "../../modules/kubernetes-engine"
  env              = var.env
  location         = var.gcp_region
  network          = data.terraform_remote_state.infra_state.outputs.outputs.network
  subnetwork       = data.terraform_remote_state.infra_state.outputs.outputs.public_subnet
  primary_node     = var.primary_node
  preemptible_node = var.preemptive_node
}

