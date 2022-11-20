package arguments

import (
	"fmt"
	"os"
	"strings"
)

// parsedArgs holds the parsedArgs arguments and parameters.
type parsedArgs struct {
	arguments  map[string]any
	parameters map[string]*string
}

// Grab grabs the arguments and parameters from the command line.
func Grab() (result *parsedArgs) {
	result = &parsedArgs{}

	var parameterIndex int
	result.arguments, parameterIndex = fetchArguments()
	result.parameters = fetchParameters(parameterIndex)
	return
}

// fetchArguments fetches the arguments from the command line.
func fetchArguments() (result map[string]any, parameterIndex int) {
	for index, value := range os.Args[1:] {
		if strings.HasPrefix(value, "-") {
			parameterIndex = index + 1
			break
		}
		result[value] = nil
	}
	return
}

// fetchParameters fetches the parameters from the command line.
func fetchParameters(startIndex int) (result map[string]*string) {
	result = make(map[string]*string)
	for index := startIndex; index < len(os.Args); index++ {
		isParameter := strings.HasPrefix(os.Args[index], "-")
		if !isParameter {
			panic(fmt.Sprintf("Incorrect parameter syntax near '%s'", os.Args[index]))
		}

		moreParametersAfter := index+1 < len(os.Args)
		nextIsParameter := moreParametersAfter &&
			(strings.HasPrefix(os.Args[index+1], "-") && !strings.ContainsRune(os.Args[index+1], ' '))
		isNamedParameter := strings.HasPrefix(os.Args[index], "--")
		parameter := strings.TrimLeft(os.Args[index], "-")

		if !moreParametersAfter || nextIsParameter {
			if isNamedParameter {
				result[parameter] = nil
				continue
			}

			for _, character := range parameter {
				result[string(character)] = nil
			}
			continue
		}

		value := &os.Args[index+1]
		if isNamedParameter {
			result[parameter] = value
		} else {
			for _, character := range parameter {
				result[string(character)] = value
			}
		}

		// The next index is a value, so skip it.
		index++
	}

	return
}

// HasParameter checks if the parameter exists.
func (p *parsedArgs) HasParameter(parameter string) bool {
	_, ok := p.parameters[parameter]
	return ok
}

// GetParameterValue gets the value of the parameter.
func (p *parsedArgs) GetParameterValue(parameter string) (string, error) {
	value, ok := p.parameters[parameter]
	if !ok {
		return "", fmt.Errorf("parameter '%s' not found", parameter)
	}
	if value == nil {
		return "", fmt.Errorf("parameter '%s' has no value", parameter)
	}
	return *value, nil
}

// HasArgument checks if the argument exists.
func (p *parsedArgs) HasArgument(argument string) (ok bool) {
	_, ok = p.arguments[argument]
	return
}
