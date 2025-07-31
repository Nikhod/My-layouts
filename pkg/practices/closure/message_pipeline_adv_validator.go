package closure

func ValidatorPractice() {
	messages := getMessages()
	messages = AreAllValidatorIsTrue(messages, isPositiveId, isLongEnough, isSMS)
}

type Validator func(object message) bool

func isLongEnough(object message) bool {
	return len(object.body) > 5
}
func isPositiveId(object message) bool {
	return object.id > 0
}
func isSMS(object message) bool {
	return object.msgType == "sms"
}

func AreAllValidatorIsTrue(objects []message, validators ...Validator) []message {
	var validated []message

	for _, obj := range objects {
		var passed bool = true
		for _, checker := range validators {
			if !(checker(obj)) {
				passed = false
			}
		}
		if passed {
			validated = append(validated, obj)
		}

	}

	return validated
}

// another way of implementing - correct too
func NikValid(msg []message, vals ...func(msg message) bool) {}
