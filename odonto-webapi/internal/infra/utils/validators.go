package utils

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}
	sum := 0
	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}
	return sum
}

// E-mail validation
func IsValidMail(email, message string) error {

	//
	if !emailRegex.MatchString(email) {
		return errors.New(message)
	}
	//
	if len(email) == 0 {
		return errors.New(message)
	}

	// Separate the email body from email the domain
	emailParts := strings.Split(email, "@")
	//Verify if the body or the domain are empty
	if len(emailParts[0]) == 0 || len(emailParts[1]) == 0 {
		return errors.New(message)
	}

	// verify if the domain exists and accept emails
	mx, err := net.LookupMX(emailParts[1])
	if err != nil || len(mx) == 0 {
		return errors.New(message)
	}
	return nil
}

func IsValidCPF(value interface{}, message string) error {
	//
	cpf, ok := value.(string)
	if !ok {
		return errors.New(message)
	}
	//
	if len(cpf) != 11 {
		return errors.New(message)
	}
	if cpf == "00000000000" ||
		cpf == "11111111111" ||
		cpf == "22222222222" ||
		cpf == "33333333333" ||
		cpf == "44444444444" ||
		cpf == "55555555555" ||
		cpf == "66666666666" ||
		cpf == "77777777777" ||
		cpf == "88888888888" ||
		cpf == "99999999999" {
		return errors.New(message)
	}
	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)
	remainsFirst := sum % 11
	d1 := 0
	if remainsFirst >= 2 {
		d1 = 11 - remainsFirst
	}
	secondPart := firstPart + strconv.Itoa(d1)
	dsum := sumDigit(secondPart, cpfSecondDigitTable)
	r2 := dsum % 11
	d2 := 0
	if r2 >= 2 {
		d2 = 11 - r2
	}
	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	if finalPart == cpf {
		return nil
	}
	return errors.New(message)
}
