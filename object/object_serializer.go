package object

import (
	"encoding/gob"
	"fmt"

	"github.com/nggenius/ngengine/utils"

	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
)

type encoder interface {
	Serialize() ([]byte, error)
}

type decoder interface {
	Deserialize([]byte) error
}

func (f *Factory) Encode(o interface{}) ([]byte, error) {
	obj := o.(Object)

	msg := protocol.NewMessage(share.MAX_BUF_LEN)
	ar := utils.NewStoreArchiver(msg.Body)
	ar.PutString(obj.ObjectType())
	if enc, ok := o.(encoder); ok {
		b, err := enc.Serialize()
		if err != nil {
			return nil, err
		}

		err = ar.PutData(b)
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
	typ, err := ar.ReadString()
	if err != nil {
		return err
	}

	if o.ObjectType() != typ {
		return fmt.Errorf("object type error, src: %s, have: %s", o.ObjectType(), typ)
	}

	if dec, ok := o.(decoder); ok {
		data, err := ar.ReadData()
		if err != nil {
			return err
		}
		return dec.Deserialize(data)
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
	typ, err := ar.ReadString()
	if err != nil {
		return nil, err
	}

	o, err := f.Create(typ)
	if err != nil {
		return nil, err
	}

	if dec, ok := o.(decoder); ok {
		data, err := ar.ReadData()
		if err != nil {
			f.Destroy(o)
			return nil, err
		}
		return o, dec.Deserialize(data)
	}

	dec := gob.NewDecoder(ar)
	err = dec.Decode(o)
	if err != nil {
		f.Destroy(o)
		return nil, err
	}
	return o, nil
}

// func (f *Factory) decodeObj(parent Container, ar *utils.LoadArchive) (interface{}, error) {
// 	var typ string
// 	err := ar.Read(&typ)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pos := 0
// 	caps := 0
// 	childs := 0

// 	err = ar.Read(&pos)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = ar.Read(&caps)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = ar.Read(&childs)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var origin rpc.Mailbox
// 	err = ar.Read(&origin)

// 	o, err := f.Create(typ)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = ar.Read(o)
// 	if err != nil {
// 		f.Destroy(o)
// 		return nil, err
// 	}

// 	obj := o.(Object)
// 	obj.Witness().SetDummy(true)
// 	obj.Witness().SetOriginal(&origin)

// 	c, ok := o.(Container)
// 	if ok {
// 		c.SetCap(caps)

// 		if parent != nil {
// 			parent.AddChildIf(pos, o)
// 		}

// 		for i := 0; i < childs; i++ {
// 			_, err := f.decodeObj(c, ar)
// 			if err != nil {
// 				f.Destroy(o)
// 				return nil, err
// 			}
// 		}
// 	}

// 	return o, nil
// }
