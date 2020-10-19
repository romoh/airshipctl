/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kustomize/api/types"
)

// +kubebuilder:object:root=true

// ReplacementTransformer plugin configuration for airship document model
type ReplacementTransformer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Replacements list of source and target field to do a replacement
	Replacements []types.Replacement `json:"replacements,omitempty" yaml:"replacements,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplacementTransformer) DeepCopyInto(out *ReplacementTransformer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Replacements != nil {
		out.Replacements = make([]types.Replacement, len(in.Replacements))
		for i, repl := range in.Replacements {
			out.Replacements[i] = types.Replacement{
				Source: &types.ReplSource{
					ObjRef:   &types.Target{},
					FieldRef: repl.Source.FieldRef,
					Value:    repl.Source.Value,
				},
				Target: &types.ReplTarget{
					ObjRef:    &types.Selector{},
					FieldRefs: repl.Target.FieldRefs,
				},
			}
			*(out.Replacements[i].Source.ObjRef) = *(in.Replacements[i].Source.ObjRef)
			*(out.Replacements[i].Target.ObjRef) = *(in.Replacements[i].Target.ObjRef)
		}
	}
}
