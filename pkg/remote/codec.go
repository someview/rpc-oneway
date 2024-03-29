/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remote

import (
	"context"
)

// Codec is the abstraction of the codec layer of Kitex.
type Codec interface {
	Encode(ctx context.Context, msg Message, out ByteBuffer) error
	Decode(ctx context.Context, msg Message, in ByteBuffer) error
	Name() string
}

// MetaEncoder is an abstraction of the encode layer that has meta and payload stage.
// Users could implement EncodeMetaAndPayload to customize header encoding logic,
// or implement EncodePayload and pass the overwritten implementation into EncodeMetaAndPayload
// to customize payload encoding logic.
type MetaEncoder interface {
	EncodeMetaAndPayload(ctx context.Context, msg Message, out ByteBuffer, me MetaEncoder) error
	EncodePayload(ctx context.Context, msg Message, out ByteBuffer) error
}

// MetaDecoder is an abstraction of the decode layer that has meta and payload stage
type MetaDecoder interface {
	DecodeMeta(ctx context.Context, msg Message, in ByteBuffer) error
	DecodePayload(ctx context.Context, msg Message, in ByteBuffer) error
}
