// Code in this file generated by gokay: github.com/zencoder/gokay
package gkexample

import (
	"errors"
	"fmt"

	"github.com/zencoder/gokay/gokay"
)

func (s Example) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN MapOfInterfaces field Validations
	errorsMapOfInterfaces := make(gokay.ErrorSlice, 0, 0)
	// NotNil
	if s.MapOfInterfaces == nil {
		errorsMapOfInterfaces = append(errorsMapOfInterfaces, errors.New("is Nil"))
	}
	emMapOfInterfaces := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfInterfaces {
		if err := gokay.Validate(v0); err != nil {
			emMapOfInterfaces[fmt.Sprintf("%v", k0)] = err
		}
	}

	if len(emMapOfInterfaces) > 0 {
		errorsMapOfInterfaces = append(errorsMapOfInterfaces, emMapOfInterfaces)
	}

	if len(errorsMapOfInterfaces) > 0 {
		em["MapOfInterfaces"] = errorsMapOfInterfaces
	}
	// END MapOfInterfaces field Validations

	if len(em) > 0 {
		return em
	} else {
		return nil
	}

}
func (s ExampleStruct) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN HexStringPtr field Validations
	errorsHexStringPtr := make(gokay.ErrorSlice, 0, 0)
	// Length
	if err := gokay.LengthString(16, s.HexStringPtr); err != nil {
		errorsHexStringPtr = append(errorsHexStringPtr, err)
	}

	// NotNil
	if s.HexStringPtr == nil {
		errorsHexStringPtr = append(errorsHexStringPtr, errors.New("is Nil"))
	}
	// Hex
	if err := gokay.IsHex(s.HexStringPtr); err != nil {
		errorsHexStringPtr = append(errorsHexStringPtr, err)
	}

	if len(errorsHexStringPtr) > 0 {
		em["HexStringPtr"] = errorsHexStringPtr
	}
	// END HexStringPtr field Validations

	// BEGIN HexString field Validations
	errorsHexString := make(gokay.ErrorSlice, 0, 0)
	// Length
	if err := gokay.LengthString(12, &s.HexString); err != nil {
		errorsHexString = append(errorsHexString, err)
	}

	// Hex
	if err := gokay.IsHex(&s.HexString); err != nil {
		errorsHexString = append(errorsHexString, err)
	}

	if len(errorsHexString) > 0 {
		em["HexString"] = errorsHexString
	}
	// END HexString field Validations

	// BEGIN BCP47StringPtr field Validations
	errorsBCP47StringPtr := make(gokay.ErrorSlice, 0, 0)
	// NotNil
	if s.BCP47StringPtr == nil {
		errorsBCP47StringPtr = append(errorsBCP47StringPtr, errors.New("is Nil"))
	}
	// BCP47
	if err := gokay.IsBCP47(s.BCP47StringPtr); err != nil {
		errorsBCP47StringPtr = append(errorsBCP47StringPtr, err)
	}

	if len(errorsBCP47StringPtr) > 0 {
		em["BCP47StringPtr"] = errorsBCP47StringPtr
	}
	// END BCP47StringPtr field Validations

	// BEGIN BCP47String field Validations
	errorsBCP47String := make(gokay.ErrorSlice, 0, 0)
	// BCP47
	if err := gokay.IsBCP47(&s.BCP47String); err != nil {
		errorsBCP47String = append(errorsBCP47String, err)
	}

	if len(errorsBCP47String) > 0 {
		em["BCP47String"] = errorsBCP47String
	}
	// END BCP47String field Validations

	// BEGIN CanBeNilWithConstraints field Validations
	errorsCanBeNilWithConstraints := make(gokay.ErrorSlice, 0, 0)
	// Length
	if err := gokay.LengthString(12, s.CanBeNilWithConstraints); err != nil {
		errorsCanBeNilWithConstraints = append(errorsCanBeNilWithConstraints, err)
	}

	if len(errorsCanBeNilWithConstraints) > 0 {
		em["CanBeNilWithConstraints"] = errorsCanBeNilWithConstraints
	}
	// END CanBeNilWithConstraints field Validations

	if len(em) > 0 {
		return em
	} else {
		return nil
	}

}

