package grpc

import "github.com/FrancescoIlario/beershop"

func protoValidationError(vr beershop.ValidationResult) []*ValidationEntry {
	ee := vr.Errors()
	vee := make([]*ValidationEntry, len(ee))
	count := 0
	for k, v := range ee {
		vee[count] = &ValidationEntry{
			Key:   k,
			Value: v,
		}
		count++
	}
	return vee
}
