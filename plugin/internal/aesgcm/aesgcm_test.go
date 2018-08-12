// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aesgcm

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key, err := Key("xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh")
	if err != nil {
		t.Error(err)
		return
	}

	message := []byte("top-secret")
	ciphertext, err := Encrypt(message, key)
	if err != nil {
		t.Error(err)
		return
	}

	plaintext, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(message, plaintext) {
		t.Errorf("Expect secret encrypted and decrypted")
	}
}

func TestInvalidKey(t *testing.T) {
	_, err := Key("xVKAGlWQiY3sOp8J")
	if err != errInvalidKeyLength {
		t.Errorf("Want Invalid Key Length error")
	}
}
