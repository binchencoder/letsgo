package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "sql.go",
    ],
    importpath = "github.com/binchencoder/letsgo/sql",
    deps = ["//third-party-go/vendor/upper.io/db.v3/lib/sqlbuilder:go_default_library"],
)

go_test(
    name = "small_tests",
    size = "small",
    srcs = ([
        "sql_test.go",
    ]),
    embed = [
        ":go_default_library",
    ],
    deps = [
        "//third-party-go/vendor/github.com/smartystreets/goconvey/convey:go_default_library",
        "//third-party-go/vendor/upper.io/db.v3:go_default_library",
        "//third-party-go/vendor/upper.io/db.v3/lib/sqlbuilder:go_default_library",
    ],
)
