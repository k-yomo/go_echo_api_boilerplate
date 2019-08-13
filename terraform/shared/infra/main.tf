module "infra" {
  source                = "../../modules/infra"
  project_name          = var.project_name
  network_name          = "vpc-network-${var.env}"
  public_ip_cidr_range  = "10.0.0.0/24"
  private_ip_cidr_range = "10.0.1.0/24"
  region                = var.gcp_region
  machine_type          = var.db_machine_type
  instance_name         = "${var.service_name}-${var.env}-mysql-db"
  db_name               = "${var.service_name}_${var.env}"
}

