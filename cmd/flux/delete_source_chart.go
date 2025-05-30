/*
Copyright 2024 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/cobra"

	sourcev1 "github.com/fluxcd/source-controller/api/v1"
)

var deleteSourceChartCmd = &cobra.Command{
	Use:   "chart [name]",
	Short: "Delete a HelmChart source",
	Long:  "The delete source chart command deletes the given HelmChart from the cluster.",
	Example: `  # Delete a HelmChart
  flux delete source chart podinfo`,
	ValidArgsFunction: resourceNamesCompletionFunc(sourcev1.GroupVersion.WithKind(sourcev1.HelmChartKind)),
	RunE: deleteCommand{
		apiType: helmChartType,
		object:  universalAdapter{&sourcev1.HelmChart{}},
	}.run,
}

func init() {
	deleteSourceCmd.AddCommand(deleteSourceChartCmd)
}
