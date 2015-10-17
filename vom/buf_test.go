// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vom

// TODO(toddw): Add more exhaustive tests of decbuf SetLimit and fill.

import (
	"io"
	"strings"
	"testing"
)

func expectEncbufBytes(t *testing.T, b *encbuf, expect string) {
	if got, want := b.Len(), len(expect); got != want {
		t.Errorf("len got %d, want %d", got, want)
	}
	if got, want := string(b.Bytes()), expect; got != want {
		t.Errorf("bytes got %q, want %q", b.Bytes(), expect)
	}
}

func TestEncbuf(t *testing.T) {
	b := newEncbuf(2 << 14)
	expectEncbufBytes(t, b, "")
	copy(b.Grow(5), "abcde*****")
	expectEncbufBytes(t, b, "abcde")
	b.WriteOneByte('1')
	expectEncbufBytes(t, b, "abcde1")
	b.WriteMaximumPossible([]byte("def"))
	expectEncbufBytes(t, b, "abcde1def")
	b.Reset()
	expectEncbufBytes(t, b, "")
	b.WriteMaximumPossible([]byte("123"))
	expectEncbufBytes(t, b, "123")
}

func testEncbufReserve(t *testing.T, base int) {
	for _, size := range []int{base - 1, base, base + 1} {
		str := strings.Repeat("x", size)
		// Test starting empty and writing size.
		b := newEncbuf(2 << 14)
		b.WriteMaximumPossible([]byte(str))
		expectEncbufBytes(t, b, str)
		b.WriteMaximumPossible([]byte(str))
		expectEncbufBytes(t, b, str+str)
		// Test starting with one byte and writing size.
		b = newEncbuf(2 << 14)
		b.WriteOneByte('A')
		expectEncbufBytes(t, b, "A")
		b.WriteMaximumPossible([]byte(str))
		expectEncbufBytes(t, b, "A"+str)
		b.WriteMaximumPossible([]byte(str))
		expectEncbufBytes(t, b, "A"+str+str)
	}
}

func TestEncbufReserve(t *testing.T) {
	testEncbufReserve(t, minBufFree)
	testEncbufReserve(t, minBufFree*2)
	testEncbufReserve(t, minBufFree*3)
	testEncbufReserve(t, minBufFree*4)
}

func TestEncbufHitMaximum(t *testing.T) {
	maxSize := 2 << 10
	b := newEncbuf(maxSize)
	b.WriteOneByte('A')
	b.WriteOneByte('B')
	str := strings.Repeat("x", maxSize)
	amtWritten := b.WriteMaximumPossible([]byte(str))
	if got, want := amtWritten, maxSize-2; got != want {
		t.Errorf("expected %d bytes to be written, but %d were", want, got)
	}
	expectEncbufBytes(t, b, "AB"+str[:len(str)-2])
}

func expectReadSmall(t *testing.T, mode string, b *decbuf, n int, expect string, expectErr error) {
	buf, err := b.ReadSmall(n)
	if err != nil {
		if got, want := err, expectErr; got != want {
			t.Errorf("%s ReadBuf err got %v, want %v", mode, got, want)
		}
		return
	}
	if got, want := string(buf)[:n], expect; got != want {
		t.Errorf("%s ReadBuf buf got %q, want %q", mode, got, want)
	}
}

func expectSkip(t *testing.T, mode string, b *decbuf, n, expectN int, expectErr error) {
	nGot, err := b.Skip(n)
	if got, want := err, expectErr; got != want {
		t.Errorf("%s Skip err got %v, want %v", mode, got, want)
	} else if got, want := nGot, expectN; got != want {
		t.Errorf("%s Skip %d bytes only skipped %d", mode, want, got)
	}
}

func expectPeekSmall(t *testing.T, mode string, b *decbuf, n int, expect string, expectErr error) {
	buf, err := b.PeekSmall(n)
	if err != nil {
		if got, want := err, expectErr; got != want {
			t.Errorf("%s PeakSmall err got %v, want %v", mode, got, want)
		}
		return
	}
	if got, want := string(buf)[:n], expect; got != want {
		t.Errorf("%s PeakSmall buf got %q, want %q", mode, got, want)
	}
}

func expectReadByte(t *testing.T, mode string, b *decbuf, expect byte, expectErr error) {
	actual, err := b.ReadByte()
	if got, want := err, expectErr; got != want {
		t.Errorf("%s ReadByte err got %v, want %v", mode, got, want)
	}
	if got, want := actual, expect; got != want {
		t.Errorf("%s ReadByte buf got %q, want %q", mode, got, want)
	}
}

