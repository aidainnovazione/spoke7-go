// Copyright (c) 2023 Vladislav Fursov (GhostIAm)
// This code is licensed under MIT license (see LICENSE for details)
// https://gist.github.com/ghostiam/48acd974f2044e25ba43f090316e6f2d

package utils

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var _ runtime.Marshaler = (*MultipartFormPb)(nil)

type MultipartFormPb struct {
	runtime.Marshaler
}

func GWMultipartForm(marshaler runtime.Marshaler) runtime.ServeMuxOption {
	return runtime.WithMarshalerOption("multipart/form-data", &MultipartFormPb{
		Marshaler: marshaler,
	})
}

func (j *MultipartFormPb) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v any) error {
		msg, ok := v.(proto.Message)
		if !ok {
			return errors.New("not proto message") //nolint:goerr113
		}

		br := bufio.NewReaderSize(r, 1024)
		pb, err := br.Peek(100)
		if err != nil {
			return fmt.Errorf("peek boundary: %w", err)
		}

		if len(pb) < 2 {
			return errors.New("boundary len < 2") //nolint:goerr113
		}

		boundary := bytes.TrimSpace(bytes.Split(pb, []byte("\n"))[0])[2:]

		values := make(url.Values)

		mr := multipart.NewReader(br, string(boundary))
		for {
			var p *multipart.Part
			p, err = mr.NextPart()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return fmt.Errorf("read next part: %w", err)
			}

			formName := p.FormName()

			var fieldDescriptor protoreflect.FieldDescriptor
			fieldDescriptor, err = fieldDescriptorByPath(msg, formName)
			if err != nil {
				return err
			}

			var data []byte
			data, err = io.ReadAll(p)
			if err != nil {
				return fmt.Errorf("read part body: %w", err)
			}

			if fieldDescriptor.Kind() == protoreflect.BytesKind {
				values.Set(formName, base64.StdEncoding.EncodeToString(data))
				continue
			}

			if p.FileName() != "" && fieldDescriptor.Kind() == protoreflect.MessageKind {
				/*
					in proto file:
						message Media {
							string filename = 1;
							string content_type = 2;
							bytes content = 3;
						}
				*/
				values.Set(formName+".filename", p.FileName())
				values.Set(formName+".content_type", p.Header.Get("Content-Type"))
				values.Set(formName+".content", base64.StdEncoding.EncodeToString(data))

				continue
			}

			values.Set(formName, string(data))
		}

		err = runtime.PopulateQueryParameters(msg, values, &utilities.DoubleArray{})
		if err != nil {
			return fmt.Errorf("populate query params: %w", err)
		}

		return nil
	})
}

func (j *MultipartFormPb) Unmarshal(data []byte, v any) error {
	return j.NewDecoder(bytes.NewReader(data)).Decode(v) //nolint: wrapcheck
}

func fieldDescriptorByPath(msg proto.Message, path string) (protoreflect.FieldDescriptor, error) {
	msgValue := msg.ProtoReflect()
	fieldPath := strings.Split(path, ".")
	var fieldDescriptor protoreflect.FieldDescriptor
	for i, fieldName := range fieldPath {
		msgFields := msgValue.Descriptor().Fields()

		fieldDescriptor = msgFields.ByName(protoreflect.Name(fieldName))
		if fieldDescriptor == nil {
			fieldDescriptor = msgFields.ByJSONName(fieldName)
		}

		// If this is the last element, we're done
		if i == len(fieldPath)-1 {
			break
		}

		// Only singular message fields are allowed
		if fieldDescriptor == nil || fieldDescriptor.Message() == nil || fieldDescriptor.Cardinality() == protoreflect.Repeated {
			//nolint:goerr113
			return nil, fmt.Errorf("invalid path: %q is not a message", path)
		}

		// Get the nested message
		msgValue = msgValue.Mutable(fieldDescriptor).Message()
	}

	if fieldDescriptor == nil {
		//nolint:goerr113
		return nil, fmt.Errorf("invalid path: %q is not a valid field", path)
	}

	return fieldDescriptor, nil
}
