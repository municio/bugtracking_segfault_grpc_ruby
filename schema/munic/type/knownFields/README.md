To update the protobuf schema file you must:

 - Update the field list file `default_fields.json` with `http://gitlab.mobile-intra.com/cloud-core/tracking_fields_reference/raw/master/default_fields.json`
 - Run the generator script that parses `default_fields.json` and prints the updated protobuf: `go run municfields_generate.go > municFields.proto`
 - Ensure differences don't break backward compatibility (no fields id changes or field type change)
 - Commit `default_fields.json` and `municFields.proto`
