/*
 * Copyright 2022 Michael Graff.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dns

import "fmt"

type buffer struct {
	data    []byte
	current int
}

type WritableBuffer buffer
type ReadableBuffer buffer

func NewWritableBuffer(capacity int) *WritableBuffer {
	return &WritableBuffer{
		data:    make([]byte, capacity),
		current: 0,
	}
}

func NewReadableBuffer(data []byte) *ReadableBuffer {
	return &ReadableBuffer{
		data:    data,
		current: 0,
	}
}

func (w *WritableBuffer) GetBytes(length int) ([]byte, error) {
	if w.current+length > len(w.data) {
		return nil, fmt.Errorf("length exceeds buffer")
	}
	ret := make([]byte, 0, length)
	offset := w.current + length
	ret = append(ret, w.data[w.current:offset]...)
	w.current += length
	return ret, nil
}
