package templatemethod

import "testing"

func Test_genAndSendOTP(t *testing.T) {
	smsOtp := &sms{}
	o := otp{
		iOtp: smsOtp,
	}

	o.genAndSendOtp(4)
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOtp(4)
}
