package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "balancer.go",
        "naming.go",
    ],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/grpc",
    deps = [
        "@binchencoder_third_party_go//vendor/github.com/stretchr/testify/mock:go_default_library",
        "@binchencoder_third_party_go//vendor/golang.org/x/net/context:go_default_library",
        "@binchencoder_third_party_go//vendor/google.golang.org/grpc:go_default_library",
        "@binchencoder_third_party_go//vendor/google.golang.org/grpc/naming:go_default_library",
    ],
)
