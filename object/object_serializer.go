package object

import (
	"encoding/gob"
	"fmt"

	"github.com/nggenius/ngengine/utils"

	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
)

type encoder interface {
	Serialize(ar *utils.StoreArchive) error
}

type decoder interface {
	Deserialize(ar *utils.LoadArchive) error
}

func (f *Factory) Encode(o interface{}) ([]byte, error) {
	obj := o.(Object)

	msg := protocol.NewMessage(share.MAX_BUF_LEN)
	ar := utils.NewStoreArchiver(msg.Body)
	ar.PutString(obj.ObjectType())
	if enc, ok := o.(encoder); ok {
		err := enc.Serialize(ar)
		if err != nil {
			return nil, err
		}

		return ar.Data(), nil
	}

	enc := gob.NewEncoder(ar)
	err := enc.Encode(o)
	if err != nil {
		return nil, err
	}

	return ar.Data(), nil
}

func (f *Factory) Sync(o Object, b []byte) error {
	ar := utils.NewLoadArchiver(b)
	typ, err := ar.GetString()
	if err != nil {
		return err
	}

	if o.ObjectType() != typ {
		return fmt.Errorf("object type error, src: %s, have: %s", o.ObjectType(), typ)
	}

	if dec, ok := o.(decoder); ok {
		return dec.Deserialize(ar)
	}

	dec := gob.NewDecoder(ar)
	err = dec.Decode(o)
	if err != nil {
		return err
	}
	return nil
}

func (f *Factory) Decode(b []byte) (interface{}, error) {
	ar := utils.NewLoadArchiver(b)
	typ, err := ar.GetString()
	if err != nil {
		return nil, err
	}

	o, err := f.Create(typ)
	if err != nil {
		return nil, err
	}

	if dec, ok := o.(decoder); ok {
		if err := dec.Deserialize(ar); err != nil {
			f.Destroy(o)
			return nil, err
		}
		return o, nil
	}

	dec := gob.NewDecoder(ar)
	err = dec.Decode(o)
	if err != nil {
		f.Destroy(o)
		return nil, err
	}
	return o, nil
}
