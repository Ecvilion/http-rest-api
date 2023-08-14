/*Когда читаем юзера из БД мы не записываем ему полу password => из-за этого валидация на него не будет работать, т.к. пароль
должен быть не нулевым*/

package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