// Validate
func (s HasValidateImplicit) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN InvalidStruct field Validations
	errorsInvalidStruct := make(gokay.ErrorSlice, 0, 0)

	if s.InvalidStruct != nil {
		if err := gokay.Validate(s.InvalidStruct); err != nil {
			errorsInvalidStruct = append(errorsInvalidStruct, err)
		}
	}

	if len(errorsInvalidStruct) > 0 {
		em["InvalidStruct"] = errorsInvalidStruct
	}
	// END InvalidStruct field Validations

	// BEGIN ValidStruct field Validations
	errorsValidStruct := make(gokay.ErrorSlice, 0, 0)
	if err := gokay.Validate(s.ValidStruct); err != nil {
		errorsValidStruct = append(errorsValidStruct, err)
	}

	if len(errorsValidStruct) > 0 {
		em["ValidStruct"] = errorsValidStruct
	}
	// END ValidStruct field Validations

	// BEGIN MapOfStruct field Validations
	errorsMapOfStruct := make(gokay.ErrorSlice, 0, 0)
	emMapOfStruct := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfStruct {
		if err := gokay.Validate(v0); err != nil {
			emMapOfStruct[fmt.Sprintf("%v", k0)] = err
		}
	}

	if len(emMapOfStruct) > 0 {
		errorsMapOfStruct = append(errorsMapOfStruct, emMapOfStruct)
	}

	if len(errorsMapOfStruct) > 0 {
		em["MapOfStruct"] = errorsMapOfStruct
	}
	// END MapOfStruct field Validations

	// BEGIN MapOfStructPtrs field Validations
	errorsMapOfStructPtrs := make(gokay.ErrorSlice, 0, 0)
	emMapOfStructPtrs := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfStructPtrs {
		if v0 != nil {
			if err := gokay.Validate(v0); err != nil {
				emMapOfStructPtrs[fmt.Sprintf("%v", k0)] = err
			}
		}
	}

	if len(emMapOfStructPtrs) > 0 {
		errorsMapOfStructPtrs = append(errorsMapOfStructPtrs, emMapOfStructPtrs)
	}

	if len(errorsMapOfStructPtrs) > 0 {
		em["MapOfStructPtrs"] = errorsMapOfStructPtrs
	}
	// END MapOfStructPtrs field Validations

	// BEGIN MapOfMaps field Validations
	errorsMapOfMaps := make(gokay.ErrorSlice, 0, 0)
	emMapOfMaps := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfMaps {
		emv0 := make(gokay.ErrorMap)
		for k1, v1 := range v0 {
			if v1 != nil {
				if err := gokay.Validate(v1); err != nil {
					emv0[fmt.Sprintf("%v", k1)] = err
				}
			}
		}

		if len(emv0) > 0 {
			emMapOfMaps[fmt.Sprintf("%v", k0)] = emv0
		}
	}

	if len(emMapOfMaps) > 0 {
		errorsMapOfMaps = append(errorsMapOfMaps, emMapOfMaps)
	}

	if len(errorsMapOfMaps) > 0 {
		em["MapOfMaps"] = errorsMapOfMaps
	}
	// END MapOfMaps field Validations

	// BEGIN MapMapsOfSlices field Validations
	errorsMapMapsOfSlices := make(gokay.ErrorSlice, 0, 0)
	emMapMapsOfSlices := make(gokay.ErrorMap)
	for k0, v0 := range s.MapMapsOfSlices {
		emv0 := make(gokay.ErrorMap)
		for k1, v1 := range v0 {
			emv1 := make(gokay.ErrorMap)
			for k2, v2 := range v1 {
				if v2 != nil {
					if err := gokay.Validate(v2); err != nil {
						emv1[fmt.Sprintf("%v", k2)] = err
					}
				}
			}

			if len(emv1) > 0 {
				emv0[fmt.Sprintf("%v", k1)] = emv1
			}
		}

		if len(emv0) > 0 {
			emMapMapsOfSlices[fmt.Sprintf("%v", k0)] = emv0
		}
	}

	if len(emMapMapsOfSlices) > 0 {
		errorsMapMapsOfSlices = append(errorsMapMapsOfSlices, emMapMapsOfSlices)
	}

	if len(errorsMapMapsOfSlices) > 0 {
		em["MapMapsOfSlices"] = errorsMapMapsOfSlices
	}
	// END MapMapsOfSlices field Validations

	// BEGIN MapOfInterfaces field Validations
	errorsMapOfInterfaces := make(gokay.ErrorSlice, 0, 0)
	emMapOfInterfaces := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfInterfaces {
		if err := gokay.Validate(v0); err != nil {
			emMapOfInterfaces[fmt.Sprintf("%v", k0)] = err
		}
	}

	if len(emMapOfInterfaces) > 0 {
		errorsMapOfInterfaces = append(errorsMapOfInterfaces, emMapOfInterfaces)
	}

	if len(errorsMapOfInterfaces) > 0 {
		em["MapOfInterfaces"] = errorsMapOfInterfaces
	}
	// END MapOfInterfaces field Validations

	// BEGIN SimpleSlice field Validations
	errorsSimpleSlice := make(gokay.ErrorSlice, 0, 0)
	emSimpleSlice := make(gokay.ErrorMap)
	for k0, v0 := range s.SimpleSlice {
		if v0 != nil {
			if err := gokay.Validate(v0); err != nil {
				emSimpleSlice[fmt.Sprintf("%v", k0)] = err
			}
		}
	}

	if len(emSimpleSlice) > 0 {
		errorsSimpleSlice = append(errorsSimpleSlice, emSimpleSlice)
	}

	if len(errorsSimpleSlice) > 0 {
		em["SimpleSlice"] = errorsSimpleSlice
	}
	// END SimpleSlice field Validations

	// BEGIN SliceOfSlicesOfSlices field Validations
	errorsSliceOfSlicesOfSlices := make(gokay.ErrorSlice, 0, 0)
	emSliceOfSlicesOfSlices := make(gokay.ErrorMap)
	for k0, v0 := range s.SliceOfSlicesOfSlices {
		emv0 := make(gokay.ErrorMap)
		for k1, v1 := range v0 {
			emv1 := make(gokay.ErrorMap)
			for k2, v2 := range v1 {
				if v2 != nil {
					if err := gokay.Validate(v2); err != nil {
						emv1[fmt.Sprintf("%v", k2)] = err
					}
				}
			}

			if len(emv1) > 0 {
				emv0[fmt.Sprintf("%v", k1)] = emv1
			}
		}

		if len(emv0) > 0 {
			emSliceOfSlicesOfSlices[fmt.Sprintf("%v", k0)] = emv0
		}
	}

	if len(emSliceOfSlicesOfSlices) > 0 {
		errorsSliceOfSlicesOfSlices = append(errorsSliceOfSlicesOfSlices, emSliceOfSlicesOfSlices)
	}

	if len(errorsSliceOfSlicesOfSlices) > 0 {
		em["SliceOfSlicesOfSlices"] = errorsSliceOfSlicesOfSlices
	}
	// END SliceOfSlicesOfSlices field Validations

	// BEGIN MapOfSlicesOfMaps field Validations
	errorsMapOfSlicesOfMaps := make(gokay.ErrorSlice, 0, 0)
	emMapOfSlicesOfMaps := make(gokay.ErrorMap)
	for k0, v0 := range s.MapOfSlicesOfMaps {
		emv0 := make(gokay.ErrorMap)
		for k1, v1 := range v0 {
			emv1 := make(gokay.ErrorMap)
			for k2, v2 := range v1 {
				if v2 != nil {
					if err := gokay.Validate(v2); err != nil {
						emv1[fmt.Sprintf("%v", k2)] = err
					}
				}
			}

			if len(emv1) > 0 {
				emv0[fmt.Sprintf("%v", k1)] = emv1
			}
		}

		if len(emv0) > 0 {
			emMapOfSlicesOfMaps[fmt.Sprintf("%v", k0)] = emv0
		}
	}

	if len(emMapOfSlicesOfMaps) > 0 {
		errorsMapOfSlicesOfMaps = append(errorsMapOfSlicesOfMaps, emMapOfSlicesOfMaps)
	}

	if len(errorsMapOfSlicesOfMaps) > 0 {
		em["MapOfSlicesOfMaps"] = errorsMapOfSlicesOfMaps
	}
	// END MapOfSlicesOfMaps field Validations

	if len(em) > 0 {
		return em
	} else {
		return nil
	}

}

