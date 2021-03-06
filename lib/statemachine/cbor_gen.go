package statemachine

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *TestState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	// t.A (uint64) (uint64)
	if len("A") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"A\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("A")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("A")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.A))); err != nil {
		return err
	}

	// t.B (uint64) (uint64)
	if len("B") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"B\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("B")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("B")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.B))); err != nil {
		return err
	}
	return nil
}

func (t *TestState) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("TestState: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.A (uint64) (uint64)
		case "A":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.A = uint64(extra)
			// t.B (uint64) (uint64)
		case "B":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.B = uint64(extra)

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *TestEvent) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	// t.A (string) (string)
	if len("A") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"A\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("A")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("A")); err != nil {
		return err
	}

	if len(t.A) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.A was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.A)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.A)); err != nil {
		return err
	}

	// t.Val (uint64) (uint64)
	if len("Val") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Val\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Val")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Val")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Val))); err != nil {
		return err
	}
	return nil
}

func (t *TestEvent) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("TestEvent: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.A (string) (string)
		case "A":

			{
				sval, err := cbg.ReadString(br)
				if err != nil {
					return err
				}

				t.A = string(sval)
			}
			// t.Val (uint64) (uint64)
		case "Val":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.Val = uint64(extra)

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
