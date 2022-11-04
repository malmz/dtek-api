project = "dtek-api"

app "dtek-api" {
  build {
    use "docker" {}

    registry {
      use "docker" {
        image = "dtek-api"
        tag   = "latest"
      }
    }
  }
  deploy {
    use "nomad" {
      datacenter = "homelab"
    }
  }
}
