package example_generates

import (
	"bytes"
	externallib "github.com/rgnu/tinyjson/example_generates/external_lib"
	lexer "github.com/tinyjson/lexer"
	"strconv"
)

func tinyjsonMarshalC39f6f78a15d523b(w *bytes.Buffer, this *externallib.ExternalClass) {
	w.WriteString("{")
	w.WriteString("\"Key\":")
	w.WriteString(strconv.Quote(string(this.Key)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC39f6f78a15d523b(lex *lexer.Lexer, this *externallib.ExternalClass) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key":
				v105, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key = (string)(v105)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func tinyjsonMarshalC55104dc76695721d(w *bytes.Buffer, this *AnonymousAnonymousStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.SimpleStruct.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.SimpleStruct.Key2)))
	w.WriteString(",")
	w.WriteString("\"Key3\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.Key3)))
	w.WriteString(",")
	w.WriteString("\"Key4\":")
	w.WriteString(strconv.Quote(string(this.Key4)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC55104dc76695721d(lex *lexer.Lexer, this *AnonymousAnonymousStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v3, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.SimpleStruct.Key1 = (string)(v3)

			case "Key2":
				v5, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.SimpleStruct.Key2 = (string)(v5)

			case "Key3":
				v7, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.Key3 = (string)(v7)

			case "Key4":
				v9, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key4 = (string)(v9)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *AnonymousAnonymousStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC55104dc76695721d(w, this)
	return w.Bytes(), nil
}

func (this *AnonymousAnonymousStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC55104dc76695721d(lex, this)
}

func tinyjsonMarshalC6054502fc5d6d268(w *bytes.Buffer, this *AnonymousStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.SimpleStruct.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.SimpleStruct.Key2)))
	w.WriteString(",")
	w.WriteString("\"Key3\":")
	w.WriteString(strconv.Quote(string(this.Key3)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC6054502fc5d6d268(lex *lexer.Lexer, this *AnonymousStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v58, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.SimpleStruct.Key1 = (string)(v58)

			case "Key2":
				v60, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.SimpleStruct.Key2 = (string)(v60)

			case "Key3":
				v62, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key3 = (string)(v62)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *AnonymousStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6054502fc5d6d268(w, this)
	return w.Bytes(), nil
}

func (this *AnonymousStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6054502fc5d6d268(lex, this)
}

func tinyjsonMarshalC2e3108dabb158644(w *bytes.Buffer, this *DoubleArray) {
	if *this == nil {
		w.WriteString("null")
	} else {
		w.WriteString("[")
		for v107, v108 := range *this {
			if v107 > 0 {
				w.WriteString(",")
			}
			if v108 == nil {
				w.WriteString("null")
			} else {
				w.WriteString("[")
				for v109, v110 := range v108 {
					if v109 > 0 {
						w.WriteString(",")
					}
					w.WriteString(strconv.Quote(string(v110)))
				}
				w.WriteString("]")
			}
		}
		w.WriteString("]")
	}
}

func tinyjsonUnmarshalC2e3108dabb158644(lex *lexer.Lexer, this *DoubleArray) error {
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ArrayIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	} else {
		*this = make([][]string, 0)
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		for {
			if lex.Controls[0] == lexer.ArrayOut {
				lex.Controls = lex.Controls[1:]

				break
			}
			var v112 []string
			if lex.Controls[0] == lexer.Nil {
				v112 = nil
				lex.Controls = lex.Controls[1:]
			} else if lex.Controls[0] != lexer.ArrayIn {
				lex.SkipValue()
				return lexer.ErrorUnexpectedType
			} else {
				v112 = make([]string, 0)
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				for {
					if lex.Controls[0] == lexer.ArrayOut {
						lex.Controls = lex.Controls[1:]

						break
					}
					v114, err := lex.ReadString()
					if err != nil {
						return err
					}

					v112 = append(v112, string(v114))
				}
			}

			*this = append(*this, v112)
		}
	}

	return nil
}

func (this *DoubleArray) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2e3108dabb158644(w, this)
	return w.Bytes(), nil
}

func (this *DoubleArray) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2e3108dabb158644(lex, this)
}

func tinyjsonMarshalC30b95ff183c471d4(w *bytes.Buffer, this *IntPtr) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	v125 := *this
	w.WriteString(strconv.FormatInt(int64(*v125), 10))
}

