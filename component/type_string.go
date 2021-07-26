// Code generated by "stringer -type=Type -linecomment ./component"; DO NOT EDIT.

package component

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidType-0]
	_ = x[CommandType-1]
	_ = x[CommunicatorType-2]
	_ = x[GuestType-3]
	_ = x[HostType-4]
	_ = x[ProviderType-5]
	_ = x[ProvisionerType-6]
	_ = x[SyncedFolderType-7]
	_ = x[AuthenticatorType-8]
	_ = x[LogPlatformType-9]
	_ = x[LogViewerType-10]
	_ = x[MapperType-11]
	_ = x[ConfigType-12]
	_ = x[PluginInfoType-13]
	_ = x[maxType-14]
}

const _Type_name = "InvalidCommandCommunicatorGuestHostProviderProvisionerSynced FolderAuthenticatorLogPlatformLogViewerMapperConfigPluginInfomaxType"

var _Type_index = [...]uint8{0, 7, 14, 26, 31, 35, 43, 54, 67, 80, 91, 100, 106, 112, 122, 129}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
