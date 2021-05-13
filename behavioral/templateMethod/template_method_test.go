package templatemethod

import (
	"testing"
)

func TestPattern(t *testing.T) {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
