// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20181130storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_UserAssignedIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity runs a test to see if a specific instance of UserAssignedIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentity instances for property testing - lazily instantiated by
// UserAssignedIdentityGenerator()
var userAssignedIdentityGenerator gopter.Gen

// UserAssignedIdentityGenerator returns a generator of UserAssignedIdentity instances for property testing.
func UserAssignedIdentityGenerator() gopter.Gen {
	if userAssignedIdentityGenerator != nil {
		return userAssignedIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity(generators)
	userAssignedIdentityGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity{}), generators)

	return userAssignedIdentityGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentity_SpecGenerator()
	gens["Status"] = UserAssignedIdentity_STATUSGenerator()
}

func Test_UserAssignedIdentity_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_Spec runs a test to see if a specific instance of UserAssignedIdentity_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentity_Spec instances for property testing - lazily instantiated by
// UserAssignedIdentity_SpecGenerator()
var userAssignedIdentity_SpecGenerator gopter.Gen

// UserAssignedIdentity_SpecGenerator returns a generator of UserAssignedIdentity_Spec instances for property testing.
// We first initialize userAssignedIdentity_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentity_SpecGenerator() gopter.Gen {
	if userAssignedIdentity_SpecGenerator != nil {
		return userAssignedIdentity_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	return userAssignedIdentity_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(UserAssignedIdentityOperatorSpecGenerator())
}

func Test_UserAssignedIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS runs a test to see if a specific instance of UserAssignedIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentity_STATUS instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSGenerator()
var userAssignedIdentity_STATUSGenerator gopter.Gen

// UserAssignedIdentity_STATUSGenerator returns a generator of UserAssignedIdentity_STATUS instances for property testing.
func UserAssignedIdentity_STATUSGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSGenerator != nil {
		return userAssignedIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(generators)
	userAssignedIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS{}), generators)

	return userAssignedIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityOperatorSpec, UserAssignedIdentityOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityOperatorSpec runs a test to see if a specific instance of UserAssignedIdentityOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityOperatorSpec(subject UserAssignedIdentityOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityOperatorSpec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentityOperatorSpec instances for property testing - lazily instantiated by
// UserAssignedIdentityOperatorSpecGenerator()
var userAssignedIdentityOperatorSpecGenerator gopter.Gen

// UserAssignedIdentityOperatorSpecGenerator returns a generator of UserAssignedIdentityOperatorSpec instances for property testing.
func UserAssignedIdentityOperatorSpecGenerator() gopter.Gen {
	if userAssignedIdentityOperatorSpecGenerator != nil {
		return userAssignedIdentityOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec(generators)
	userAssignedIdentityOperatorSpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityOperatorSpec{}), generators)

	return userAssignedIdentityOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec(gens map[string]gopter.Gen) {
	gens["ConfigMaps"] = gen.PtrOf(UserAssignedIdentityOperatorConfigMapsGenerator())
}

func Test_UserAssignedIdentityOperatorConfigMaps_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityOperatorConfigMaps via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps, UserAssignedIdentityOperatorConfigMapsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps runs a test to see if a specific instance of UserAssignedIdentityOperatorConfigMaps round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps(subject UserAssignedIdentityOperatorConfigMaps) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityOperatorConfigMaps
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentityOperatorConfigMaps instances for property testing - lazily instantiated by
// UserAssignedIdentityOperatorConfigMapsGenerator()
var userAssignedIdentityOperatorConfigMapsGenerator gopter.Gen

// UserAssignedIdentityOperatorConfigMapsGenerator returns a generator of UserAssignedIdentityOperatorConfigMaps instances for property testing.
func UserAssignedIdentityOperatorConfigMapsGenerator() gopter.Gen {
	if userAssignedIdentityOperatorConfigMapsGenerator != nil {
		return userAssignedIdentityOperatorConfigMapsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityOperatorConfigMapsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityOperatorConfigMaps{}), generators)

	return userAssignedIdentityOperatorConfigMapsGenerator
}
