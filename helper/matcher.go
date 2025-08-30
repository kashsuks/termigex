package helper

import "regexp"

func FindMathes(regex, testStr string) ([]string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	matches := pattern.FindAllString(testStr, -1)
	return matches, nil
}