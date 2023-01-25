package swagger

import (
	"sort"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Endpoint struct {
	ConditionalProperties []string `json:"conditionalProperties"`
	Method                string   `json:"method"`
	Path                  string   `json:"path"`
	StatusCode            int      `json:"status_code"`
}

type NonManadatoryFields struct {
	SwaggerPath string     `json:"swagger_path"`
	Endpoints   []Endpoint `json:"endpoints"`
}

func ParseSchema(swaggerPath string, logger *logrus.Logger) (*NonManadatoryFields, error) {
	doc, err := loads.Spec(swaggerPath)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":         err,
			"swaggerPath": swaggerPath,
		}).Debug("ParseSchema: loads.Spec(swaggerPath)")
		return nil, errors.Wrapf(err, "ParseSchema: loads.Spec(path), swaggerPath=%+v", swaggerPath)
	}

	expanded, err := doc.Expanded(nil)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":         err,
			"doc":         doc,
			"swaggerPath": swaggerPath,
		}).Debug("ParseSchema: doc.Expanded(nil)")
		return nil, errors.Wrapf(err, "ParseSchema: doc.Expanded(nil), swaggerPath=%+v", swaggerPath)
	}

	nonManadatoryFields := &NonManadatoryFields{
		SwaggerPath: swaggerPath,
		Endpoints:   []Endpoint{},
	}
	spec := expanded.Spec()
	for path, props := range spec.Paths.Paths {
		// All the list of methods:
		// /Users/mbana/go/pkg/mod/github.com/go-openapi/spec@v0.17.2/path_item.go
		if props.Get != nil {
			nonManadatoryFields.PrintResponses(props.Get, path, "GET", logger)
		}
		if props.Put != nil {
			nonManadatoryFields.PrintResponses(props.Put, path, "PUT", logger)
		}
		if props.Post != nil {
			nonManadatoryFields.PrintResponses(props.Post, path, "POST", logger)
		}
		if props.Delete != nil {
			nonManadatoryFields.PrintResponses(props.Delete, path, "DELETE", logger)
		}
		if props.Options != nil {
			nonManadatoryFields.PrintResponses(props.Options, path, "OPTIONS", logger)
		}
		if props.Head != nil {
			nonManadatoryFields.PrintResponses(props.Head, path, "HEAD", logger)
		}
		if props.Patch != nil {
			nonManadatoryFields.PrintResponses(props.Patch, path, "PATCH", logger)
		}
	}

	return nonManadatoryFields, nil
}

func (f *NonManadatoryFields) PrintResponses(operation *spec.Operation, path string, method string, logger *logrus.Logger) {
	for statusCode, response := range operation.Responses.StatusCodeResponses {
		if response.ResponseProps.Schema == nil {
			logger.WithFields(logrus.Fields{
				"path":       path,
				"method":     method,
				"statusCode": statusCode,
			}).Debug("PrintResponses: response.ResponseProps.Schema=nil")
			continue
		}

		conditionalProperties := []string{}
		for propName, prop := range response.ResponseProps.Schema.Properties {
			opts := GetOptionalProperties(propName, prop, "\t", []string{})
			conditionalProperties = append(conditionalProperties, opts...)
		}

		endpoint := Endpoint{
			ConditionalProperties: conditionalProperties,
			Method:                method,
			Path:                  path,
			StatusCode:            statusCode,
		}
		f.Endpoints = append(f.Endpoints, endpoint)
	}
}

func GetOptionalProperties(propName string, prop spec.Schema, indent string, path []string) []string {
	// Old working conditional:
	// newPath := path
	// if prop.Type.Contains("array") {
	// 	newPath = append(path, propName+"[*]")
	// } else {
	// 	newPath = append(path, propName)
	// }
	newPath := path
	if prop.Type.Contains("array") {
		newPath = append(newPath, propName+"[*]")
	} else {
		newPath = append(newPath, propName)
	}

	// if (prop.Items == nil || prop.Items.Schema == nil) && len(prop.Properties) <= 0 {
	if (prop.Items == nil || prop.Items.Schema == nil) && !(len(prop.Properties) > 0) {
		return []string{strings.Join(newPath, ".")}
	}

	optionalProperties := []string{}
	if prop.Items != nil && prop.Items.Schema != nil {
		required := prop.Items.Schema.Required
		if len(required) > 0 {
			sort.Strings(required)
		}

		for propName, prop := range prop.Items.Schema.Properties {
			pos := sort.SearchStrings(required, propName)
			found := pos < len(required) && required[pos] == propName
			if !found {
				opts := GetOptionalProperties(propName, prop, indent+"\t", newPath)
				optionalProperties = append(optionalProperties, opts...)
			}
		}
	} else {
		required := prop.Required
		if len(required) > 0 {
			sort.Strings(required)
		}

		for propName, prop := range prop.Properties {
			pos := sort.SearchStrings(required, propName)
			found := pos < len(required) && required[pos] == propName
			if !found {
				opts := GetOptionalProperties(propName, prop, indent+"\t", newPath)
				optionalProperties = append(optionalProperties, opts...)
			}
		}
	}

	return optionalProperties
}
