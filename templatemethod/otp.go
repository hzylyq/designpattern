package templatemethod

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	message(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp iOtp
}

func (o *otp) genAndSendOtp(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.message(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
