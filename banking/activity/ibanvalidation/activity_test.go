/*
 * Copyright © 2018. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package ibanvalidation

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())
	//setup attrs
	tc.SetInput("IBAN", "DE89 3704 0044 0532 0130 00")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("result")
	assert.Equal(t, result, "valid")

	t.Log(result)
}

func TestInvalidEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())
	//setup attrs
	tc.SetInput("IBAN", "DE89 3704 0044 0532 0130 001")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("result")
	assert.Equal(t, result, "invalid")

	t.Log(result)
}
