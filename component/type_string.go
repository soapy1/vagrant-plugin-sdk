// Code generated by "stringer -type=Type -linecomment"; DO NOT EDIT.

package component

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidType-0]
	_ = x[ProviderType-1]
	_ = x[ProvisionerType-2]
	_ = x[maxType-3]
}

const _Type_name = "InvalidProviderProvisionermaxType"

var _Type_index = [...]uint8{0, 7, 15, 26, 33}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
