package protomappers

import (
	"context"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	"github.com/hashicorp/vagrant-plugin-sdk/config"
	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/datadir"
	"github.com/hashicorp/vagrant-plugin-sdk/helper/path"
	"github.com/hashicorp/vagrant-plugin-sdk/helper/types"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/cacher"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/dynamic"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/pluginclient"
	plugincomponent "github.com/hashicorp/vagrant-plugin-sdk/internal/plugin"
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/internal/plugin/core"
	pluginterminal "github.com/hashicorp/vagrant-plugin-sdk/internal/plugin/terminal"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/pluginargs"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

var WellKnownTypes = []interface{}{
	Boolean,
	BooleanPtr,
	BooleanProto,
	BooleanPtrProto,
	Bytes,
	BytesProto,
	Double,
	DoublePtr,
	DoubleProto,
	DoublePtrProto,
	Float,
	FloatPtr,
	FloatProto,
	FloatPtrProto,
	Int32,
	Int32Ptr,
	Int32Proto,
	Int32PtrProto,
	Int64,
	Int64Ptr,
	Int64Proto,
	Int64PtrProto,
	String,
	StringPtr,
	StringProto,
	StringPtrProto,
	Timestamp,
	TimestampPtr,
	TimestampProto,
	TimestampPtrProto,
	UInt32,
	UInt32Ptr,
	UInt32Proto,
	UInt32PtrProto,
	UInt64,
	UInt64Ptr,
	UInt64Proto,
	UInt64PtrProto,
	ValueToBool,
	ValueToList,
	ValueToNull,
	ValueToNumber,
	ValueToString,
	ValueToStruct,
	ValueToString,
}

// All is the list of all mappers as raw function pointers.
var All = []interface{}{
	Array,
	ArrayProto,
	Symbol,
	SymbolProto,
	ValueToString,
	Basis,
	BasisProto,
	Box,
	BoxProto,
	BoxCollection,
	BoxCollectionProto,
	BoxMetadata,
	BoxMetadataProto,
	BoxResponse,
	BoxResponseProto,
	Host,
	HostProto,
	Guest,
	GuestProto,
	Class,
	ClassProto,
	Command,
	CommandProto,
	CommandInfo,
	CommandInfoProto,
	CommandParams,
	CommandInfoFromResponse,
	CommunicatorCommand,
	CommunicatorCommandProto,
	Communicator,
	CommunicatorProto,
	ConfigData,
	ConfigDataProto,
	ConfigMerge,
	ConfigMergeProto,
	ConfigFinalize,
	ConfigFinalizeProto,
	DatadirBasis,
	DatadirBasisProto,
	DatadirProject,
	DatadirProjectProto,
	DatadirTarget,
	DatadirTargetProto,
	DatadirComponent,
	DatadirComponentProto,
	Direct,
	DirectProto,
	Flags,
	FlagsProto,
	Folders,
	FoldersProto,
	FolderToVagrantfileSyncedFolder,
	Hash,
	HashProto,
	JobInfo,
	JobInfoProto,
	Logger,
	LoggerProto,
	MachineProject,
	MachineState,
	MachineStateProto,
	SshInfo,
	SshInfoProto,
	CorePluginManager,
	CorePluginManagerProto,
	MapKeyInterface,
	MapKeyString,
	// TODO: This has been disabled since we want to
	//       have map values converted into Args.Hash
	//       messages. Delete this and function once
	//       confirmed we aren't relying on it for anything
	// MapToProto,
	Metadata,
	MetadataProto,
	MetadataSet,
	MetadataSetProto,
	NamedCapability,
	NamedCapabilityProto,
	Options,
	OptionsProto,
	Path,
	PathProto,
	Plugin,
	PluginProto,
	Plugins,
	PluginsProto,
	PluginManager,
	PluginManagerProto,
	Project,
	ProjectProto,
	ProtoToMap,
	Provider,
	ProviderProto,
	Provisioner,
	ProvisionerProto,
	Push,
	PushProto,
	Range,
	RangeProto,
	RawRubyValue,
	RawRubyValueProto,
	Seeds,
	SeedsProto,
	State,
	StateProto,
	StateBag,
	StateBagProto,
	SyncedFolder,
	SyncedFolderProto,
	Target,
	TargetProto,
	TargetIndex,
	TargetIndexProto,
	TargetMachine,
	TargetMachineProto,
	TargetProject,
	TargetToMachine,
	TerminalUI,
	TerminalUIProto,
	Vagrantfile,
	VagrantfileProto,
	VagrantfileSyncedFolderToFolder,
}

var AllFns []*argmapper.Func
var WellKnownFns []*argmapper.Func

func MapKeyInterface(
	input map[string]interface{},
) map[interface{}]interface{} {
	r := make(map[interface{}]interface{}, len(input))
	for k, v := range input {
		r[k] = v
	}
	return r
}

func MapKeyString(
	input map[interface{}]interface{},
) (map[string]interface{}, error) {
	r := make(map[string]interface{}, len(input))
	for ik, v := range input {
		k, ok := ik.(string)
		if !ok {
			return nil, fmt.Errorf("key value is not a string (%#v)", ik)
		}
		r[k] = v
	}

	return r, nil
}

// Known type mappers

func Boolean(
	input *wrapperspb.BoolValue,
) bool {
	return input.Value
}

func BooleanPtr(
	input *wrapperspb.BoolValue,
) *bool {
	return &input.Value
}

func BooleanProto(
	input bool,
) *wrapperspb.BoolValue {
	return &wrapperspb.BoolValue{
		Value: input,
	}
}

func BooleanPtrProto(
	input *bool,
) *wrapperspb.BoolValue {
	if input == nil {
		return nil
	}
	return &wrapperspb.BoolValue{
		Value: *input,
	}
}

func Bytes(
	input *wrapperspb.BytesValue,
) []byte {
	return input.Value
}

func BytesProto(
	input []byte,
) *wrapperspb.BytesValue {
	return &wrapperspb.BytesValue{
		Value: input,
	}
}

func Double(
	input *wrapperspb.DoubleValue,
) float64 {
	return input.Value
}

func DoublePtr(
	input *wrapperspb.DoubleValue,
) *float64 {
	return &input.Value
}

func DoubleProto(
	input float64,
) *wrapperspb.DoubleValue {
	return &wrapperspb.DoubleValue{
		Value: input,
	}
}

func DoublePtrProto(
	input *float64,
) *wrapperspb.DoubleValue {
	if input == nil {
		return nil
	}
	return &wrapperspb.DoubleValue{
		Value: *input,
	}
}

func Float(
	input *wrapperspb.FloatValue,
) float32 {
	return input.Value
}

func FloatPtr(
	input *wrapperspb.FloatValue,
) *float32 {
	return &input.Value
}

func FloatProto(
	input float32,
) *wrapperspb.FloatValue {
	return &wrapperspb.FloatValue{
		Value: input,
	}
}

func FloatPtrProto(
	input *float32,
) *wrapperspb.FloatValue {
	if input == nil {
		return nil
	}
	return &wrapperspb.FloatValue{
		Value: *input,
	}
}

func Int32(
	input *wrapperspb.Int32Value,
) int32 {
	return input.Value
}

func Int32Ptr(
	input *wrapperspb.Int32Value,
) *int32 {
	return &input.Value
}

func Int32Proto(
	input int32,
) *wrapperspb.Int32Value {
	return &wrapperspb.Int32Value{
		Value: input,
	}
}

func Int32PtrProto(
	input *int32,
) *wrapperspb.Int32Value {
	if input == nil {
		return nil
	}
	return &wrapperspb.Int32Value{
		Value: *input,
	}
}

func Int64(
	input *wrapperspb.Int64Value,
) int64 {
	return input.Value
}

func Int64Ptr(
	input *wrapperspb.Int64Value,
) *int64 {
	return &input.Value
}

func Int64Proto(
	input int64,
) *wrapperspb.Int64Value {
	return &wrapperspb.Int64Value{
		Value: input,
	}
}

func Int64PtrProto(
	input *int64,
) *wrapperspb.Int64Value {
	if input == nil {
		return nil
	}
	return &wrapperspb.Int64Value{
		Value: *input,
	}
}

func String(
	input *wrapperspb.StringValue,
) string {
	return input.Value
}

func StringPtr(
	input *wrapperspb.StringValue,
) *string {
	return &input.Value
}

func StringProto(
	input string,
) *wrapperspb.StringValue {
	return &wrapperspb.StringValue{
		Value: input,
	}
}

func StringPtrProto(
	input *string,
) *wrapperspb.StringValue {
	if input == nil {
		return nil
	}
	return &wrapperspb.StringValue{
		Value: *input,
	}
}

func Timestamp(
	input *timestamppb.Timestamp,
) time.Time {
	return input.AsTime()
}

func TimestampPtr(
	input *timestamppb.Timestamp,
) *time.Time {
	v := input.AsTime()
	return &v
}

func TimestampProto(
	input time.Time,
) *timestamppb.Timestamp {
	return timestamppb.New(input)
}

func TimestampPtrProto(
	input *time.Time,
) *timestamppb.Timestamp {
	if input == nil {
		return nil
	}
	return timestamppb.New(*input)
}

func UInt32(
	input *wrapperspb.UInt32Value,
) uint32 {
	return input.Value
}

func UInt32Ptr(
	input *wrapperspb.UInt32Value,
) *uint32 {
	return &input.Value
}

func UInt32Proto(
	input uint32,
) *wrapperspb.UInt32Value {
	return &wrapperspb.UInt32Value{
		Value: input,
	}
}

func UInt32PtrProto(
	input *uint32,
) *wrapperspb.UInt32Value {
	if input == nil {
		return nil
	}
	return &wrapperspb.UInt32Value{
		Value: *input,
	}
}

func UInt64(
	input *wrapperspb.UInt64Value,
) uint64 {
	return input.Value
}

func UInt64Ptr(
	input *wrapperspb.UInt64Value,
) *uint64 {
	return &input.Value
}

func UInt64Proto(
	input uint64,
) *wrapperspb.UInt64Value {
	return &wrapperspb.UInt64Value{
		Value: input,
	}
}

func UInt64PtrProto(
	input *uint64,
) *wrapperspb.UInt64Value {
	if input == nil {
		return nil
	}
	return &wrapperspb.UInt64Value{
		Value: *input,
	}
}

func ValueToBool(
	input *structpb.Value,
) (bool, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_BoolValue)(nil)) {
		return false, fmt.Errorf("value is not bool kind")
	}

	return input.GetBoolValue(), nil
}

