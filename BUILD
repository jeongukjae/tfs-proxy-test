load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")
load("@bazel_gazelle//:def.bzl", "gazelle")

buildifier(name = "buildifier")

# gazelle:prefix tfx-proxy-test
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "tfx-proxy-test_lib",
    srcs = ["main.go"],
    importpath = "tfx-proxy-test",
    visibility = ["//visibility:private"],
    deps = [
        "@org_golang_google_grpc//:go_default_library",
        "@com_github_jeongukjae_tfs_proto//:tensorflow_serving_apis_go_proto",
        "@com_github_jeongukjae_tfs_proto//:tensorflow_core_framework_go_proto",
        "@com_github_jeongukjae_tfs_proto//:tensorflow_core_example_go_proto",
    ],
)

go_binary(
    name = "tfx-proxy-test",
    embed = [":tfx-proxy-test_lib"],
    visibility = ["//visibility:public"],
)