func tinyjsonUnmarshalC30b95ff183c471d4(lex *lexer.Lexer, this *IntPtr) error {
	v127, err := lex.ReadInt()
	if err != nil {
		return err
	}
	if err == lexer.ErrorNilValue {
		err = nil
		*this = nil
		return nil
	}
	v128 := int(v127)
	v129 := &v128
	*this = IntPtr(v129)

	return nil
}

func tinyjsonMarshalC1b6cffa2ba517936(w *bytes.Buffer, this *IntPtrPtr) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	v71 := *this
	if *v71 == nil {
		w.WriteString("null")
		return
	}
	v72 := *v71
	w.WriteString(strconv.FormatInt(int64(*v72), 10))
}

func tinyjsonUnmarshalC1b6cffa2ba517936(lex *lexer.Lexer, this *IntPtrPtr) error {
	v74, err := lex.ReadInt()
	if err != nil {
		return err
	}
	if err == lexer.ErrorNilValue {
		err = nil
		*this = nil
		return nil
	}
	v75 := int(v74)
	v76 := &v75
	v77 := &v76
	*this = IntPtrPtr(v77)

	return nil
}

func tinyjsonMarshalC430c8b35bb9457d8(w *bytes.Buffer, this *MapFloatString) {
	if *this == nil {
		w.WriteString("null")
	} else {
		var v30 bool
		w.WriteString("{")
		for v28, v29 := range *this {
			if v30 {
				w.WriteString(",")
			} else {
				v30 = true
			}
			w.WriteString("\"" + strconv.FormatFloat(float64(v28), 'g', -1, 64) + "\"")
			w.WriteString(":")
			w.WriteString(strconv.Quote(string(v29)))
		}
		w.WriteString("}")
	}
}

func tinyjsonUnmarshalC430c8b35bb9457d8(lex *lexer.Lexer, this *MapFloatString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[float64]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v32, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			v33, _ := strconv.ParseFloat(v32, 64)
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v35, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[float64(v33)] = string(v35)
		}
	}
	return nil
}

func (this *MapFloatString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC430c8b35bb9457d8(w, this)
	return w.Bytes(), nil
}

func (this *MapFloatString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC430c8b35bb9457d8(lex, this)
}

func tinyjsonMarshalC25845c95d4491d1b(w *bytes.Buffer, this *MapIntString) {
	if *this == nil {
		w.WriteString("null")
	} else {
		var v12 bool
		w.WriteString("{")
		for v10, v11 := range *this {
			if v12 {
				w.WriteString(",")
			} else {
				v12 = true
			}
			w.WriteString("\"" + strconv.FormatInt(int64(v10), 10) + "\"")
			w.WriteString(":")
			w.WriteString(strconv.Quote(string(v11)))
		}
		w.WriteString("}")
	}
}

func tinyjsonUnmarshalC25845c95d4491d1b(lex *lexer.Lexer, this *MapIntString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[int]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v14, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			v15, _ := strconv.ParseInt(v14, 10, 64)
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v17, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[int(v15)] = string(v17)
		}
	}
	return nil
}

func (this *MapIntString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC25845c95d4491d1b(w, this)
	return w.Bytes(), nil
}

func (this *MapIntString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC25845c95d4491d1b(lex, this)
}

func tinyjsonMarshalC490bd268b68e6a3f(w *bytes.Buffer, this *MapMap) {
	if *this == nil {
		w.WriteString("null")
	} else {
		var v85 bool
		w.WriteString("{")
		for v83, v84 := range *this {
			if v85 {
				w.WriteString(",")
			} else {
				v85 = true
			}
			w.WriteString(strconv.Quote(string(v83)))
			w.WriteString(":")
			if v84 == nil {
				w.WriteString("null")
			} else {
				var v88 bool
				w.WriteString("{")
				for v86, v87 := range v84 {
					if v88 {
						w.WriteString(",")
					} else {
						v88 = true
					}
					w.WriteString(strconv.Quote(string(v86)))
					w.WriteString(":")
					w.WriteString(strconv.Quote(string(v87)))
				}
				w.WriteString("}")
			}
		}
		w.WriteString("}")
	}
}

