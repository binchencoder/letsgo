package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cli.go"],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/etcd",
    deps = [
        "@com_github_coreos_etcd//client:go_default_library",
        "@com_github_stretchr_testify//mock:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
