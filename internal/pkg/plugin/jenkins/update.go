package jenkins

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/helm"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// Update updates jenkins with provided options.
func Update(options map[string]interface{}) (map[string]interface{}, error) {
	// Initialize Operator with Operations
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			setDefaultValue(defaultHelmConfig),
			helm.Validate,
			replaceStroageClass,
		},
		ExecuteOperations:   helm.DefaultUpdateOperations,
		TerminateOperations: helm.DefaultTerminateOperations,
		GetStateOperation:   getHelmResourceAndCustomResource,
	}

	// Execute all Operations in Operator
	status, err := operator.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}
	log.Debugf("Return map: %v", status)
	return status, nil
}
