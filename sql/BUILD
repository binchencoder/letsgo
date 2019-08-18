package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["sql.go"],
    importpath = "binchencoder.com/letsgo/sql",
    deps = ["@io_upper_db_v3//lib/sqlbuilder:go_default_library"],
)

go_test(
    name = "small_tests",
    size = "small",
    srcs = [],
    embed = [
        ":go_default_library",
    ],
    deps = [
        "@com_github_smartystreets_goconvey//convey:go_default_library",
        "@io_upper_db_v3//:go_default_library",
        "@io_upper_db_v3//lib/sqlbuilder:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["sql_test.go"],
    embed = [":go_default_library"],
    deps = [
        "@com_github_smartystreets_goconvey//convey:go_default_library",
        "@io_upper_db_v3//:go_default_library",
        "@io_upper_db_v3//lib/sqlbuilder:go_default_library",
    ],
)