func ValueToList(
	input *structpb.Value,
) ([]*structpb.Value, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_ListValue)(nil)) {
		return nil, fmt.Errorf("value is not list kind")
	}

	return input.GetListValue().Values, nil
}

func ValueToNull(
	input *structpb.Value,
) (interface{}, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_NullValue)(nil)) {
		return nil, fmt.Errorf("value is not null kind")
	}

	return nil, nil
}

func ValueToNumber(
	input *structpb.Value,
) (float64, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_NumberValue)(nil)) {
		return 0, fmt.Errorf("value is not number kind")
	}

	return input.GetNumberValue(), nil
}

func ValueToString(
	input *structpb.Value,
) (string, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_StringValue)(nil)) {
		return "", fmt.Errorf("value is not string kind")
	}
	return input.GetStringValue(), nil
}

func ValueToStruct(
	input *structpb.Value,
) (map[string]*structpb.Value, error) {
	if reflect.TypeOf(input.Kind) != reflect.TypeOf((*structpb.Value_StructValue)(nil)) {
		return nil, fmt.Errorf("value is not struct kind")
	}

	return input.GetStructValue().GetFields(), nil
}

// Custom mappers

// NOTE: This does not convert the proto back to the Seeds fully. It
// will only convert the base type, but the contents will remain as
// any values to prevent large numbers of grpc service/client setups
func Seeds(
	input *vagrant_plugin_sdk.Args_Seeds,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*core.Seeds, error) {
	result := core.NewSeeds()
	t := make([]interface{}, len(input.Typed))
	for i := 0; i < len(input.Typed); i++ {
		t[i] = input.Typed[i]
	}
	result.Typed = t

	for k := range input.Named {
		result.Named[k] = input.Named[k]
	}

	return result, nil
}

func SeedsProto(
	input *core.Seeds,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_Seeds, error) {
	result := &vagrant_plugin_sdk.Args_Seeds{
		Named: make(map[string]*anypb.Any),
		Typed: make([]*anypb.Any, len(input.Typed)),
	}

	for i := 0; i < len(input.Typed); i++ {
		a, ok := input.Typed[i].(*anypb.Any)
		if !ok {
			return SeedsProtoFull(input, log, internal, ctx)
		}
		result.Typed[i] = a
	}

	for k := range input.Named {
		a, ok := input.Named[k].(*anypb.Any)
		if !ok {
			return SeedsProtoFull(input, log, internal, ctx)
		}
		result.Named[k] = a
	}

	return result, nil
}

func SeedsProtoFull(
	input *core.Seeds,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_Seeds, error) {
	result := &vagrant_plugin_sdk.Args_Seeds{
		Typed: make([]*anypb.Any, 0, len(input.Typed)),
		Named: make(map[string]*anypb.Any, len(input.Named)),
	}

	d := &component.Direct{Arguments: []interface{}{}}
	for _, v := range input.Typed {
		if a, ok := v.(*anypb.Any); ok {
			result.Typed = append(result.Typed, a)
			continue
		}
		d.Arguments = append(d.Arguments, v)
	}

	if len(d.Arguments) > 0 {
		t, err := DirectProto(
			d, log, internal, ctx,
		)

		if err != nil {
			return nil, err
		}

		for _, v := range t.Arguments {
			result.Typed = append(result.Typed, v)
		}
	}

	d = &component.Direct{Arguments: []interface{}{}}
	names := []string{}

	for k, v := range input.Named {
		if a, ok := v.(*anypb.Any); ok {
			result.Named[k] = a
			continue
		}
		names = append(names, k)
		d.Arguments = append(d.Arguments, v)
	}

	if len(d.Arguments) > 0 {
		t, err := DirectProto(
			d, log, internal, ctx,
		)

		if err != nil {
			return nil, err
		}

		for i := 0; i < len(names); i++ {
			result.Named[names[i]] = t.Arguments[i]
		}
	}

	return result, nil
}

func TargetToMachine(
	input core.Target,
) (core.Machine, error) {
	m, err := input.Specialize((*core.Machine)(nil))
	if err != nil {
		return nil, err
	}
	return m.(core.Machine), nil
}

func Range(
	input *vagrant_plugin_sdk.Args_Range,
) (types.Range, error) {
	return types.NewRange(input.Start, input.End)
}

func RangeProto(
	input types.Range,
) *vagrant_plugin_sdk.Args_Range {
	return &vagrant_plugin_sdk.Args_Range{
		Start: input.Initial(),
		End:   input.Final(),
	}
}

func Symbol(
	input *vagrant_plugin_sdk.Args_Symbol,
) types.Symbol {
	return types.Symbol(input.Str)
}

func SymbolProto(
	input types.Symbol,
) *vagrant_plugin_sdk.Args_Symbol {
	return &vagrant_plugin_sdk.Args_Symbol{
		Str: string(input),
	}
}

func Array(
	input *vagrant_plugin_sdk.Args_Array,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) ([]interface{}, error) {
	result, err := Direct(
		&vagrant_plugin_sdk.Args_Direct{
			Arguments: input.List,
		},
		log,
		internal,
		ctx,
	)

	if err != nil {
		return nil, err
	}

	return result.Arguments, nil
}

func ArrayProto(
	input []interface{},
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_Array, error) {
	result, err := DirectProto(
		&component.Direct{
			Arguments: input,
		},
		log,
		internal,
		ctx,
	)

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Args_Array{
		List: result.Arguments,
	}, nil
}

func Class(
	input *vagrant_plugin_sdk.Args_Class,
) types.Class {
	return types.Class(input.Name)
}

func ClassProto(
	input types.Class,
) *vagrant_plugin_sdk.Args_Class {
	return &vagrant_plugin_sdk.Args_Class{
		Name: string(input),
	}
}

func RawRubyValue(
	input *vagrant_plugin_sdk.Config_RawRubyValue,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*types.RawRubyValue, error) {
	v, err := Hash(input.Data, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{}, len(v))
	for key, val := range v {
		skey, ok := key.(string)
		if !ok {
			symkey, ok := key.(types.Symbol)
			if !ok {
				return nil, fmt.Errorf("invalid root key type %T", key)
			}
			skey = string(symkey)
		}
		data[skey] = val
	}

	result := &types.RawRubyValue{
		Data: data,
	}
	if input.Source != nil {
		result.Source = Class(input.Source)
	}

	return result, nil
}

