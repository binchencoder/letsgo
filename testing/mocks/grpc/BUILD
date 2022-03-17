package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "balancer.go",
    ],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/grpc",
    deps = [
        "@com_github_stretchr_testify//mock:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
