package requirements

import "go.arcalot.io/log"

func BasicRequirements(filenames []string, logger log.Logger) (bool, error) {
	meets_reqs := true
	output := ""

	if present, err := HasFilename(filenames, "README.md"); err != nil {
		return false, err
	} else if !present {
		output = "Missing README.md\n"
		logger.Infof(output)
		meets_reqs = false
	}

	if present, err := HasFilename(filenames, "Dockerfile"); err != nil {
		return false, err
	} else if !present {
		output = "Missing Dockerfile\n"
		logger.Infof(output)
		meets_reqs = false
	}

	if present, err := HasFilename(filenames, "(?i).*test.*"); err != nil {
		return false, err
	} else if !present {
		// match case-insensitive 'test'?
		output = "Missing a test file\n"
		logger.Infof(output)
		meets_reqs = false
	}

	return meets_reqs, nil
}
