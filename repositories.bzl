load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "binchencoder_third_party_go",
        commit = "d315c8c6aeab36114ee245515150906d434e92b3",
        importpath = "gitee.com/binchencoder/third-party-go",
    )
    go_repository(
        name = "binchencoder_ease_gateway",
        commit = "267a6471666091ec09c88a3eae061b31162c5bf7",
        importpath = "gitee.com/binchencoder/ease-gateway",
    )
    go_repository(
        name = "binchencoder_letsgo",
        commit = "6adc4f84411faa8b1d9fef647d222b6979f175d7",
        importpath = "gitee.com/binchencoder/letsgo",
    )

    go_repository(
        name = "grpc_ecosystem_grpc_gateway",
        commit = "ad529a448ba494a88058f9e5be0988713174ac86",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
    )
    go_repository(
        name = "com_github_fatih_color",
        commit = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
        importpath = "github.com/fatih/color",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        commit = "ae52aff18558bd92cbe681549bfe9e8cbffd5903",
        importpath = "github.com/klauspost/compress",
    )
    go_repository(
        name = "com_github_klauspost_cpuid",
        commit = "05a8198c0f5a27739aec358908d7e12c64ce6eb7",
        importpath = "github.com/klauspost/cpuid",
    )
    # go_repository(
    #     name = "com_github_golang_net",
    #     commit = "4829fb13d2c62012c17688fa7f629f371014946d",
    #     importpath = "github.com/golang/net",
    # )
    # go_repository(
    #     name = "com_github_golang_protobuf",
    #     commit = "4c88cc3f1a34ffade77b79abc53335d1e511f25b",
    #     importpath = "github.com/golang/protobuf",
    # )
    # go_repository(
    #     name = "org_golang_google_grpc",
    #     importpath = "google.golang.org/grpc",
    #     commit = "383e8b2c3b9e36c4076b235b32537292176bae20",
    # )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        commit = "6d4652c779c4add9e1a69db058dabafddba21c37",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        commit = "6724a57986aff9bff1a1770e9347036def7c89f6",
        importpath = "github.com/rogpeppe/fastuuid",
    )
    go_repository(
        name = "in_gopkg_resty_v1",
        commit = "fa5875c0caa5c260ab78acec5a244215a730247f",
        importpath = "gopkg.in/resty.v1",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        commit = "0ca9ea5df5451ffdf184b4428c902747c2c11cd7",
        importpath = "github.com/ghodss/yaml",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
        importpath = "gopkg.in/yaml.v2",
    )