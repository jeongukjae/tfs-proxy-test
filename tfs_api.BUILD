load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "tensorflow_core_framework_proto",
    srcs = glob(["tensorflow/core/framework/**"]),
)

proto_library(
    name = "tensorflow_core_example_proto",
    srcs = glob(["tensorflow/core/example/**"]),
)

proto_library(
    name = "tensorflow_core_protobuf_proto",
    srcs = glob(["tensorflow/core/protobuf/**"]),
    deps = [
        ":tensorflow_core_framework_proto",
        "@com_google_protobuf//:any_proto",
    ],
)

proto_library(
    name = "tensorflow_serving_config_proto",
    srcs = glob(["tensorflow_serving/config/**"]),
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
)

proto_library(
    name = "tensorflow_serving_apis_proto",
    srcs = glob(["tensorflow_serving/apis/**"]),
    deps = [
        ":tensorflow_core_framework_proto",
        ":tensorflow_core_example_proto",
        ":tensorflow_core_protobuf_proto",
        ":tensorflow_serving_config_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:wrappers_proto",
    ],
)

go_proto_library(
    name = "tensorflow_serving_apis_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/tensorflow/serving/tensorflow_serving/apis",
    proto = "//:tensorflow_serving_apis_proto",
    deps = [
        ":tensorflow_core_example_go_proto",
        ":tensorflow_core_framework_go_proto",
        ":tensorflow_core_protobuf_go_proto",
        ":tensorflow_serving_config_go_proto",
    ],
)

go_proto_library(
    name = "tensorflow_serving_config_go_proto",
    importpath = "github.com/tensorflow/serving/tensorflow_serving/config",
    proto = "//:tensorflow_serving_config_proto",
)

# tensorflow protos
go_proto_library(
    name = "tensorflow_core_framework_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/framework",
    proto = "//:tensorflow_core_framework_proto",
)

go_proto_library(
    name = "tensorflow_core_example_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/example",
    proto = "//:tensorflow_core_example_proto",
)

go_proto_library(
    name = "tensorflow_core_protobuf_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf",
    proto = "//:tensorflow_core_protobuf_proto",
    deps = [
        ":tensorflow_core_framework_go_proto",
    ],
)
