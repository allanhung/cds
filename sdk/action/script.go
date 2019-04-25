package action

import (
	"github.com/ovh/cds/sdk"
)

// Script action definition.
var Script = sdk.Action{
	Name: sdk.ScriptAction,
	Description: `CDS Builtin Action.
Execute a script, written in script attribute.`,
	Parameters: []sdk.Parameter{
		{
			Name: "script",
			Description: `Content of your script.
You can put #!/bin/bash, or #!/bin/perl at first line.
Make sure that the binary used is in
the pre-requisites of action`,
			Type: sdk.TextParameter,
		},
	},
}
