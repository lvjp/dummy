output "registry" {
  value = scaleway_container_namespace.dummy.registry_endpoint
}

output "graphana_password" {
  value     = scaleway_cockpit_grafana_user.dummy.password
  sensitive = true
}
