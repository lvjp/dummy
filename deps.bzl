load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_afex_hystrix_go",
        importpath = "github.com/afex/hystrix-go",
        sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
        version = "v0.0.0-20180502004556-fa1af6a1f4f5",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:ANLJcKmQm4nIaog7xdr/id6FM6zm5hHnfZrvtKPxqGg=",
        version = "v2.3.1",
    )

    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )

    go_repository(
        name = "com_github_armon_go_metrics",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:O2sNqxBdvq8Eq5xmzljcYzAORli6RWCvEym4cJf9m18=",
        version = "v0.3.9",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:QN1nsY27ssD/JmW4s83qmSb+uL6DG4GmCDzjmJB4xUI=",
        version = "v1.40.45",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:ZbovGV/qo40nrOJ4q8G33AGICzaPI45FHQWJ9650pF4=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_cloudwatch",
        importpath = "github.com/aws/aws-sdk-go-v2/service/cloudwatch",
        sum = "h1:w/fPGB0t5rWwA43mux4e9ozFSH5zF1moQemlA131PWc=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:AEwwwXQZtUwP5Mz506FeXXrKBe0jA8gVM+1gEcSRooc=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_casbin_casbin_v2",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:/poEwPSovi4bTOcP752/CsTQiRz2xycyVKFG7GUhbDw=",
        version = "v2.37.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:G2HAfAmvm/GcKan2oOQpBXOd2tT2G57ZnZGWa1PxPBQ=",
        version = "v4.1.1",
    )

    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )

    go_repository(
        name = "com_github_clbanning_mxj",
        importpath = "github.com/clbanning/mxj",
        sum = "h1:HuhwZtbyvyOw+3Z1AowPkU87JkJUSv751ELWaiTpj8I=",
        version = "v1.8.4",
    )

    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )

    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:mRhaKNwANqRgUBGKmnI5ZxEk7QXmjQeCcuYFMX2bfcc=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_github_franela_goreq",
        importpath = "github.com/franela/goreq",
        sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
        version = "v0.0.0-20171204163338-bcd34c9993f8",
    )

    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        sum = "h1:e4o3o3IsBfAKQh5Qbbiqyfu97Ku7jrO/JbohvztANh4=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )

    go_repository(
        name = "com_github_go_zookeeper_zk",
        importpath = "github.com/go-zookeeper/zk",
        sum = "h1:4mx0EYENAdX/B/rbunjlt5+4RTA/a9SMHBRuSKdGxPM=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
        version = "v5.0.4",
    )

    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )

    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:RAqyYixv1p7uEnocuy8P1nru5wprCh/MH2BIlW5z5/o=",
        version = "v4.0.0",
    )

    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )

    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )

    go_repository(
        name = "com_github_hashicorp_consul_api",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:MwZJp86nlnL+6+W1Zly4JUuVn9YHhMggBirMpHGD7kw=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:K4ev2ib4LdQETX5cSZBG0DVLk1jwGqSPXBjdah3veNs=",
        version = "v0.16.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:EBWvyu9tcRszt3Bxp3KNssBMP1KuHWyO51lz9+786iM=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_github_hdrhistogram_hdrhistogram_go",
        importpath = "github.com/HdrHistogram/hdrhistogram-go",
        sum = "h1:5IcZpTvzydCQeHzK4Ef/D5rrSqwxob0t8PQPMybUNFM=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_hudl_fargo",
        importpath = "github.com/hudl/fargo",
        sum = "h1:ZDDILMbB37UlAVLlWcJ2Iz1XuahZZTDZfdCKeclfq2s=",
        version = "v1.4.0",
    )

    go_repository(
        name = "com_github_influxdata_influxdb1_client",
        importpath = "github.com/influxdata/influxdb1-client",
        sum = "h1:HqW4xhhynfjrtEiiSGcQUd6vrK23iMam1FO8rI7mwig=",
        version = "v0.0.0-20200827194710-b269163b24ab",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )

    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )

    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:P76CopJELS0TiO2mebmnzgWaajssP/EszplttgQxcgc=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_github_knetic_govaluate",
        importpath = "github.com/Knetic/govaluate",
        sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
        version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
    )

    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )

    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:DOKFKCQ7FNG2L1rbrmstDN4QVRdS89Nkh85u68Uwp98=",
        version = "v0.0.18",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_miekg_dns",
        importpath = "github.com/miekg/dns",
        sum = "h1:JKfpVSCB84vrAmHzyrsxB5NAr5kLoMXZArPSw7Qlgyg=",
        version = "v1.1.43",
    )
    go_repository(
        name = "com_github_minio_highwayhash",
        importpath = "github.com/minio/highwayhash",
        sum = "h1:Aak5U0nElisjDCfPSG79Tgzkn2gl66NxOMspRrKnA/g=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:6h7AQ0yhTcIsmFmnAwQls75jp2Gzs4iB8W7pjMO+rqo=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )

    go_repository(
        name = "com_github_nats_io_jwt_v2",
        importpath = "github.com/nats-io/jwt/v2",
        sum = "h1:i/O6cmIsjpcQyWDYNcq2JyZ3/VTF8SJ4JWluI5OhpvI=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:+0ndxwUPz3CmQ2vjbXdkC1fo3FdiOQDim4gl3Mge8Qo=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_github_nats_io_nats_server_v2",
        importpath = "github.com/nats-io/nats-server/v2",
        sum = "h1:wsnVaaXH9VRSg+A2MVg5Q727/CqxnmPLGFQ3YZYKTQg=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:cgM5tL53EvYRU+2YLXIK0G2mJtK12Ft9oeooSZMA2G8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_op_go_logging",
        importpath = "github.com/op/go-logging",
        sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
        version = "v0.0.0-20160315200505-970db520ece7",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:UwtQQx2pyPIgWYHRg+epgdx1/HnBQTgN3/oIYEJTQzU=",
        version = "v0.2.5",
    )
    go_repository(
        name = "com_github_performancecopilot_speed_v4",
        importpath = "github.com/performancecopilot/speed/v4",
        sum = "h1:VxEDCmdkfbQYDlcr/GC9YoN9PQ6p8ulk9xVsepYy9ZY=",
        version = "v4.0.0",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )

    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:5fCgGYogn0hFdhyhLbw7hEsWxufKtY9klyvdNfFlFhM=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:UBgGFHqYdG/TPFD1B1ogZywDqEkwp3fBMvqdiQ7Xew4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:EKsfXEYo4JpWMHH5cg+KOUWeuJSov1Id8zGR8eeI1YM=",
        version = "v0.42.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:wzCHvIvM5SxWqYvwgVL7yJY8Lz3PKn49KQtpgMYJfhI=",
        version = "v0.9.0",
    )

    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:73kH8U+JUqXU8lRuOHeVHaa/SZPifC7BkcraZVejAe8=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:qd7wPTDkN6KQx2VmMBLrpHkiyQwgFXRnkOLacUiaSNY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        importpath = "github.com/rs/zerolog",
        sum = "h1:cO+d60CHkknCbvzEWxP0S9K6KqyTjrCNUy1LdQLCGPc=",
        version = "v1.29.1",
    )

    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_sony_gobreaker",
        importpath = "github.com/sony/gobreaker",
        sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
        version = "v0.4.1",
    )

    go_repository(
        name = "com_github_sourcegraph_conc",
        importpath = "github.com/sourcegraph/conc",
        sum = "h1:OQTbbt6P72L20UqAkXXuLOj79LfEanQ+YQFNpLA9ySo=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_streadway_amqp",
        importpath = "github.com/streadway/amqp",
        sum = "h1:kuuDrUJFZL1QYL9hUNuCxNObNzB0bV/ZG5jV3RWAQgo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_streadway_handy",
        importpath = "github.com/streadway/handy",
        sum = "h1:mOtuXaRAbVZsxAHVdPR3IjfmN8T1h2iczJLynhLybf8=",
        version = "v0.0.0-20200128134331-0f66f006fb2e",
    )

    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )

    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_vividcortex_gohistogram",
        importpath = "github.com/VividCortex/gohistogram",
        sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration",
        importpath = "github.com/xhit/go-str2duration",
        sum = "h1:BcV5u025cITWxEQKGWr1URRzrcXtu7uk8+luz3Yuhwc=",
        version = "v1.2.0",
    )

    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )

    go_repository(
        name = "in_gopkg_gcfg_v1",
        importpath = "gopkg.in/gcfg.v1",
        sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
        version = "v1.2.3",
    )
    go_repository(
        name = "in_gopkg_warnings_v0",
        importpath = "gopkg.in/warnings.v0",
        sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
        version = "v0.1.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )

    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:GsV3S+OfZEOCNXdtNkBSR7kgLobAa/SO6tCxRa0GAYw=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:2aQv6F436YnN7I4VbI8PPYrBhu+SmrTaADcf8Mi/6PU=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:ftQ0nOOHMcbMS3KIaDQ0g5Qcd6bhaBrQT6b89DfwLTs=",
        version = "v2.305.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:62Eh0XOro+rDwkrypAGDfgmNh5Joq+z+W9HZdlXMzek=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:gqCw0LfLxScz8irSi8exQc7fyQ0fKQU/qnC/X8+V/1M=",
        version = "v0.23.0",
    )

    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
        version = "v1.6.7",
    )

    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:ysnBoUyeL/H6RCvNRhWHjKoDEmguI+mPU+qHgK8qv/w=",
        version = "v0.0.0-20210917145530-b395a37504d4",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:AGJ0Ih4mHjSeibYkFGh1dD9KJ/eOtZ93I6hoHhukQ5Q=",
        version = "v1.40.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:kPPoIgf3TsEvrm0PFe15JQ+570QVxYzEvvHqChK+cng=",
        version = "v1.30.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:3erb+vDS8lU1sxfDHF4/hhWyaXnhIaO+7RgL4fDZORA=",
        version = "v0.0.0-20210915214749-c084706c2272",
    )

    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:rJrUqqhjsgNp7KqAIc25s9pZnjU7TUcSY7HcVZjdn1g=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:HuArIo48skDwlrvM3sEdHXElYslAMsf3KwRkkW4MC4s=",
        version = "v0.5.0",
    )

    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:wsuoTGHzEhffawBOhz5CYhcrV4IdKZbEyZjBMuTp12o=",
        version = "v0.1.0",
    )

    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:3jlCCIQZPdOYu1h8BkNvLz8Kgwtae2cagcG/VamtZRU=",
        version = "v0.7.0",
    )

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:4BRB4x83lYWy72KwLD/qYDuTu7q9PjSagHvijDw7cLo=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:7zkz7BUtwNFFqcowJ+RIgu2MaV/MapERkDIy+mwPyjs=",
        version = "v0.0.0-20210723032227-1f47c861a9ac",
    )

    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )

    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:ue41HOKd1vGURxrmeKIgELGb3jPW9DMUDGtsinblHwI=",
        version = "v1.19.1",
    )
