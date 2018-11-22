// Code generated by "stringer -output stringer.go -type=DeclarationKind,ScopeKind,ChanDir,sentinel"; DO NOT EDIT.

package gc

import "strconv"

const _DeclarationKind_name = "ConstDeclarationFuncDeclarationImportDeclarationMethodDeclarationTypeDeclarationVarDeclaration"

var _DeclarationKind_index = [...]uint8{0, 16, 31, 48, 65, 80, 94}

func (i DeclarationKind) String() string {
	if i < 0 || i >= DeclarationKind(len(_DeclarationKind_index)-1) {
		return "DeclarationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeclarationKind_name[_DeclarationKind_index[i]:_DeclarationKind_index[i+1]]
}

const _ScopeKind_name = "UniverseScopePackageScopeFileScopeBlockScope"

var _ScopeKind_index = [...]uint8{0, 13, 25, 34, 44}

func (i ScopeKind) String() string {
	if i < 0 || i >= ScopeKind(len(_ScopeKind_index)-1) {
		return "ScopeKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScopeKind_name[_ScopeKind_index[i]:_ScopeKind_index[i+1]]
}

const _ChanDir_name = "TxChanRxChan"

var _ChanDir_index = [...]uint8{0, 6, 12}

func (i ChanDir) String() string {
	i -= 1
	if i < 0 || i >= ChanDir(len(_ChanDir_index)-1) {
		return "ChanDir(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ChanDir_name[_ChanDir_index[i]:_ChanDir_index[i+1]]
}

const _sentinel_name = "checkingcheckedcheckedFalsecheckedTrue"

var _sentinel_index = [...]uint8{0, 8, 15, 27, 38}

func (i sentinel) String() string {
	i -= 1
	if i >= sentinel(len(_sentinel_index)-1) {
		return "sentinel(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _sentinel_name[_sentinel_index[i]:_sentinel_index[i+1]]
}