func tinyjsonUnmarshalC490bd268b68e6a3f(lex *lexer.Lexer, this *MapMap) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[string]map[string]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v90, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v91 := map[string]string{}
			data := lex.Data()
			if lex.Controls[0] == lexer.Nil {
				v91 = nil
				lex.Controls = lex.Controls[1:]
			} else if lex.Controls[0] != lexer.ObjectIn {
				lex.SkipValue()
			} else {
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				v91 = make(map[string]string, 0)
				for {
					if lex.Controls[0] == lexer.ObjectOut {
						lex.Controls = lex.Controls[1:]
						break
					}
					v92, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
					lex.Controls = lex.Controls[1:]
					lex.Actions = lex.Actions[2:]
					v94, err := lex.ReadString()
					if err != nil {
						return err
					}

					v91[v92] = string(v94)
				}
			}
			(*this)[v90] = v91
		}
	}
	return nil
}

func (this *MapMap) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC490bd268b68e6a3f(w, this)
	return w.Bytes(), nil
}

func (this *MapMap) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC490bd268b68e6a3f(lex, this)
}

func tinyjsonMarshalC78629a0f5f3f164f(w *bytes.Buffer, this *MapStringString) {
	if *this == nil {
		w.WriteString("null")
	} else {
		var v117 bool
		w.WriteString("{")
		for v115, v116 := range *this {
			if v117 {
				w.WriteString(",")
			} else {
				v117 = true
			}
			w.WriteString(strconv.Quote(string(v115)))
			w.WriteString(":")
			w.WriteString(strconv.Quote(string(v116)))
		}
		w.WriteString("}")
	}
}

func tinyjsonUnmarshalC78629a0f5f3f164f(lex *lexer.Lexer, this *MapStringString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[string]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v119, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v121, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[v119] = string(v121)
		}
	}
	return nil
}

func (this *MapStringString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC78629a0f5f3f164f(w, this)
	return w.Bytes(), nil
}

func (this *MapStringString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC78629a0f5f3f164f(lex, this)
}

func tinyjsonMarshalC866cb397916001e(w *bytes.Buffer, this *MyArray) {
	if *this == nil {
		w.WriteString("null")
	} else {
		w.WriteString("[")
		for v36, v37 := range *this {
			if v36 > 0 {
				w.WriteString(",")
			}
			w.WriteString(strconv.Quote(string(v37)))
		}
		w.WriteString("]")
	}
}

func tinyjsonUnmarshalC866cb397916001e(lex *lexer.Lexer, this *MyArray) error {
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ArrayIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	} else {
		*this = make([]string, 0)
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		for {
			if lex.Controls[0] == lexer.ArrayOut {
				lex.Controls = lex.Controls[1:]

				break
			}
			v40, err := lex.ReadString()
			if err != nil {
				return err
			}

			*this = append(*this, string(v40))
		}
	}

	return nil
}

func (this *MyArray) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC866cb397916001e(w, this)
	return w.Bytes(), nil
}

func (this *MyArray) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC866cb397916001e(lex, this)
}

func tinyjsonMarshalC57e9d1860d1d68d8(w *bytes.Buffer, this *MyExternal) {
	tinyjsonMarshalC39f6f78a15d523b(w, (*externallib.ExternalClass)(this))
}

func tinyjsonUnmarshalC57e9d1860d1d68d8(lex *lexer.Lexer, this *MyExternal) error {
	tinyjsonUnmarshalC39f6f78a15d523b(lex, (*externallib.ExternalClass)(this))
	return nil
}

func (this *MyExternal) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC57e9d1860d1d68d8(w, this)
	return w.Bytes(), nil
}

func (this *MyExternal) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC57e9d1860d1d68d8(lex, this)
}

func tinyjsonMarshalC380704bb7b4d7c03(w *bytes.Buffer, this *MyExternalAlias) {
	tinyjsonMarshalC39f6f78a15d523b(w, (*externallib.ExternalClass)(this))
}

func tinyjsonUnmarshalC380704bb7b4d7c03(lex *lexer.Lexer, this *MyExternalAlias) error {
	tinyjsonUnmarshalC39f6f78a15d523b(lex, (*externallib.ExternalClass)(this))
	return nil
}

func (this *MyExternalAlias) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC380704bb7b4d7c03(w, this)
	return w.Bytes(), nil
}

func (this *MyExternalAlias) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC380704bb7b4d7c03(lex, this)
}

func tinyjsonMarshalC2606cd2b57d29245(w *bytes.Buffer, this *MyFloat) {
	w.WriteString(strconv.FormatFloat(float64(*this), 'g', -1, 64))
}

func tinyjsonUnmarshalC2606cd2b57d29245(lex *lexer.Lexer, this *MyFloat) error {
	v135, err := lex.ReadFloat()
	if err != nil {
		return err
	}
	*this = MyFloat(v135)

	return nil
}

