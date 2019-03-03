// Code generated by codecgen - DO NOT EDIT.

package client

import (
	"errors"
	codec1978 "github.com/ugorji/go/codec"
	"runtime"
	"strconv"
	"time"
)

const (
	// ----- content types ----
	codecSelferCcUTF81819 = 1
	codecSelferCcRAW1819  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1819  = 10
	codecSelferValueTypeMap1819    = 9
	codecSelferValueTypeString1819 = 6
	codecSelferValueTypeInt1819    = 2
	codecSelferValueTypeUint1819   = 3
	codecSelferValueTypeFloat1819  = 4
	codecSelferBitsize1819         = uint8(32 << (^uint(0) >> 63))
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct1819 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1819 struct{}

func init() {
	if codec1978.GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		panic("codecgen version mismatch: current: 8, need " + strconv.FormatInt(int64(codec1978.GenVersion), 10) + ". Re-generate file: " + file)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 time.Time
		_ = v0
	}
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false // struct tag has 'toArray'
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				r.WriteMapStart(3)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Action))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `action`)
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Action))
				}
			}
			var yyn6 bool
			if x.Node == nil {
				yyn6 = true
				goto LABEL6
			}
		LABEL6:
			if yyr2 || yy2arr2 {
				if yyn6 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.Node == nil {
						r.EncodeNil()
					} else {
						x.Node.CodecEncodeSelf(e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `node`)
				r.WriteMapElemValue()
				if yyn6 {
					r.EncodeNil()
				} else {
					if x.Node == nil {
						r.EncodeNil()
					} else {
						x.Node.CodecEncodeSelf(e)
					}
				}
			}
			var yyn9 bool
			if x.PrevNode == nil {
				yyn9 = true
				goto LABEL9
			}
		LABEL9:
			if yyr2 || yy2arr2 {
				if yyn9 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.PrevNode == nil {
						r.EncodeNil()
					} else {
						x.PrevNode.CodecEncodeSelf(e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `prevNode`)
				r.WriteMapElemValue()
				if yyn9 {
					r.EncodeNil()
				} else {
					if x.PrevNode == nil {
						r.EncodeNil()
					} else {
						x.PrevNode.CodecEncodeSelf(e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1819 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1819 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1819)
		}
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = (string)(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if true && x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}

				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if true && x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}

				x.PrevNode.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj7 int
	var yyb7 bool
	var yyhl7 bool = l >= 0
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = (string)(r.DecodeString())
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}

		x.Node.CodecDecodeSelf(d)
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}

		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj7++
		if yyhl7 {
			yyb7 = yyj7 > l
		} else {
			yyb7 = r.CheckBreak()
		}
		if yyb7 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj7-1, "")
	}
	r.ReadArrayEnd()
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false // struct tag has 'toArray'
			var yyq2 = [8]bool{     // should field at this index be written?
				true,                // Key
				x.Dir,               // Dir
				true,                // Value
				true,                // Nodes
				true,                // CreatedIndex
				true,                // ModifiedIndex
				x.Expiration != nil, // Expiration
				x.TTL != 0,          // TTL
			}
			_ = yyq2
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(8)
			} else {
				var yynn2 int
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `key`)
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferCcUTF81819, `dir`)
					r.WriteMapElemValue()
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `value`)
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF81819, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `nodes`)
				r.WriteMapElemValue()
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `createdIndex`)
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferCcUTF81819, `modifiedIndex`)
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			}
			var yyn21 bool
			if x.Expiration == nil {
				yyn21 = true
				goto LABEL21
			}
		LABEL21:
			if yyr2 || yy2arr2 {
				if yyn21 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if yyq2[6] {
						if x.Expiration == nil {
							r.EncodeNil()
						} else {
							yy22 := *x.Expiration
							if false {
							} else {
								r.EncodeTime(yy22)
							}
						}
					} else {
						r.EncodeNil()
					}
				}
			} else {
				if yyq2[6] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferCcUTF81819, `expiration`)
					r.WriteMapElemValue()
					if yyn21 {
						r.EncodeNil()
					} else {
						if x.Expiration == nil {
							r.EncodeNil()
						} else {
							yy24 := *x.Expiration
							if false {
							} else {
								r.EncodeTime(yy24)
							}
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[7] {
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[7] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferCcUTF81819, `ttl`)
					r.WriteMapElemValue()
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1819 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1819 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1819)
		}
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = (string)(r.DecodeString())
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				x.Nodes.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = (uint64)(r.DecodeUint64())
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = (uint64)(r.DecodeUint64())
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if true && x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}

				if false {
				} else {
					*x.Expiration = r.DecodeTime()
				}
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = (int64)(r.DecodeInt64())
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj13 int
	var yyb13 bool
	var yyhl13 bool = l >= 0
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = (bool)(r.DecodeBool())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = (string)(r.DecodeString())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		x.Nodes.CodecDecodeSelf(d)
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = (uint64)(r.DecodeUint64())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = (uint64)(r.DecodeUint64())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}

		if false {
		} else {
			*x.Expiration = r.DecodeTime()
		}
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = (int64)(r.DecodeInt64())
	}
	for {
		yyj13++
		if yyhl13 {
			yyb13 = yyj13 > l
		} else {
			yyb13 = r.CheckBreak()
		}
		if yyb13 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj13-1, "")
	}
	r.ReadArrayEnd()
}

func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			h.encNodes((Nodes)(x), e)
		}
	}
}

func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		h.decNodes((*Nodes)(x), d)
	}
}

func (x codecSelfer1819) encNodes(v Nodes, e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteArrayStart(len(v))
	for _, yyv1 := range v {
		r.WriteArrayElem()
		if yyv1 == nil {
			r.EncodeNil()
		} else {
			yyv1.CodecEncodeSelf(e)
		}
	}
	r.WriteArrayEnd()
}

func (x codecSelfer1819) decNodes(v *Nodes, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []*Node{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 8)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]*Node, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		// var yydn1 bool
		for ; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 8)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]*Node, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)

			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, nil)
				yyc1 = true

			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = nil
				} else {
					if yyv1[yyj1] == nil {
						yyv1[yyj1] = new(Node)
					}
					yyv1[yyj1].CodecDecodeSelf(d)
				}

			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]*Node, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
