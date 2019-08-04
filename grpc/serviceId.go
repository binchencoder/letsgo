package grpc

import (
	"fmt"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	mdKeyServiceId = "mdkey_serviceid"
)

// WithServiceId returns a copy of outgoing context, with service ID as part of
// the metadata.
// If the outgoing context already has metadata, the metadata will be joined.
func WithServiceId(outgoing context.Context, serviceId int) context.Context {
	md := metadata.Pairs(mdKeyServiceId, strconv.Itoa(serviceId))

	// Pass existing metadata down.
	if existingMd, ok := metadata.FromOutgoingContext(outgoing); ok {
		md = metadata.Join(existingMd, md)
	}

	return metadata.NewOutgoingContext(outgoing, md)
}

// GetServiceId gets service ID from incoming context metadata.
func GetServiceId(incoming context.Context) (int, error) {
	val, err := getLastVal(incoming, mdKeyServiceId)
	if err != nil {
		return -1, err
	}

	id, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func getLastVal(incoming context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(incoming)
	if !ok {
		return "", fmt.Errorf("Failed to get metadata from incoming context")
	}

	vals, found := md[key]
	if !found {
		return "", fmt.Errorf("Failed to find key %q from metadata", key)
	}
	if len(vals) == 0 {
		return "", fmt.Errorf("Found empty list of values for key %q in metadata", key)
	}

	// Returns the last value, which is the latest.
	return vals[len(vals)-1], nil
}
