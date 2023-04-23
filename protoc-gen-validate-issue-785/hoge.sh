vpath=""
protoc \
  --validate_out=lang=go:./gen \
  -I . \
  -I ${vpath} example.proto
