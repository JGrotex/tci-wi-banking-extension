/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package ibanvalidation

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivIBAN   = "IBAN"
	ovResult = "result"
)

var activityLog = logger.GetLogger("banking-activity-IBANvalidation")

type IBANvalidationActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &IBANvalidationActivity{metadata: metadata}
}

func (a *IBANvalidationActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *IBANvalidationActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing IBAN validation activity")
	//Read Inputs
	if context.GetInput(ivIBAN) == nil {
		// IBAN string is not configured
		// return error to the engine
		return false, activity.NewError("IBAN string is not configured", "IBAN-4001", nil)
	}
	IBANstr := context.GetInput(ivIBAN).(string)
	//activityLog.Info("IBAN: " + IBANstr)

	// execute validation - Start
	var OK, well_formated, ibanerr = IsCorrectIban(IBANstr, true) // passed: IBAN string , debug true/false
	activityLog.Info("IBAN well_formated: " + well_formated)
	if ibanerr != nil {
		activityLog.Info("IBAN format error")
		//	return false, activity.NewError("IBAN Error "+ibanerr.Error(), "IBAN-4002", nil)
	}

	if OK {
		//Use separator in concatenation
		context.SetOutput(ovResult, "valid")
	} else {
		//No separator in concatenation
		context.SetOutput(ovResult, "invalid")
	}
	//}
	return true, nil
}