func RawRubyValueProto(
	input *types.RawRubyValue,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Config_RawRubyValue, error) {
	i := make(map[interface{}]interface{}, len(input.Data))
	for key, val := range input.Data {
		i[key] = val
	}
	iproto, err := HashProto(i, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	return &vagrant_plugin_sdk.Config_RawRubyValue{
		Data:   iproto,
		Source: ClassProto(input.Source),
	}, nil
}

func ConfigData(
	input *vagrant_plugin_sdk.Args_ConfigData,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*component.ConfigData, error) {
	v, err := Hash(input.Data, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{}, len(v))
	for key, val := range v {
		skey, ok := key.(string)
		if !ok {
			symkey, ok := key.(types.Symbol)
			if !ok {
				return nil, fmt.Errorf("invalid root key type %T", key)
			}
			skey = string(symkey)
		}
		data[skey] = val
	}

	result := &component.ConfigData{
		Data: data,
	}
	if input.Source != nil {
		result.Source = input.Source.Name
	}

	return result, nil
}

func ConfigDataProto(
	input *component.ConfigData,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_ConfigData, error) {
	i := make(map[interface{}]interface{}, len(input.Data))
	for key, val := range input.Data {
		i[key] = val
	}
	iproto, err := HashProto(i, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	return &vagrant_plugin_sdk.Args_ConfigData{
		Data: iproto,
		Source: &vagrant_plugin_sdk.Args_Class{
			Name: input.Source,
		},
	}, nil
}

func ConfigMerge(
	input *vagrant_plugin_sdk.Config_Merge,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*component.ConfigMerge, error) {
	base, err := ConfigData(input.Base, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	overlay, err := ConfigData(input.Overlay, log, internal, ctx)
	if err != nil {
		return nil, err
	}

	return &component.ConfigMerge{
		Base:    base,
		Overlay: overlay,
	}, nil
}

func ConfigMergeProto(
	input *component.ConfigMerge,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Config_Merge, error) {
	base, err := ConfigDataProto(input.Base, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	overlay, err := ConfigDataProto(input.Overlay, log, internal, ctx)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Config_Merge{
		Base:    base,
		Overlay: overlay,
	}, nil
}

func ConfigFinalize(
	input *vagrant_plugin_sdk.Config_Finalize,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*component.ConfigFinalize, error) {
	data, err := ConfigData(input.Config, log, internal, ctx)
	if err != nil {
		return nil, err
	}

	return &component.ConfigFinalize{
		Config: data,
	}, nil
}

func ConfigFinalizeProto(
	input *component.ConfigFinalize,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Config_Finalize, error) {
	data, err := ConfigDataProto(input.Config, log, internal, ctx)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Config_Finalize{
		Config: data,
	}, nil
}

func Hash(
	input *vagrant_plugin_sdk.Args_Hash,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (result map[interface{}]interface{}, err error) {
	result = make(map[interface{}]interface{}, len(input.GetEntries()))

	for _, e := range input.GetEntries() {
		r, err := Direct(
			&vagrant_plugin_sdk.Args_Direct{
				Arguments: []*anypb.Any{
					e.Key,
					e.Value,
				},
			},
			log,
			internal,
			ctx,
		)

		if err != nil {
			return nil, err
		}

		result[r.Arguments[0]] = r.Arguments[1]
	}

	return
}

func HashProto(
	input map[interface{}]interface{},
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_Hash, error) {
	content := make([]*vagrant_plugin_sdk.Args_HashEntry, 0, len(input))

	for k, v := range input {
		r, err := DirectProto(
			&component.Direct{
				Arguments: []interface{}{k, v},
			},
			log,
			internal,
			ctx,
		)

		if err != nil {
			return nil, err
		}

		content = append(content, &vagrant_plugin_sdk.Args_HashEntry{
			Key:   r.Arguments[0],
			Value: r.Arguments[1],
		})
	}

	return &vagrant_plugin_sdk.Args_Hash{
		Entries: content,
	}, nil
}

func Options(
	input *vagrant_plugin_sdk.Args_Options,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (result types.Options, err error) {
	return Hash(
		input.Options,
		log,
		internal,
		ctx,
	)
}

func OptionsProto(
	input types.Options,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (opts *vagrant_plugin_sdk.Args_Options, err error) {
	h, err := HashProto(input, log, internal, ctx)
	if err != nil {
		return
	}
	return &vagrant_plugin_sdk.Args_Options{
		Options: h,
	}, nil
}

func Folders(
	input *vagrant_plugin_sdk.Args_Folders,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (result types.Folders, err error) {
	return Hash(
		input.Folders,
		log,
		internal,
		ctx,
	)
}

func FoldersProto(
	input types.Folders,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (opts *vagrant_plugin_sdk.Args_Folders, err error) {
	h, err := HashProto(input, log, internal, ctx)
	if err != nil {
		return
	}
	return &vagrant_plugin_sdk.Args_Folders{
		Folders: h,
	}, nil
}

func NamedCapability(
	input *vagrant_plugin_sdk.Args_NamedCapability,
) *component.NamedCapability {
	return &component.NamedCapability{
		Capability: input.Capability,
	}
}

func NamedCapabilityProto(
	input *component.NamedCapability,
) *vagrant_plugin_sdk.Args_NamedCapability {
	return &vagrant_plugin_sdk.Args_NamedCapability{
		Capability: input.Capability,
	}
}

func Direct(
	input *vagrant_plugin_sdk.Args_Direct,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*component.Direct, error) {
	args := make([]interface{}, len(input.Arguments))

	for i := 0; i < len(args); i++ {
		v := input.Arguments[i]
		// List item are Any values so start with decoding
		_, val, err := dynamic.DecodeAny(v)
		if err != nil {
			return nil, err
		}

		// If the value is nil, check for that before we start
		// trying maps
		if _, ok := val.(*vagrant_plugin_sdk.Args_Null); ok {
			args[i] = nil
			continue
		}

		// First attempt to map the decoded value using
		// the well known type protos
		nv, err := dynamic.MapFromWellKnownProto(val.(proto.Message))

		// If we didn't generate an error, set the value and move on
		if err == nil {
			args[i] = nv
			continue
		}

		// Next attempt a blind map to convert the value into something
		// we may have a converter for
		nv, err = dynamic.BlindMap(val, internal.Mappers(),
			argmapper.Typed(internal, ctx, log))

		// Again, if there's no error, set the value and move on
		if err == nil {
			args[i] = nv
			continue
		}

		// Log the mapping failure to the debug output
		log.Warn("failed to map decoded direct argument",
			"value", val,
			"error", err,
		)

		// return nil, err

		// TODO(spox): I don't think we should do this but instead
		// force an error when we can't convert

		// Set the decoded value into the result set since it's
		// the best we can do
		args[i] = val
	}

	return &component.Direct{Arguments: args}, nil
}

func DirectProto(
	input *component.Direct,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Args_Direct, error) {
	list := make([]*anypb.Any, len(input.Arguments))
	for i := 0; i < len(list); i++ {
		arg := input.Arguments[i]
		var v interface{}
		v, err := dynamic.MapToWellKnownProto(arg)
		if err != nil {
			v, err = dynamic.UnknownMap(arg, (*proto.Message)(nil),
				internal.Mappers(),
				argmapper.Typed(internal),
				argmapper.Typed(ctx),
				argmapper.Typed(log),
			)
			if err != nil {
				return nil, err
			}
		}

		if v == nil {
			v = &vagrant_plugin_sdk.Args_Null{}
		}

		if list[i], err = dynamic.EncodeAny(v.(proto.Message)); err != nil {
			return nil, err
		}
	}

	return &vagrant_plugin_sdk.Args_Direct{
		Arguments: list,
	}, nil
}

func HostProto(
	input component.Host,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Host, error) {
	cid := fmt.Sprintf("%p", input)
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_Host), nil
	}
	p := &plugincomponent.HostPlugin{
		BasePlugin: basePlugin(input, internal),
		Impl:       input,
	}

	internal.Logger().Trace("wrapping host plugin",
		"host", input)

	id, ep, err := wrapClient(input, p, internal)
	if err != nil {
		internal.Logger().Warn("failed to wrap host plugin",
			"host", input,
			"error", err)

		return nil, err
	}
	proto := &vagrant_plugin_sdk.Args_Host{
		Network:  ep.Network(),
		Addr:     ep.String(),
		StreamId: id,
	}
	internal.Cache().Register(cid, proto)
	return proto, nil
}

func Host(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Host,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Host, error) {
	p := &plugincomponent.HostPlugin{
		BasePlugin: basePlugin(nil, internal),
	}
	internal.Logger().Trace("connecting to wrapped host plugin",
		"connection-info", input)

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		internal.Logger().Warn("failed to connect to wrapped host plugin",
			"connection-info", input,
			"error", err)

		return nil, err
	}

	return client.(core.Host), nil
}

func GuestProto(
	input component.Guest,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Guest, error) {
	p := &plugincomponent.GuestPlugin{
		BasePlugin: basePlugin(input, internal),
		Impl:       input,
	}

	internal.Logger().Trace("wrapping guest plugin", "guest", input)
	id, ep, err := wrapClient(input, p, internal)
	if err != nil {
		internal.Logger().Warn("failed to wrap guest plugin", "guest", input, "error", err)
		return nil, err
	}
	return &vagrant_plugin_sdk.Args_Guest{
		Network:  ep.Network(),
		Addr:     ep.String(),
		StreamId: id,
	}, nil
}

func Guest(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Guest,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Guest, error) {
	p := &plugincomponent.GuestPlugin{
		BasePlugin: basePlugin(nil, internal),
	}
	internal.Logger().Trace("connecting to wrapped guest plugin", "connection-info", input)
	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		internal.Logger().Warn("failed to connect to wrapped guest plugin", "connection-info", input, "error", err)
		return nil, err
	}

	return client.(core.Guest), nil
}

// Flags maps

func Flags(input []*vagrant_plugin_sdk.Command_Flag) (flags []*component.CommandFlag, err error) {
	flags = make([]*component.CommandFlag, len(input))

	for i, f := range input {
		flags[i] = &component.CommandFlag{
			LongName:     f.LongName,
			ShortName:    f.ShortName,
			Description:  f.Description,
			DefaultValue: f.DefaultValue,
			Aliases:      f.Aliases,
		}

		switch f.Type {
		case vagrant_plugin_sdk.Command_Flag_STRING:
			flags[i].Type = component.FlagString
		case vagrant_plugin_sdk.Command_Flag_BOOL:
			flags[i].Type = component.FlagBool
		}
	}
	return
}

// Flags maps
func FlagsProto(input []*component.CommandFlag) (output []*vagrant_plugin_sdk.Command_Flag, err error) {
	output = make([]*vagrant_plugin_sdk.Command_Flag, len(input))

	for i, f := range input {
		output[i] = &vagrant_plugin_sdk.Command_Flag{
			LongName:     f.LongName,
			ShortName:    f.ShortName,
			Description:  f.Description,
			DefaultValue: f.DefaultValue,
			Aliases:      f.Aliases,
		}

		switch f.Type {
		case component.FlagString:
			output[i].Type = vagrant_plugin_sdk.Command_Flag_STRING
		case component.FlagBool:
			output[i].Type = vagrant_plugin_sdk.Command_Flag_BOOL
		default:
			err = fmt.Errorf("invalid flag type - %s", f.Type.String())
			return
		}
	}

	return
}

func MapToProto(input map[string]interface{}) (*structpb.Struct, error) {
	return structpb.NewStruct(input)
}

func ProtoToMap(input *structpb.Struct) (map[string]interface{}, error) {
	return input.AsMap(), nil
}

func MachineState(input *vagrant_plugin_sdk.Args_Target_Machine_State) (*core.MachineState, error) {
	var result core.MachineState
	return &result, mapstructure.Decode(input, &result)
}

func MachineStateProto(input *core.MachineState) (*vagrant_plugin_sdk.Args_Target_Machine_State, error) {
	var result vagrant_plugin_sdk.Args_Target_Machine_State
	return &result, mapstructure.Decode(input, &result)
}

func SshInfo(input *vagrant_plugin_sdk.Args_Connection_SSHInfo) (*core.SshInfo, error) {
	var result core.SshInfo
	return &result, mapstructure.Decode(input, &result)
}

func SshInfoProto(input *core.SshInfo) (*vagrant_plugin_sdk.Args_Connection_SSHInfo, error) {
	var result vagrant_plugin_sdk.Args_Connection_SSHInfo
	return &result, mapstructure.Decode(input, &result)
}

func CorePluginManager(ctx context.Context,
	input *vagrant_plugin_sdk.Args_CorePluginManager,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.CorePluginManager, error) {
	cid := input.Addr
	if cid != "" {
		if ch := internal.Cache().Get(cid); ch != nil {
			log.Trace("using cached core plugin manager", "cid", cid)
			return ch.(core.CorePluginManager), nil
		}
	}

	// Create our plugin
	p := &plugincore.CorePluginManagerPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)

	log.Trace("cache miss on core plugin manager", "cid", cid)

	return client.(core.CorePluginManager), nil
}

func CorePluginManagerProto(
	impl core.CorePluginManager,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_CorePluginManager, error) {
	cid := fmt.Sprintf("%p", impl)
	if ch := internal.Cache().Get(cid); ch != nil {
		log.Trace("using cached core plugin manager proto", "cid", cid)
		return ch.(*vagrant_plugin_sdk.Args_CorePluginManager), nil
	}

	// Create our plugin
	p := &plugincore.CorePluginManagerPlugin{
		BasePlugin: basePlugin(impl, internal),
		Impl:       impl,
	}

	id, ep, err := wrapClient(impl, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_CorePluginManager{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}

	internal.Cache().Register(cid, proto)

	log.Trace("cache miss on core plugin manager proto", "cid", cid)

	return proto, nil
}

func Box(ctx context.Context,
	input *vagrant_plugin_sdk.Args_Box,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Box, error) {
	// Create our plugin
	p := &plugincore.BoxPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Box), nil
}

func BoxProto(
	box core.Box,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Box, error) {
	n, err := box.Name()
	if err != nil {
		return nil, err
	}
	v, err := box.Version()
	if err != nil {
		return nil, err
	}
	pr, err := box.Provider()
	if err != nil {
		return nil, err
	}
	cid := n + "-" + v + "-" + pr
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_Box), nil
	}

	// Create our plugin
	p := &plugincore.BoxPlugin{
		BasePlugin: basePlugin(box, internal),
		Impl:       box,
	}

	id, ep, err := wrapClient(box, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Box{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	log.Trace("registered box into cache",
		"cid", cid,
		"proto", proto,
		"cache", hclog.Fmt("%p", internal.Cache()),
	)

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func BoxCollection(ctx context.Context,
	input *vagrant_plugin_sdk.Args_BoxCollection,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.BoxCollection, error) {
	// Create our plugin
	p := &plugincore.BoxCollectionPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.BoxCollection), nil
}

func BoxCollectionProto(
	boxCollection core.BoxCollection,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_BoxCollection, error) {
	cid := fmt.Sprintf("%p", boxCollection)
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_BoxCollection), nil
	}

	// Create our plugin
	p := &plugincore.BoxCollectionPlugin{
		BasePlugin: basePlugin(boxCollection, internal),
		Impl:       boxCollection,
	}

	log.Warn("wrapping box to generate proto", "cid", cid)

	id, ep, err := wrapClient(boxCollection, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_BoxCollection{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	log.Trace("registered box collection into cache",
		"cid", cid,
		"proto", proto,
		"cache", hclog.Fmt("%p", internal.Cache()),
	)

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func BoxMetadata(ctx context.Context,
	input *vagrant_plugin_sdk.Args_BoxMetadata,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.BoxMetadata, error) {
	// Create our plugin
	p := &plugincore.BoxMetadataPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.BoxMetadata), nil
}

func BoxMetadataProto(
	boxMetadata core.BoxMetadata,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_BoxMetadata, error) {
	n := boxMetadata.BoxName()
	cid := "box_metadata" + n
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_BoxMetadata), nil
	}

	// Create our plugin
	p := &plugincore.BoxMetadataPlugin{
		BasePlugin: basePlugin(boxMetadata, internal),
		Impl:       boxMetadata,
	}

	id, ep, err := wrapClient(boxMetadata, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_BoxMetadata{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	log.Trace("registered box metadata into cache",
		"cid", cid,
		"proto", proto,
		"cache", hclog.Fmt("%p", internal.Cache()),
	)

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func BoxResponse(ctx context.Context,
	input *vagrant_plugin_sdk.Target_Machine_BoxResponse,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Box, error) {
	// Return nil if the response is empty
	if input.Box == nil {
		return nil, nil
	}

	// Create our plugin
	p := &plugincore.BoxPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input.Box, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Box), nil
}

func BoxResponseProto(
	box core.Box,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Target_Machine_BoxResponse, error) {
	result := &vagrant_plugin_sdk.Target_Machine_BoxResponse{}

	// A nil box and an empty result are valid
	if box == nil {
		return result, nil
	}

	n, err := box.Name()
	if err != nil {
		return nil, err
	}
	v, err := box.Version()
	if err != nil {
		return nil, err
	}
	pr, err := box.Provider()
	if err != nil {
		return nil, err
	}
	cid := n + "-" + v + "-" + pr
	if ch := internal.Cache().Get(cid); ch != nil {
		result.Box = ch.(*vagrant_plugin_sdk.Args_Box)
		return result, nil
	}

	// Create our plugin
	p := &plugincore.BoxPlugin{
		BasePlugin: basePlugin(box, internal),
		Impl:       box,
	}

	id, ep, err := wrapClient(box, p, internal)
	if err != nil {
		return nil, err
	}

	result.Box = &vagrant_plugin_sdk.Args_Box{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	log.Trace("registered box into cache",
		"cid", cid,
		"proto", result.Box,
		"cache", hclog.Fmt("%p", internal.Cache()),
	)

	internal.Cache().Register(cid, result.Box)

	return result, nil
}

// JobInfo maps Args.JobInfo to component.JobInfo.
func JobInfo(input *vagrant_plugin_sdk.Args_JobInfo) (*component.JobInfo, error) {
	var result component.JobInfo
	return &result, mapstructure.Decode(input, &result)
}

// JobInfoProto
func JobInfoProto(input *component.JobInfo) (*vagrant_plugin_sdk.Args_JobInfo, error) {
	var result vagrant_plugin_sdk.Args_JobInfo
	return &result, mapstructure.Decode(input, &result)
}

func DatadirBasis(input *vagrant_plugin_sdk.Args_DataDir_Basis) *datadir.Basis {
	dir := datadir.NewBasicDir(input.ConfigDir, input.CacheDir, input.DataDir, input.TempDir)
	return &datadir.Basis{Dir: dir}
}

func DatadirProject(input *vagrant_plugin_sdk.Args_DataDir_Project) *datadir.Project {
	dir := datadir.NewBasicDir(input.ConfigDir, input.CacheDir, input.DataDir, input.TempDir)
	return &datadir.Project{Dir: dir}
}

func DatadirTarget(input *vagrant_plugin_sdk.Args_DataDir_Target) *datadir.Target {
	dir := datadir.NewBasicDir(input.ConfigDir, input.CacheDir, input.DataDir, input.TempDir)
	return &datadir.Target{Dir: dir}
}

func DatadirComponent(input *vagrant_plugin_sdk.Args_DataDir_Project) *datadir.Component {
	dir := datadir.NewBasicDir(input.ConfigDir, input.CacheDir, input.DataDir, input.TempDir)
	return &datadir.Component{Dir: dir}
}

func DatadirBasisProto(input *datadir.Basis) *vagrant_plugin_sdk.Args_DataDir_Basis {
	return &vagrant_plugin_sdk.Args_DataDir_Basis{
		CacheDir:  input.CacheDir().String(),
		DataDir:   input.DataDir().String(),
		TempDir:   input.TempDir().String(),
		ConfigDir: input.ConfigDir().String(),
	}
}

func DatadirProjectProto(input *datadir.Project) *vagrant_plugin_sdk.Args_DataDir_Project {
	return &vagrant_plugin_sdk.Args_DataDir_Project{
		CacheDir:  input.CacheDir().String(),
		DataDir:   input.DataDir().String(),
		TempDir:   input.TempDir().String(),
		ConfigDir: input.ConfigDir().String(),
	}
}

func DatadirTargetProto(input *datadir.Target) *vagrant_plugin_sdk.Args_DataDir_Target {
	return &vagrant_plugin_sdk.Args_DataDir_Target{
		CacheDir:  input.CacheDir().String(),
		DataDir:   input.DataDir().String(),
		TempDir:   input.TempDir().String(),
		ConfigDir: input.ConfigDir().String(),
	}
}

func DatadirComponentProto(input *datadir.Component) *vagrant_plugin_sdk.Args_DataDir_Component {
	return &vagrant_plugin_sdk.Args_DataDir_Component{
		CacheDir:  input.CacheDir().String(),
		DataDir:   input.DataDir().String(),
		TempDir:   input.TempDir().String(),
		ConfigDir: input.ConfigDir().String(),
	}
}

// Logger maps *vagrant_plugin_sdk.Args_Logger to an hclog.Logger
func Logger(input *vagrant_plugin_sdk.Args_Logger) hclog.Logger {
	// We use the default logger as the base. Within a plugin we always set
	// it so we can confidently use this. This lets plugins potentially mess
	// with this but that's a risk we have to take.
	return hclog.L().ResetNamed(input.Name)
}

func LoggerProto(log hclog.Logger) *vagrant_plugin_sdk.Args_Logger {
	return &vagrant_plugin_sdk.Args_Logger{
		Name: log.Name(),
	}
}

// TerminalUI maps *vagrant_plugin_sdk.Args_TerminalUI to an hclog.TerminalUI
func TerminalUI(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_TerminalUI,
	log hclog.Logger,
	internal pluginargs.Internal,
) (terminal.UI, error) {
	cid := input.Addr
	if cid != "" {
		if ch := internal.Cache().Get(cid); ch != nil {
			return ch.(terminal.UI), nil
		}
	}

	// Create our plugin
	p := &pluginterminal.UIPlugin{
		Mappers: internal.Mappers(),
		Logger:  log,
	}

	internal.Logger().Trace("connecting to wrapped ui",
		"stream_id", input.StreamId,
	)

	client, err := wrapConnect(ctx, p, input, internal)

	if err != nil {
		internal.Logger().Warn("failed to connect to wrapped ui",
			"steam_id", input.StreamId,
			"error", err,
		)

		return nil, err
	}

	internal.Logger().Trace("connected to wrapped ui",
		"ui", client,
		"stream_id", input.StreamId,
	)

	internal.Cache().Register(cid, client)

	return client.(terminal.UI), nil
}

func TerminalUIProto(
	ui terminal.UI,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_TerminalUI, error) {
	cid := fmt.Sprintf("%p", ui)
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_TerminalUI), nil
	}

	// Create our plugin
	p := &pluginterminal.UIPlugin{
		Impl:    ui,
		Mappers: internal.Mappers(),
		Logger:  log.ResetNamed("vagrant.wrapped"),
	}

	internal.Logger().Trace("wrapping ui", "ui", ui)
	id, ep, err := wrapClient(ui, p, internal)

	if err != nil {
		internal.Logger().Trace("failed to wrap ui", "ui", ui, "error", err)
		return nil, err
	}

	internal.Logger().Trace("wrapped ui", "ui", ui, "stream_id", id, "endpoint", ep)
	proto := &vagrant_plugin_sdk.Args_TerminalUI{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func MetadataSet(input *vagrant_plugin_sdk.Args_MetadataSet) *component.MetadataSet {
	return &component.MetadataSet{
		Metadata: input.Metadata,
	}
}

func MetadataSetProto(meta *component.MetadataSet) *vagrant_plugin_sdk.Args_MetadataSet {
	return &vagrant_plugin_sdk.Args_MetadataSet{Metadata: meta.Metadata}
}

func Plugin(
	ctx context.Context,
	input *vagrant_plugin_sdk.PluginManager_Plugin,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*core.NamedPlugin, error) {
	t, err := component.FindType(input.Type)
	if err != nil {
		return nil, err
	}
	result := &core.NamedPlugin{
		Name: input.Name,
		Type: t.String(),
	}
	if input.Options != nil {
		result.Options, err = component.UnmarshalOptionsProto(t, input.Options)
		if err != nil {
			return nil, fmt.Errorf("problem unmarshalling plugin options: %s", err)
		}
	}

	if input.Plugin == nil {
		return result, nil
	}

	args := []argmapper.Arg{
		argmapper.ConverterFunc(internal.Mappers()...),
		argmapper.Typed(ctx, log, internal),
	}
	_, v, err := dynamic.DecodeAny(input.Plugin)
	raw, err := dynamic.Map(v, component.TypeMap[t], args...)
	if err != nil {
		return nil, err
	}
	result.Plugin = raw

	return result, nil
}

func PluginProto(
	input *core.NamedPlugin,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.PluginManager_Plugin, error) {
	t, err := component.FindType(input.Type)
	if err != nil {
		return nil, err
	}
	result := &vagrant_plugin_sdk.PluginManager_Plugin{
		Name: input.Name,
		Type: t.String(),
	}

	if input.Options != nil {
		result.Options, err = component.ProtoAny(input.Options)
		if err != nil {
			return nil, fmt.Errorf("problem converting plugin options to proto: %s", err)
		}
	}

	if input.Plugin == nil {
		return result, nil
	}

	raw, err := dynamic.UnknownMap(input.Plugin,
		(*proto.Message)(nil),
		internal.Mappers(),
		argmapper.Typed(log, internal),
	)
	if err != nil {
		return nil, err
	}
	v, err := dynamic.EncodeAny(raw.(proto.Message))
	if err != nil {
		return nil, err
	}
	result.Plugin = v

	return result, nil
}

func Plugins(
	ctx context.Context,
	input *vagrant_plugin_sdk.PluginManager_PluginsResponse,
	log hclog.Logger,
	internal pluginargs.Internal,
) ([]*core.NamedPlugin, error) {
	result := make([]*core.NamedPlugin, len(input.Plugins))
	for i, np := range input.Plugins {
		raw, err := dynamic.Map(np,
			(**core.NamedPlugin)(nil),
			argmapper.ConverterFunc(internal.Mappers()...),
			argmapper.Typed(ctx, log, internal),
		)
		if err != nil {
			return nil, err
		}
		result[i] = raw.(*core.NamedPlugin)
	}

	return result, nil
}

func PluginsProto(
	input []*core.NamedPlugin,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.PluginManager_PluginsResponse, error) {
	result := &vagrant_plugin_sdk.PluginManager_PluginsResponse{
		Plugins: make([]*vagrant_plugin_sdk.PluginManager_Plugin, len(input)),
	}

	for i, np := range input {

		raw, err := PluginProto(np, log, internal)
		if err != nil {
			return nil, err
		}
		result.Plugins[i] = raw
	}

	return result, nil
}

func PluginManager(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_PluginManager,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.PluginManager, error) {
	p := &plugincore.PluginManagerPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.PluginManager), err
}

func PluginManagerProto(
	input core.PluginManager,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_PluginManager, error) {
	p := &plugincore.PluginManagerPlugin{
		BasePlugin: basePlugin(input, internal),
		Impl:       input,
	}
	id, ep, err := wrapClient(input, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_PluginManager{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}

	return proto, nil
}

// StateBag maps StateBag proto to core.StateBag.
func StateBag(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_StateBag,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.StateBag, error) {
	// Create our plugin
	p := &plugincore.StateBagPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.StateBag), nil
}

func StateBagProto(
	bag core.StateBag,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_StateBag, error) {
	cid := fmt.Sprintf("%p", bag)
	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_StateBag), nil
	}

	// Create our plugin
	p := &plugincore.StateBagPlugin{
		BasePlugin: basePlugin(bag, internal),
		Impl:       bag,
	}

	id, ep, err := wrapClient(bag, p, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_StateBag{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func CommandParams(
	input *vagrant_plugin_sdk.Command_Arguments,
) (c *component.CommandParams) {
	c = &component.CommandParams{
		Arguments: input.Args,
		Flags:     map[string]interface{}{},
	}

	for _, f := range input.Flags {
		switch f.Type {
		case vagrant_plugin_sdk.Command_Arguments_Flag_BOOL:
			c.Flags[f.Name] = f.GetBool()
		case vagrant_plugin_sdk.Command_Arguments_Flag_STRING:
			c.Flags[f.Name] = f.GetString_()
		}
	}

	return
}

func CommandInfoFromResponse(
	input *vagrant_plugin_sdk.Command_CommandInfoResp,
) *vagrant_plugin_sdk.Command_CommandInfo {
	return input.CommandInfo
}

func CommandInfo(input *vagrant_plugin_sdk.Command_CommandInfo) (*component.CommandInfo, error) {
	flags, err := Flags(input.Flags)

	subcommands := []*component.CommandInfo{}
	for _, cmd := range input.Subcommands {
		subcommand, err := CommandInfo(cmd)
		if err != nil {
			return nil, err
		}
		subcommands = append(subcommands, subcommand)
	}

	result := &component.CommandInfo{
		Flags:       flags,
		Name:        input.Name,
		Help:        input.Help,
		Synopsis:    input.Synopsis,
		Subcommands: subcommands,
		Primary:     input.Primary,
	}
	return result, err
}

func CommandInfoProto(input *component.CommandInfo) (*vagrant_plugin_sdk.Command_CommandInfo, error) {
	var result vagrant_plugin_sdk.Command_CommandInfo
	err := mapstructure.Decode(input, &result)
	if err != nil {
		return nil, err
	}
	result.Flags, err = FlagsProto(input.Flags)
	subcmds := []*vagrant_plugin_sdk.Command_CommandInfo{}
	for _, cmd := range input.Subcommands {
		toAdd, err := CommandInfoProto(cmd)
		if err != nil {
			return nil, err
		}
		subcmds = append(subcmds, toAdd)
	}
	result.Subcommands = subcmds
	return &result, err
}

func StateProto(s core.State) *vagrant_plugin_sdk.Args_Target_State {
	var state vagrant_plugin_sdk.Args_Target_State_State
	switch s {
	case core.CREATED:
		state = vagrant_plugin_sdk.Args_Target_State_CREATED
	case core.DESTROYED:
		state = vagrant_plugin_sdk.Args_Target_State_DESTROYED
	case core.PENDING:
		state = vagrant_plugin_sdk.Args_Target_State_PENDING
	default:
		state = vagrant_plugin_sdk.Args_Target_State_UNKNOWN
	}
	return &vagrant_plugin_sdk.Args_Target_State{
		State: state,
	}
}

func State(s *vagrant_plugin_sdk.Args_Target_State) (state core.State) {
	switch s.State {
	case vagrant_plugin_sdk.Args_Target_State_CREATED:
		state = core.CREATED
	case vagrant_plugin_sdk.Args_Target_State_DESTROYED:
		state = core.DESTROYED
	case vagrant_plugin_sdk.Args_Target_State_PENDING:
		state = core.PENDING
	default:
		state = core.UNKNOWN
	}
	return
}

func MetadataProto(m map[string]string) *vagrant_plugin_sdk.Args_MetadataSet {
	return &vagrant_plugin_sdk.Args_MetadataSet{
		Metadata: m,
	}
}

func Metadata(m *vagrant_plugin_sdk.Args_MetadataSet) map[string]string {
	return m.Metadata
}

func BasisProto(
	b core.Basis,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Basis, error) {
	cid := fmt.Sprintf("%p", b)

	if ch := internal.Cache().Get(cid); ch != nil {
		log.Trace("using cached basis proto", "cid", cid)
		return ch.(*vagrant_plugin_sdk.Args_Basis), nil
	}

	bp := &plugincore.BasisPlugin{
		BasePlugin: basePlugin(b, internal),
		Impl:       b,
	}

	id, ep, err := wrapClient(b, bp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Basis{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}

	internal.Cache().Register(cid, proto)

	log.Trace("cache miss on basis proto", "cid", cid)

	return proto, nil
}

func Basis(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Basis,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Basis, error) {
	cid := input.Addr
	if cid != "" {
		if ch := internal.Cache().Get(cid); ch != nil {
			log.Trace("using cached basis", "cid", cid)
			return ch.(core.Basis), nil
		}
	}

	b := &plugincore.BasisPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, b, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)

	log.Trace("cache miss on basis", "cid", cid)

	return client.(core.Basis), nil
}

func CommunicatorCommand(
	c *vagrant_plugin_sdk.Communicator_Command,
) ([]string, error) {
	return []string{c.Command}, nil
}

func CommunicatorCommandProto(
	c []string,
) (*vagrant_plugin_sdk.Communicator_Command, error) {
	return &vagrant_plugin_sdk.Communicator_Command{
		Command: strings.Join(c[:], " "),
	}, nil
}

func CommunicatorProto(
	c component.Communicator,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Communicator, error) {
	cp := &plugincomponent.CommunicatorPlugin{
		BasePlugin: basePlugin(c, internal),
		Impl:       c,
	}

	id, ep, err := wrapClient(c, cp, internal)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Args_Communicator{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}, nil
}

func Communicator(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Communicator,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Communicator, error) {
	p := &plugincomponent.CommunicatorPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Communicator), nil
}

func PushProto(
	c component.Push,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Push, error) {
	cp := &plugincomponent.PushPlugin{
		BasePlugin: basePlugin(c, internal),
		Impl:       c,
	}

	id, ep, err := wrapClient(c, cp, internal)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Args_Push{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}, nil
}

func Push(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Push,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Push, error) {
	p := &plugincomponent.PushPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Push), nil
}

func ProjectProto(
	p core.Project,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Project, error) {
	cid := fmt.Sprintf("%p", p)

	if ch := internal.Cache().Get(cid); ch != nil {
		return ch.(*vagrant_plugin_sdk.Args_Project), nil
	}

	pp := &plugincore.ProjectPlugin{
		BasePlugin: basePlugin(p, internal),
		Impl:       p,
	}

	id, ep, err := wrapClient(p, pp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Project{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	internal.Cache().Register(cid, proto)

	return proto, nil
}

func Project(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Project,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Project, error) {
	cid := input.Addr
	if cid != "" {
		if ch := internal.Cache().Get(cid); ch != nil {
			return ch.(core.Project), nil
		}
	}

	p := &plugincore.ProjectPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)

	return client.(core.Project), nil
}

func CommandProto(
	c component.Command,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Command, error) {
	cp := &plugincomponent.CommandPlugin{
		BasePlugin: basePlugin(c, internal),
		Impl:       c,
	}

	id, ep, err := wrapClient(c, cp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Command{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String()}

	return proto, nil
}

func Command(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Command,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Command, error) {
	p := &plugincomponent.CommandPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, p, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Command), nil
}

func VagrantfileSyncedFolderToFolder(
	f *vagrant_plugin_sdk.Target_Machine_SyncedFoldersResponse_Folder,
) (*core.Folder, error) {
	var result *core.Folder
	return result, mapstructure.Decode(f, &result)
}

func FolderToVagrantfileSyncedFolder(
	f *core.Folder,
	log hclog.Logger,
	internal pluginargs.Internal,
	ctx context.Context,
) (*vagrant_plugin_sdk.Target_Machine_SyncedFoldersResponse_Folder, error) {
	result := &vagrant_plugin_sdk.Target_Machine_SyncedFoldersResponse_Folder{}
	opts, err := HashProto(f.Options, log, internal, ctx)
	if err != nil {
		return nil, err
	}
	result.Source = f.Source.String()
	result.Destination = f.Destination.String()
	result.Options = opts
	return result, err
}

func SyncedFolderProto(
	s component.SyncedFolder,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_SyncedFolder, error) {
	rid := fmt.Sprintf("%p", s)
	if at := internal.Cache().Get(rid); at != nil {
		log.Trace("using cached synced folder value",
			"value", at,
		)

		return at.(*vagrant_plugin_sdk.Args_SyncedFolder), nil
	}

	sp := &plugincomponent.SyncedFolderPlugin{
		BasePlugin: basePlugin(s, internal),
		Impl:       s,
	}

	id, endpoint, err := wrapClient(s, sp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_SyncedFolder{
		StreamId: id,
		Network:  endpoint.Network(),
		Addr:     endpoint.String(),
	}

	internal.Cache().Register(rid, proto)

	return proto, nil
}

func SyncedFolder(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Provider,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.SyncedFolder, error) {
	s := &plugincomponent.SyncedFolderPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, s, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.SyncedFolder), nil
}

func ProviderProto(
	t component.Provider,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Provider, error) {
	cid := fmt.Sprintf("%p", t)
	if c := internal.Cache().Get(cid); c != nil {
		log.Trace("cache hit on provider proto",
			"cid", cid,
		)

		return c.(*vagrant_plugin_sdk.Args_Provider), nil
	}

	tp := &plugincomponent.ProviderPlugin{
		BasePlugin: basePlugin(t, internal),
		Impl:       t,
	}

	id, endpoint, err := wrapClient(t, tp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Provider{
		StreamId: id,
		Network:  endpoint.Network(),
		Addr:     endpoint.String(),
	}

	internal.Cache().Register(cid, proto)
	log.Trace("registering provider to cache",
		"cid", cid,
	)

	return proto, nil
}

func Provider(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Provider,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Provider, error) {
	cid := input.Addr
	if cid != "" {
		if c := internal.Cache().Get(cid); c != nil {
			log.Trace("cache hit on provider",
				"cid", cid,
			)

			return c.(core.Provider), nil
		}
	}

	t := &plugincomponent.ProviderPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, t, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)
	log.Trace("registering provider to cache",
		"cid", cid,
	)

	return client.(core.Provider), nil
}

func ProvisionerProto(
	t component.Provisioner,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Provisioner, error) {
	tp := &plugincomponent.ProvisionerPlugin{
		BasePlugin: basePlugin(t, internal),
		Impl:       t,
	}

	id, endpoint, err := wrapClient(t, tp, internal)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Args_Provisioner{
		StreamId: id,
		Network:  endpoint.Network(),
		Addr:     endpoint.String(),
	}, nil
}

func Provisioner(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Provisioner,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Provisioner, error) {
	t := &plugincomponent.ProvisionerPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, t, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.Provisioner), nil
}

func PathProto(
	p path.Path,
) *vagrant_plugin_sdk.Args_Path {
	return &vagrant_plugin_sdk.Args_Path{
		Path: p.String(),
	}
}

func Path(
	input *vagrant_plugin_sdk.Args_Path,
) path.Path {
	return path.NewPath(input.Path)
}

func MachineProject(
	m core.Machine,
) (core.Project, error) {
	return m.Project()
}

func TargetProject(
	t core.Target,
) (core.Project, error) {
	return t.Project()
}

func TargetProto(
	t core.Target,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Target, error) {
	rid, err := t.ResourceId()
	if err != nil {
		return nil, err
	}
	cid := fmt.Sprintf("target-%s-%p", rid, t)
	if at := internal.Cache().Get(cid); at != nil {
		log.Trace("using cached target value",
			"value", at,
		)

		return at.(*vagrant_plugin_sdk.Args_Target), nil
	}

	tp := &plugincore.TargetPlugin{
		BasePlugin: basePlugin(t, internal),
		Impl:       t,
	}

	id, endpoint, err := wrapClient(t, tp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Target{
		StreamId: id,
		Network:  endpoint.Network(),
		Addr:     endpoint.String(),
	}

	log.Trace("registering target proto to cache",
		"rid", rid,
		"proto", proto,
	)
	internal.Cache().Register(cid, proto)
	return proto, nil
}

func Target(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Target,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Target, error) {
	cid := fmt.Sprintf("target-%s", input.Addr)
	if v, ok := internal.Cache().Fetch(cid); ok {
		return v.(core.Target), nil
	}

	t := &plugincore.TargetPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, t, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)
	return client.(core.Target), nil
}

func TargetMachineProto(
	m core.Machine,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Target_Machine, error) {
	rid, err := m.ResourceId()
	if err != nil {
		return nil, err
	}

	cid := fmt.Sprintf("machine-%s", rid)
	if v, ok := internal.Cache().Fetch(cid); ok {
		return v.(*vagrant_plugin_sdk.Args_Target_Machine), nil
	}

	tp := &plugincore.TargetMachinePlugin{
		BasePlugin: basePlugin(m, internal),
		Impl:       m,
		TargetImpl: m,
	}

	id, endpoint, err := wrapClient(m, tp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Target_Machine{
		StreamId: id,
		Network:  endpoint.Network(),
		Addr:     endpoint.String(),
	}

	log.Trace("registering machine proto to cache",
		"rid", cid,
		"proto", proto,
	)
	internal.Cache().Register(cid, proto)
	return proto, nil
}

func TargetMachine(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Target_Machine,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Machine, error) {
	cid := fmt.Sprintf("machine-%s", input.Addr)
	if v, ok := internal.Cache().Fetch(cid); ok {
		return v.(core.Machine), nil
	}

	m := &plugincore.TargetMachinePlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, m, input, internal)
	if err != nil {
		return nil, err
	}

	internal.Cache().Register(cid, client)
	return client.(core.Machine), nil
}

func TargetIndexProto(
	t core.TargetIndex,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_TargetIndex, error) {
	ti := &plugincore.TargetIndexPlugin{
		BasePlugin: basePlugin(t, internal),
		Impl:       t,
	}

	id, ep, err := wrapClient(t, ti, internal)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Args_TargetIndex{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}, nil
}

func TargetIndex(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_TargetIndex,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.TargetIndex, error) {
	ti := &plugincore.TargetIndexPlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, ti, input, internal)
	if err != nil {
		return nil, err
	}

	return client.(core.TargetIndex), nil
}

func VagrantfileProto(
	v core.Vagrantfile,
	log hclog.Logger,
	internal pluginargs.Internal,
) (*vagrant_plugin_sdk.Args_Vagrantfile, error) {
	rid := fmt.Sprintf("%p", v)
	if av := internal.Cache().Get(rid); av != nil {
		log.Trace("using cache vagrantfile value",
			"value", av,
		)
		return av.(*vagrant_plugin_sdk.Args_Vagrantfile), nil
	}
	bp := &plugincore.VagrantfilePlugin{
		BasePlugin: basePlugin(v, internal),
		Impl:       v,
	}

	id, ep, err := wrapClient(v, bp, internal)
	if err != nil {
		return nil, err
	}

	proto := &vagrant_plugin_sdk.Args_Vagrantfile{
		StreamId: id,
		Network:  ep.Network(),
		Addr:     ep.String(),
	}

	log.Trace("registering vagrantfile proto to cache",
		"rid", rid,
		"proto", proto,
	)
	internal.Cache().Register(rid, proto)

	return proto, nil
}

func Vagrantfile(
	ctx context.Context,
	input *vagrant_plugin_sdk.Args_Vagrantfile,
	log hclog.Logger,
	internal pluginargs.Internal,
) (core.Vagrantfile, error) {
	b := &plugincore.VagrantfilePlugin{
		BasePlugin: basePlugin(nil, internal),
	}

	client, err := wrapConnect(ctx, b, input, internal)
	if err != nil {
		return nil, err
	}

	c, ok := client.(core.Vagrantfile)
	if !ok {
		return nil, fmt.Errorf("could not cast client to core.Vagrantfile interface (%#v)", client)
	}

	return c, nil
}

type connInfo interface {
	GetStreamId() uint32
	GetNetwork() string
	GetAddr() string
}

type hasTarget interface {
	SetTarget(net.Addr)
	Target() net.Addr
}

// When a core plugin is received, the proto will match the
// ConnInfo interface which provides the information needed
// setup a new client. Depending on the origin of the proto
// the client will either establish a direct connection to
// the service, or will connect via the broker.
func wrapConnect(
	ctx context.Context,
	p plugin.GRPCPlugin,
	i connInfo,
	internal pluginargs.Internal,
) (interface{}, error) {
	internal.Logger().Trace("connecting to wrapped plugin",
		"plugin", hclog.Fmt("%T", p),
		"connection", i,
		"broker", hclog.Fmt("%p", internal.Broker()))

	var err error
	var conn *grpc.ClientConn
	var addr net.Addr
	if target := i.GetAddr(); target != "" {
		switch i.GetNetwork() {
		case "tcp":
			addr, err = net.ResolveTCPAddr("tcp", target)
		case "unix":
			addr, err = net.ResolveUnixAddr("unix", target)
		default:
			return nil, fmt.Errorf(
				"Unknown target address type: %s", i.GetNetwork())
		}

		internal.Logger().Trace("connecting to wrapped plugin via direct target",
			"plugin", hclog.Fmt("%T", p),
			"target", target)

		// TODO(spox): grab dial options from pluginclient config
		conn, err = grpc.Dial("unused", grpc.WithDialer(
			func(_ string, _ time.Duration) (net.Conn, error) {
				return net.Dial(i.GetNetwork(), target)
			}), grpc.WithInsecure(),
		)
	} else {
		internal.Logger().Trace("connecting to wrapped plugin via broker",
			"plugin", hclog.Fmt("%T", p),
			"stream_id", i.GetStreamId(),
			"broker", hclog.Fmt("%p", internal.Broker()))

		conn, err = internal.Broker().Dial(i.GetStreamId())
	}
	if err != nil {
		internal.Logger().Warn("failed to connect to wrapped plugin",
			"plugin", hclog.Fmt("%T", p),
			"connection", i,
			"broker", hclog.Fmt("%p", internal.Broker()),
			"error", err)

		return nil, err
	}

	internal.Cleanup().Do(func() error { return conn.Close() })

	client, err := p.GRPCClient(ctx, internal.Broker(), conn)
	if err != nil {
		internal.Logger().Warn("failed to create client for wrapped plugin",
			"plugin", hclog.Fmt("%T", p),
			"connection", i,
			"broker", hclog.Fmt("%p", internal.Broker()),
			"error", err)

		return nil, err
	}

	if closer, ok := client.(io.Closer); ok {
		internal.Cleanup().Do(func() error { return closer.Close() })
	}

	if cache, ok := client.(cacher.HasCache); ok {
		cache.SetCache(internal.Cache())
	}

	internal.Logger().Trace("new client built for wrapped plugin",
		"plugin", hclog.Fmt("%T", p),
		"client", client,
		"connection", i,
		"broker", hclog.Fmt("%p", internal.Broker()))

	if addr != nil {
		if ec, ok := client.(hasTarget); ok {
			internal.Logger().Trace("setting direct target on new client",
				"plugin", hclog.Fmt("%T", p),
				"target", addr)

			ec.SetTarget(addr)
		} else {
			internal.Logger().Trace("client does not support direct targets for wrapped plugins",
				"plugin", hclog.Fmt("%T", p),
				"client", hclog.Fmt("%T", client))
		}
	}

	return client, nil
}

// This takes a plugin (which generally uses a client as the plugin implementation)
// and creates a new server for remote connections via the internal broker.
func wrapClientStandalone(
	impl interface{},
	p plugin.GRPCPlugin,
	broker *plugin.GRPCBroker,
	logger hclog.Logger,
) (id uint32, target net.Addr, closer func() error, err error) {
	// If an existing target exists for the implementation, use
	// that value for where to connect
	if iep, ok := impl.(hasTarget); ok {
		if target = iep.Target(); target != nil {
			logger.Trace("using preset wrapped plugin target",
				"plugin", hclog.Fmt("%T", p),
				"target", target)

			return
		}
	} else {
		logger.Trace("implementation does not support direct targets for wrapped plugins",
			"plugin", hclog.Fmt("%T", p),
			"implementation", hclog.Fmt("%T", impl),
		)
	}

	// Fetch the next available steam ID from the broker
	id = broker.NextId()

	// Since we want to register the target endpoint directly for
	// access off the configured broker, we need to get the listener
	// and setup the server directly instead of letting the plugin
	// library handle it for us
	l, err := broker.Accept(id)
	if err != nil {
		logger.Warn("failed to establish connection stream",
			"error", err)

		return
	}
	target = l.Addr()

	// Grab the shared plugin configuration so the expected
	// server configuration can be applied
	config := pluginclient.ClientConfig(logger)
	sopts := []grpc.ServerOption{}
	if config.TLSConfig != nil {
		sopts = append(sopts, grpc.Creds(credentials.NewTLS(config.TLSConfig)))
	}

	logger.Trace("starting listener for wrapped plugin",
		"broker", hclog.Fmt("%p", broker),
		"plugin", hclog.Fmt("%T", p),
		"stream_id", id,
		"target", target)

	server := plugin.DefaultGRPCServer(sopts)
	if err = p.GRPCServer(broker, server); err != nil {
		return
	}

	// Register a shutdown of this wrapped plugin server in our
	// cleanup so we don't leave it hanging around when closed
	closer = func() error {
		logger.Trace("shutting down listener for wrapped plugin",
			"broker", hclog.Fmt("%p", broker),
			"plugin", hclog.Fmt("%T", p),
			"stream_id", id,
			"target", target)

		server.GracefulStop()
		return nil
	}

	// Start serving
	go server.Serve(l)

	return
}

// This takes a plugin (which generally uses a client as the plugin implementation)
// and creates a new server for remote connections via the internal broker.
func wrapClient(
	impl interface{},
	p plugin.GRPCPlugin,
	internal pluginargs.Internal,
) (id uint32, target net.Addr, err error) {
	id, target, closer, err := wrapClientStandalone(
		impl,
		p,
		internal.Broker(),
		internal.Logger(),
	)

	if err != nil {
		return
	}

	internal.Cleanup().Do(closer)

	return
}

// creates a BasePlugin configuration using an optional
// source client and the internal args
func basePlugin(
	src interface{},
	internal pluginargs.Internal,
) *plugincomponent.BasePlugin {
	if w, ok := src.(wrappable); ok {
		return w.Wrap()
	}
	return &plugincomponent.BasePlugin{
		Cache:   internal.Cache(),
		Cleanup: internal.Cleanup(),
		Mappers: internal.Mappers(),
		Logger:  internal.Logger(),
		Wrapped: true,
	}
}

func init() {
	for _, fn := range All {
		mFn, err := argmapper.NewFunc(fn)
		if err != nil {
			panic(err)
		}
		plugincomponent.MapperFns = append(plugincomponent.MapperFns, mFn)
		plugincomponent.ProtomapperAllMap[reflect.TypeOf(fn)] = struct{}{}
	}
	AllFns := make([]*argmapper.Func, len(All))
	copy(AllFns, plugincomponent.MapperFns)

	for _, fn := range WellKnownTypes {
		mFn, err := argmapper.NewFunc(fn)
		if err != nil {
			panic(err)
		}
		dynamic.WellKnownTypeFns = append(dynamic.WellKnownTypeFns, mFn)
	}
	WellKnownFns = make([]*argmapper.Func, len(WellKnownTypes))
	copy(WellKnownFns, dynamic.WellKnownTypeFns)

	totLen := len(plugincomponent.MapperFns) + len(dynamic.WellKnownTypeFns)
	config.Mappers = make([]*argmapper.Func, totLen)
	copy(config.Mappers, dynamic.WellKnownTypeFns)
	copy(config.Mappers[len(dynamic.WellKnownTypeFns):totLen], plugincomponent.MapperFns)
}

type pluginMetadata interface {
	SetRequestMetadata(k, v string)
}

type wrappable interface {
	Wrap() *plugincomponent.BasePlugin
}
