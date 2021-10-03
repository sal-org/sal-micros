package constant

// PaymentURLs - payment complete endpoints - if any of them is success, then payment is captured and order is created
var PaymentURLs = []string{
	"/client/counsellor/paymentcomplete", // client counsellor appointment
	"/client/listener/paymentcomplete",   // client listener appointment
	"/client/therapist/paymentcomplete",  // client therapist appointment
	"/client/event/paymentcomplete",      // client event booking

	"/counsellor/event/paymentcomplete",       // counsellor event booking
	"/counsellor/event/block/paymentcomplete", // counsellor event blocking

	"/listener/event/paymentcomplete", // listener event booking

	"/therapist/event/paymentcomplete",       // therapist event booking
	"/therapist/event/block/paymentcomplete", // therapist event blocking
}
