package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cli.go",
    ],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/etcd",
    deps = [
        "@binchencoder_third_party_go//vendor/github.com/coreos/etcd/client:go_default_library",
        "@binchencoder_third_party_go//vendor/github.com/stretchr/testify/mock:go_default_library",
        "@binchencoder_third_party_go//vendor/golang.org/x/net/context:go_default_library",
    ],
)