// Validate
func (s NotNilTestStruct) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN NotNilMap field Validations
	errorsNotNilMap := make(gokay.ErrorSlice, 0, 0)
	// NotNil
	if s.NotNilMap == nil {
		errorsNotNilMap = append(errorsNotNilMap, errors.New("is Nil"))
	}
	emNotNilMap := make(gokay.ErrorMap)
	for k0, v0 := range s.NotNilMap {
		if err := gokay.Validate(v0); err != nil {
			emNotNilMap[fmt.Sprintf("%v", k0)] = err
		}
	}

	if len(emNotNilMap) > 0 {
		errorsNotNilMap = append(errorsNotNilMap, emNotNilMap)
	}

	if len(errorsNotNilMap) > 0 {
		em["NotNilMap"] = errorsNotNilMap
	}
	// END NotNilMap field Validations

	// BEGIN NotNilSlice field Validations
	errorsNotNilSlice := make(gokay.ErrorSlice, 0, 0)
	// NotNil
	if s.NotNilSlice == nil {
		errorsNotNilSlice = append(errorsNotNilSlice, errors.New("is Nil"))
	}

	if len(errorsNotNilSlice) > 0 {
		em["NotNilSlice"] = errorsNotNilSlice
	}
	// END NotNilSlice field Validations

	if len(em) > 0 {
		return em
	}
	return nil
}
