package(default_visibility = ["//personalweb:__pkg__"])

load("@io_bazel_rules_closure//closure:defs.bzl", "web_library")
load("//defs:vulcanize.bzl", "html_binary")

web_library(
  name = "index_lib",
  srcs = [
    "app.html",
  ],
  path = "/"
)

html_binary(
  name = "index",
  input_path = "/app.html",
  output_path = "/index.html",
  deps = [
    ":index_lib",
  ]
)