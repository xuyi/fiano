//go:build !manifestcodegen
// +build !manifestcodegen

// Code generated by "menifestcodegen". DO NOT EDIT.
// To reproduce: go run github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/manifestcodegen/cmd/manifestcodegen github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest

package manifest

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/pretty"
)

var (
	// Just to avoid errors in "import" above in case if it wasn't used below
	_ = binary.LittleEndian
	_ = (fmt.Stringer)(nil)
	_ = (io.Reader)(nil)
	_ = pretty.Header
	_ = strings.Join
)

// NewSignature returns a new instance of Signature with
// all default values set.
func NewSignature() *Signature {
	s := &Signature{}
	// Set through tag "required":
	s.Version = 0x10
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *Signature) Validate() error {
	// See tag "require"
	if s.Version != 0x10 {
		return fmt.Errorf("field 'Version' expects value '0x10', but has %v", s.Version)
	}

	return nil
}

// ReadFrom reads the Signature from 'r' in format defined in the document #575623.
func (s *Signature) ReadFrom(r io.Reader) (int64, error) {
	totalN := int64(0)

	// SigScheme (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.SigScheme)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'SigScheme': %w", err)
		}
		totalN += int64(n)
	}

	// Version (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.Version)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Version': %w", err)
		}
		totalN += int64(n)
	}

	// KeySize (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.KeySize)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'KeySize': %w", err)
		}
		totalN += int64(n)
	}

	// HashAlg (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.HashAlg)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'HashAlg': %w", err)
		}
		totalN += int64(n)
	}

	// Data (ManifestFieldType: arrayDynamic)
	{
		size := uint16(s.KeySize.InBytes())
		s.Data = make([]byte, size)
		n, err := len(s.Data), binary.Read(r, binary.LittleEndian, s.Data)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Data': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *Signature) RehashRecursive() {
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *Signature) Rehash() {
}

// WriteTo writes the Signature into 'w' in format defined in
// the document #575623.
func (s *Signature) WriteTo(w io.Writer) (int64, error) {
	totalN := int64(0)
	s.Rehash()

	// SigScheme (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.SigScheme)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'SigScheme': %w", err)
		}
		totalN += int64(n)
	}

	// Version (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.Version)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Version': %w", err)
		}
		totalN += int64(n)
	}

	// KeySize (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.KeySize)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'KeySize': %w", err)
		}
		totalN += int64(n)
	}

	// HashAlg (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.HashAlg)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'HashAlg': %w", err)
		}
		totalN += int64(n)
	}

	// Data (ManifestFieldType: arrayDynamic)
	{
		n, err := len(s.Data), binary.Write(w, binary.LittleEndian, s.Data)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Data': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// SigSchemeSize returns the size in bytes of the value of field SigScheme
func (s *Signature) SigSchemeTotalSize() uint64 {
	return 2
}

// VersionSize returns the size in bytes of the value of field Version
func (s *Signature) VersionTotalSize() uint64 {
	return 1
}

// KeySizeSize returns the size in bytes of the value of field KeySize
func (s *Signature) KeySizeTotalSize() uint64 {
	return 2
}

// HashAlgSize returns the size in bytes of the value of field HashAlg
func (s *Signature) HashAlgTotalSize() uint64 {
	return 2
}

// DataSize returns the size in bytes of the value of field Data
func (s *Signature) DataTotalSize() uint64 {
	return uint64(len(s.Data))
}

// SigSchemeOffset returns the offset in bytes of field SigScheme
func (s *Signature) SigSchemeOffset() uint64 {
	return 0
}

// VersionOffset returns the offset in bytes of field Version
func (s *Signature) VersionOffset() uint64 {
	return s.SigSchemeOffset() + s.SigSchemeTotalSize()
}

// KeySizeOffset returns the offset in bytes of field KeySize
func (s *Signature) KeySizeOffset() uint64 {
	return s.VersionOffset() + s.VersionTotalSize()
}

// HashAlgOffset returns the offset in bytes of field HashAlg
func (s *Signature) HashAlgOffset() uint64 {
	return s.KeySizeOffset() + s.KeySizeTotalSize()
}

// DataOffset returns the offset in bytes of field Data
func (s *Signature) DataOffset() uint64 {
	return s.HashAlgOffset() + s.HashAlgTotalSize()
}

// Size returns the total size of the Signature.
func (s *Signature) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.SigSchemeTotalSize()
	size += s.VersionTotalSize()
	size += s.KeySizeTotalSize()
	size += s.HashAlgTotalSize()
	size += s.DataTotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *Signature) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "Signature", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Sig Scheme", "", &s.SigScheme, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Version", "", &s.Version, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Key Size", "", &s.KeySize, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Hash Alg", "", &s.HashAlg, opts...)...)
	// ManifestFieldType is arrayDynamic
	lines = append(lines, pretty.SubValue(depth+1, "Data", "", s.dataPrettyValue(), opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}