func (this *MyFloat) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2606cd2b57d29245(w, this)
	return w.Bytes(), nil
}

func (this *MyFloat) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2606cd2b57d29245(lex, this)
}

func tinyjsonMarshalC4d65822107fcfd52(w *bytes.Buffer, this *MyInt) {
	w.WriteString(strconv.FormatInt(int64(*this), 10))
}

func tinyjsonUnmarshalC4d65822107fcfd52(lex *lexer.Lexer, this *MyInt) error {
	v101, err := lex.ReadInt()
	if err != nil {
		return err
	}
	*this = MyInt(v101)

	return nil
}

func (this *MyInt) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC4d65822107fcfd52(w, this)
	return w.Bytes(), nil
}

func (this *MyInt) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC4d65822107fcfd52(lex, this)
}

func tinyjsonMarshalC365a858149c6e2d1(w *bytes.Buffer, this *MyString) {
	w.WriteString(strconv.Quote(string(*this)))
}

func tinyjsonUnmarshalC365a858149c6e2d1(lex *lexer.Lexer, this *MyString) error {
	v64, err := lex.ReadString()
	if err != nil {
		return err
	}
	*this = (MyString)(v64)

	return nil
}

func (this *MyString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC365a858149c6e2d1(w, this)
	return w.Bytes(), nil
}

func (this *MyString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC365a858149c6e2d1(lex, this)
}

func tinyjsonMarshalC1a02070f169c1121(w *bytes.Buffer, this *MyStringAlias) {
	tinyjsonMarshalC365a858149c6e2d1(w, (*MyString)(this))
}

func tinyjsonUnmarshalC1a02070f169c1121(lex *lexer.Lexer, this *MyStringAlias) error {
	tinyjsonUnmarshalC365a858149c6e2d1(lex, (*MyString)(this))
	return nil
}

func (this *MyStringAlias) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1a02070f169c1121(w, this)
	return w.Bytes(), nil
}

func (this *MyStringAlias) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1a02070f169c1121(lex, this)
}

func tinyjsonMarshalC68255aaf95e94627(w *bytes.Buffer, this *MyUInt16) {
	w.WriteString(strconv.FormatUint(uint64(*this), 10))
}

func tinyjsonUnmarshalC68255aaf95e94627(lex *lexer.Lexer, this *MyUInt16) error {
	v124, err := lex.ReadInt()
	if err != nil {
		return err
	}
	*this = MyUInt16(v124)

	return nil
}

func (this *MyUInt16) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC68255aaf95e94627(w, this)
	return w.Bytes(), nil
}

func (this *MyUInt16) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC68255aaf95e94627(lex, this)
}

func tinyjsonMarshalC6e661e92759805f5(w *bytes.Buffer, this *SimpleStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.Key2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC6e661e92759805f5(lex *lexer.Lexer, this *SimpleStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v97, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key1 = (string)(v97)

			case "Key2":
				v99, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key2 = (string)(v99)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *SimpleStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6e661e92759805f5(w, this)
	return w.Bytes(), nil
}

func (this *SimpleStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6e661e92759805f5(lex, this)
}

func tinyjsonMarshalC2584c47f2cdf5b8a(w *bytes.Buffer, this *StructA1) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC2584c47f2cdf5b8a(lex *lexer.Lexer, this *StructA1) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v55, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v55)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA1) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2584c47f2cdf5b8a(w, this)
	return w.Bytes(), nil
}

func (this *StructA1) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2584c47f2cdf5b8a(lex, this)
}

func tinyjsonMarshalC1a714cf86b83d0e2(w *bytes.Buffer, this *StructA2) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString(",")
	w.WriteString("\"c\":")
	w.WriteString(strconv.Quote(string(this.StructC.C)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC1a714cf86b83d0e2(lex *lexer.Lexer, this *StructA2) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v20, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v20)

			case "c":
				v22, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructC.C = (string)(v22)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA2) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1a714cf86b83d0e2(w, this)
	return w.Bytes(), nil
}

func (this *StructA2) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1a714cf86b83d0e2(lex, this)
}

func tinyjsonMarshalC1408d2ac22c4d294(w *bytes.Buffer, this *StructA3) {
	w.WriteString("{")
	w.WriteString("\"a3\":")
	w.WriteString(strconv.Quote(string(this.A3)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC1408d2ac22c4d294(lex *lexer.Lexer, this *StructA3) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a3":
				v43, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A3 = (string)(v43)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA3) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1408d2ac22c4d294(w, this)
	return w.Bytes(), nil
}

func (this *StructA3) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1408d2ac22c4d294(lex, this)
}