func expectPeekByte(t *testing.T, mode string, b *decbuf, expect byte, expectErr error) {
	actual, err := b.PeekByte()
	if got, want := err, expectErr; got != want {
		t.Errorf("%s PeekByte err got %v, want %v", mode, got, want)
	}
	if got, want := actual, expect; got != want {
		t.Errorf("%s PeekByte buf got %q, want %q", mode, got, want)
	}
}

func expectReadIntoBuf(t *testing.T, mode string, b *decbuf, n, expectN int, expect string, expectErr error) {
	buf := make([]byte, n)
	nGot, err := b.ReadIntoBuf(buf)
	if got, want := err, expectErr; got != want {
		t.Errorf("%s ReadIntoBuf err got %v, want %v", mode, got, want)
	} else if got, want := nGot, expectN; got != want {
		t.Errorf("%s ReadIntoBuf read %d but expected %d", mode, got, want)
	}
	if err == nil {
		if got, want := string(buf), expect; got != want {
			t.Errorf("%s ReadIntoBuf buf got %q, want %q", mode, got, want)
		}
	}
}

func TestDecbufReadSmall(t *testing.T) {
	fn := func(mode string, b *decbuf) {
		expectReadSmall(t, mode, b, 1, "a", nil)
		expectReadSmall(t, mode, b, 2, "bc", nil)
		expectReadSmall(t, mode, b, 3, "def", nil)
		expectReadSmall(t, mode, b, 4, "ghij", nil)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
	}
	for _, mode := range allReadModes {
		fn(mode.String(), newDecbuf(mode.testReader(abcReader(10))))
	}
	fn("fromBytes", newDecbufFromBytes(abcBytes(10)))
}

func TestDecbufSkip(t *testing.T) {
	fn := func(mode string, b *decbuf) {
		expectSkip(t, mode, b, 1, 1, nil)
		expectReadSmall(t, mode, b, 2, "bc", nil)
		expectSkip(t, mode, b, 3, 3, nil)
		expectReadSmall(t, mode, b, 2, "gh", nil)
		expectSkip(t, mode, b, 2, 2, nil)
		expectSkip(t, mode, b, 1, 0, io.EOF)
		expectSkip(t, mode, b, 1, 0, io.EOF)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
	}
	for _, mode := range allReadModes {
		fn(mode.String(), newDecbuf(mode.testReader(abcReader(10))))
	}
	fn("fromBytes", newDecbufFromBytes(abcBytes(10)))
}

func TestDecbufPeekSmall(t *testing.T) {
	fn1 := func(mode string, b *decbuf) {
		// Start peeking at beginning.
		expectPeekSmall(t, mode, b, 1, "a", nil)
		expectPeekSmall(t, mode, b, 2, "ab", nil)
		expectPeekSmall(t, mode, b, 3, "abc", nil)
		expectPeekSmall(t, mode, b, 4, "", io.EOF)
		expectReadSmall(t, mode, b, 3, "abc", nil)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectPeekSmall(t, mode, b, 1, "", io.EOF)
	}
	fn2 := func(mode string, b *decbuf) {
		// Start peeking after reading 1 byte, which fills the buffer
		expectReadSmall(t, mode, b, 1, "a", nil)
		expectPeekSmall(t, mode, b, 1, "b", nil)
		expectPeekSmall(t, mode, b, 2, "bc", nil)
		expectPeekSmall(t, mode, b, 3, "bcd", nil)
		expectPeekSmall(t, mode, b, 4, "", io.EOF)
		expectReadSmall(t, mode, b, 3, "bcd", nil)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectPeekSmall(t, mode, b, 1, "", io.EOF)
	}
	for _, mode := range []readMode{readAll, readHalf, readAllEOF, readHalfEOF} {
		fn1(mode.String(), newDecbuf(mode.testReader(abcReader(3))))
		fn2(mode.String(), newDecbuf(mode.testReader(abcReader(4))))
	}
	fn1("fromBytes", newDecbufFromBytes(abcBytes(3)))
	fn2("fromBytes", newDecbufFromBytes(abcBytes(4)))
}

