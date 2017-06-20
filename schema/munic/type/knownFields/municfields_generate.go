package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/doc"
	"io/ioutil"
	"strings"
)

var header string = `syntax = "proto3";

import "munic/type/utils/boxed.proto";
import "google/type/latlng.proto";

package munic.type.knownFields;

option go_package = "gitlab.mobile-intra.com/cloud-next/schema/compiled/go/munic/types/knownFields";
option java_package = "com.munic.model.type.knownFields";
option java_outer_classname = "MunicFieldsProto";
option java_multiple_files = true;
option java_generate_equals_and_hash = true;
`

type Field struct {
	Name        string
	Id          int    `json:"field"`
	Type        string `json:"field_type"`
	Description string
}

var typeMap = map[string]string{
	"integer": "munic.type.utils.Integer",
	"string":  "munic.type.utils.String",
	"boolean": "munic.type.utils.Boolean",
	"bytes":   "munic.type.utils.Bytes",
	"latlng":  "google.type.LatLng",
}

func (f *Field) SpecialCases() *Field {
	switch f.Name {
	case "POSITION":
		return &Field{Name: f.Name, Id: 1, Type: "latlng", Description: "Stores the position of the asset."}
	case "DIO_IN_TOR", "DIO_OUT_TOR",
		"TACHOGRAPH_VIN",
		"MDI_OBD_VIN", "MDI_OBD_VIN_ALT", "MDI_OBD_VIN_HASH", "MDI_OBD_SQUISH_VIN",
		"MDI_OBD_PID_1", "MDI_OBD_PID_2", "MDI_OBD_PID_3", "MDI_OBD_PID_4", "MDI_OBD_PID_5",
		"MDI_SENSORS_RECORDER_DATA", "MDI_SENSORS_RECORDER_CALIBRATION",
		"MDI_DIAG_1", "MDI_DIAG_2", "MDI_DIAG_3":
		f.Type = "bytes"
		return f
	default:
		return f
	}
}

func (f *Field) EncodeProto() string {
	f = f.SpecialCases()
	ftype, ok := typeMap[f.Type]
	if !ok {
		ftype = f.Type
	}

	description := strings.Replace(f.Description, "\r\n", " ", -1)
	description = strings.Replace(f.Description, "\n", " ", -1)
	description = strings.Replace(description, "\r", " ", -1)
	if description == "" {
		description = "TODO: comment this field"
	}
	description = wrap(description, 80)
	description = strings.TrimSuffix(description, "\n")
	if len(description) > 0 {
		return fmt.Sprintf("%s\n\t%s %s = %d;\n", description, ftype, f.Name, f.Id)
	} else {
		return fmt.Sprintf("\t%s %s = %d;\n", ftype, f.Name, f.Id)
	}
}

func wrap(text string, cols int) string {
	var buf bytes.Buffer
	doc.ToText(&buf, text, "\t// ", "// ", 80)
	return buf.String()
}

func get_content(fname string) []Field {
	body, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err.Error())
	}

	var fields []Field
	json.Unmarshal(body, &fields)
	return fields
}

func main() {
	fs := get_content("default_fields.json")
	fmt.Println(header)
	fmt.Println("message MunicFields {")
	sep := ""
	for _, f := range fs {
		if f.Id > 250 {
			continue
		}
		fmt.Printf("%s%s", sep, f.EncodeProto())
		sep = "\n"
	}
	fmt.Println("}")
}
