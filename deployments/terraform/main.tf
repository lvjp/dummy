resource "scaleway_container_namespace" "dummy" {
  name = "dummy"

  destroy_registry = true
}

resource "scaleway_container" "dummy" {
  name         = "dummy"
  namespace_id = scaleway_container_namespace.dummy.id

  deploy       = true
  privacy      = "public"
  memory_limit = 128
  timeout      = 30

  registry_image = "rg.fr-par.scw.cloud/funcscwdummyufo53dla/dummy:latest"
}