func TestDecbufPeekSmall1(t *testing.T) {
	fn1 := func(mode string, b *decbuf) {
		// Start peeking at beginning.
		expectPeekSmall(t, mode, b, 1, "a", nil)
		expectPeekSmall(t, mode, b, 2, "ab", nil)
		expectPeekSmall(t, mode, b, 3, "abc", nil)
		expectPeekSmall(t, mode, b, 2, "ab", nil)
		expectPeekSmall(t, mode, b, 1, "a", nil)
		expectPeekSmall(t, mode, b, 4, "", io.EOF)
		expectReadSmall(t, mode, b, 3, "abc", nil)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectPeekSmall(t, mode, b, 1, "", io.EOF)
	}
	fn2 := func(mode string, b *decbuf) {
		// Start peeking after reading 1 byte, which fills the buffer
		expectReadSmall(t, mode, b, 1, "a", nil)
		expectPeekSmall(t, mode, b, 1, "b", nil)
		expectPeekSmall(t, mode, b, 2, "bc", nil)
		expectPeekSmall(t, mode, b, 3, "bcd", nil)
		expectPeekSmall(t, mode, b, 2, "bc", nil)
		expectPeekSmall(t, mode, b, 1, "b", nil)
		expectPeekSmall(t, mode, b, 4, "", io.EOF)
		expectReadSmall(t, mode, b, 3, "bcd", nil)
		expectReadSmall(t, mode, b, 1, "", io.EOF)
		expectPeekSmall(t, mode, b, 1, "", io.EOF)
	}
	for _, mode := range []readMode{readOneByte, readOneByteEOF} {
		fn1(mode.String(), newDecbuf(mode.testReader(abcReader(3))))
		fn2(mode.String(), newDecbuf(mode.testReader(abcReader(4))))
	}
	// Don't try newDecbufFromBytes for this test, since it requires filling the
	// buffer one byte at a time.
}

func TestDecbufReadByte(t *testing.T) {
	fn := func(mode string, b *decbuf) {
		expectReadByte(t, mode, b, 'a', nil)
		expectReadByte(t, mode, b, 'b', nil)
		expectReadByte(t, mode, b, 'c', nil)
		expectReadByte(t, mode, b, 0, io.EOF)
		expectReadByte(t, mode, b, 0, io.EOF)
	}
	for _, mode := range allReadModes {
		fn(mode.String(), newDecbuf(mode.testReader(abcReader(3))))
	}
	fn("fromBytes", newDecbufFromBytes(abcBytes(3)))
}

func TestDecbufPeekByte(t *testing.T) {
	fn := func(mode string, b *decbuf) {
		expectPeekByte(t, mode, b, 'a', nil)
		expectPeekByte(t, mode, b, 'a', nil)
		expectReadByte(t, mode, b, 'a', nil)

		expectPeekByte(t, mode, b, 'b', nil)
		expectPeekByte(t, mode, b, 'b', nil)
		expectReadByte(t, mode, b, 'b', nil)

		expectPeekByte(t, mode, b, 'c', nil)
		expectPeekByte(t, mode, b, 'c', nil)
		expectReadByte(t, mode, b, 'c', nil)

		expectPeekByte(t, mode, b, 0, io.EOF)
		expectPeekByte(t, mode, b, 0, io.EOF)
		expectReadByte(t, mode, b, 0, io.EOF)
	}
	for _, mode := range allReadModes {
		fn(mode.String(), newDecbuf(mode.testReader(abcReader(3))))
	}
	fn("fromBytes", newDecbufFromBytes(abcBytes(3)))
}

func TestDecbufReadIntBuf(t *testing.T) {
	fn1 := func(mode string, b *decbuf) {
		// Start ReadFull from beginning.
		expectReadIntoBuf(t, mode, b, 3, 3, "abc", nil)
		expectReadIntoBuf(t, mode, b, 3, 3, "def", nil)
		expectReadIntoBuf(t, mode, b, 1, 0, "", io.EOF)
		expectReadIntoBuf(t, mode, b, 1, 0, "", io.EOF)
	}
	fn2 := func(mode string, b *decbuf) {
		// Start ReadFull after reading 1 byte, which fills the buffer.
		expectReadSmall(t, mode, b, 1, "a", nil)
		expectReadIntoBuf(t, mode, b, 2, 2, "bc", nil)
		expectReadIntoBuf(t, mode, b, 3, 3, "def", nil)
		expectReadIntoBuf(t, mode, b, 1, 0, "", io.EOF)
		expectReadIntoBuf(t, mode, b, 1, 0, "", io.EOF)
	}
	for _, mode := range allReadModes {
		fn1(mode.String(), newDecbuf(mode.testReader(abcReader(6))))
		fn2(mode.String(), newDecbuf(mode.testReader(abcReader(6))))
	}
	fn1("fromBytes", newDecbufFromBytes(abcBytes(6)))
	fn2("fromBytes", newDecbufFromBytes(abcBytes(6)))
}