func tinyjsonMarshalC3c04951aa42655d9(w *bytes.Buffer, this *StructA4) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString(",")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC3c04951aa42655d9(lex *lexer.Lexer, this *StructA4) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v80, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v80)

			case "d":
				v82, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v82)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA4) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC3c04951aa42655d9(w, this)
	return w.Bytes(), nil
}

func (this *StructA4) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC3c04951aa42655d9(lex, this)
}

func tinyjsonMarshalCc697f48392907a0(w *bytes.Buffer, this *StructA5) {
	w.WriteString("{")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalCc697f48392907a0(lex *lexer.Lexer, this *StructA5) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "d":
				v70, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v70)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA5) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalCc697f48392907a0(w, this)
	return w.Bytes(), nil
}

func (this *StructA5) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalCc697f48392907a0(lex, this)
}

func tinyjsonMarshalC6ec34c367674cb74(w *bytes.Buffer, this *StructA6) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.StructB.B)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC6ec34c367674cb74(lex *lexer.Lexer, this *StructA6) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v67, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructB.B = (string)(v67)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA6) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6ec34c367674cb74(w, this)
	return w.Bytes(), nil
}

func (this *StructA6) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6ec34c367674cb74(lex, this)
}

func tinyjsonMarshalC268447a4189deb99(w *bytes.Buffer, this *StructA7) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.StructD.D)))
	w.WriteString(",")
	w.WriteString("\"b\":")
	tinyjsonMarshalC592e17f7b068d9db(w, (*StructB)(&this.StructB))
	w.WriteString(",")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC268447a4189deb99(lex *lexer.Lexer, this *StructA7) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v46, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D = (string)(v46)

			case "b":
				tinyjsonUnmarshalC592e17f7b068d9db(lex, (*StructB)(&this.StructB))
			case "d":
				v49, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v49)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA7) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC268447a4189deb99(w, this)
	return w.Bytes(), nil
}

func (this *StructA7) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC268447a4189deb99(lex, this)
}

func tinyjsonMarshalC592e17f7b068d9db(w *bytes.Buffer, this *StructB) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.B)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC592e17f7b068d9db(lex *lexer.Lexer, this *StructB) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v52, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.B = (string)(v52)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func tinyjsonMarshalC56ec3f2525632186(w *bytes.Buffer, this *StructInStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString("{")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.Key1.Key2)))
	w.WriteString("}")
	w.WriteString("}")
}

func tinyjsonUnmarshalC56ec3f2525632186(lex *lexer.Lexer, this *StructInStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				data := lex.Data()
				if lex.Controls[0] == lexer.Nil {
					return nil
				} else if lex.Controls[0] != lexer.ObjectIn {
					lex.SkipValue()
					return lexer.ErrorUnexpectedType
				}
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				for {
					switch lex.Controls[0] {
					case lexer.ObjectOut:
						lex.Controls = lex.Controls[1:]
						return nil
					case lexer.Key:
						key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
						lex.Controls = lex.Controls[1:]
						lex.Actions = lex.Actions[2:]
						switch key {
						case "Key2":
							v133, err := lex.ReadString()
							if err != nil {
								return err
							}
							this.Key1.Key2 = (string)(v133)

						default:
							lex.SkipValue()
						}
					}
				}
			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructInStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC56ec3f2525632186(w, this)
	return w.Bytes(), nil
}

func (this *StructInStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC56ec3f2525632186(lex, this)
}

func tinyjsonMarshalC28b621587cb3ad0b(w *bytes.Buffer, this *TaggedStruct) {
	w.WriteString("{")
	w.WriteString("\"key_1\":")
	w.WriteString(strconv.Quote(string(this.Key1)))
	w.WriteString(",")
	w.WriteString("\"key_2\":")
	w.WriteString(strconv.Quote(string(this.Key2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC28b621587cb3ad0b(lex *lexer.Lexer, this *TaggedStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "key_1":
				v25, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key1 = (string)(v25)

			case "key_2":
				v27, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key2 = (string)(v27)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *TaggedStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC28b621587cb3ad0b(w, this)
	return w.Bytes(), nil
}

func (this *TaggedStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC28b621587cb3ad0b(lex, this)
}
