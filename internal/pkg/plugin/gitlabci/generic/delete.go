package generic

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer/ci"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer/ci/cifile"
)

func Delete(options configmanager.RawOptions) (bool, error) {
	operator := &installer.Operator{
		PreExecuteOperations: installer.PreExecuteOperations{
			ci.SetDefault(ciType),
			validate,
		},
		ExecuteOperations: installer.ExecuteOperations{
			cifile.DeleteCIFiles,
		},
	}

	// Execute all Operations in Operator
	_, err := operator.Execute(options)
	if err != nil {
		return false, err
	}
	return true, nil
}
