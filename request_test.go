package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

var (
	messageCreate        = []byte("\x42\x00\x78\x01\x00\x00\x01\x50\x42\x00\x77\x01\x00\x00\x00\x38\x42\x00\x69\x01\x00\x00\x00\x20\x42\x00\x6A\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x6B\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0D\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0F\x01\x00\x00\x01\x08\x42\x00\x5C\x05\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x79\x01\x00\x00\x00\xF0\x42\x00\x57\x05\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00\x42\x00\x91\x01\x00\x00\x00\xD8\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x17\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x41\x6C\x67\x6F\x72\x69\x74\x68\x6D\x00\x42\x00\x0B\x05\x00\x00\x00\x04\x00\x00\x00\x03\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x14\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x4C\x65\x6E\x67\x74\x68\x00\x00\x00\x00\x42\x00\x0B\x02\x00\x00\x00\x04\x00\x00\x00\x80\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x18\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x55\x73\x61\x67\x65\x20\x4D\x61\x73\x6B\x42\x00\x0B\x02\x00\x00\x00\x04\x00\x00\x00\x0C\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x28\x42\x00\x0A\x07\x00\x00\x00\x0C\x49\x6E\x69\x74\x69\x61\x6C\x20\x44\x61\x74\x65\x00\x00\x00\x00\x42\x00\x0B\x09\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00\x30\x39")
	messageGet           = []byte("\x42\x00\x78\x01\x00\x00\x00\x90\x42\x00\x77\x01\x00\x00\x00\x38\x42\x00\x69\x01\x00\x00\x00\x20\x42\x00\x6A\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x6B\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0D\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0F\x01\x00\x00\x00\x48\x42\x00\x5C\x05\x00\x00\x00\x04\x00\x00\x00\x0A\x00\x00\x00\x00\x42\x00\x79\x01\x00\x00\x00\x30\x42\x00\x94\x07\x00\x00\x00\x24\x34\x39\x61\x31\x63\x61\x38\x38\x2D\x36\x62\x65\x61\x2D\x34\x66\x62\x32\x2D\x62\x34\x35\x30\x2D\x37\x65\x35\x38\x38\x30\x32\x63\x33\x30\x33\x38\x00\x00\x00\x00")
	messageCreateKeyPair = []byte("\x42\x00\x78\x01\x00\x00\x01\xE8\x42\x00\x77\x01\x00\x00\x00\x38\x42\x00\x69\x01\x00\x00\x00\x20\x42\x00\x6A\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x6B\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0D\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x0F\x01\x00\x00\x01\xA0\x42\x00\x5C\x05\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00\x42\x00\x79\x01\x00\x00\x01\x88\x42\x00\x1F\x01\x00\x00\x00\x70\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x17\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x41\x6C\x67\x6F\x72\x69\x74\x68\x6D\x00\x42\x00\x0B\x05\x00\x00\x00\x04\x00\x00\x00\x04\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x14\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x4C\x65\x6E\x67\x74\x68\x00\x00\x00\x00\x42\x00\x0B\x02\x00\x00\x00\x04\x00\x00\x08\x00\x00\x00\x00\x00\x42\x00\x65\x01\x00\x00\x00\x80\x42\x00\x08\x01\x00\x00\x00\x40\x42\x00\x0A\x07\x00\x00\x00\x04\x4E\x61\x6D\x65\x00\x00\x00\x00\x42\x00\x0B\x01\x00\x00\x00\x28\x42\x00\x55\x07\x00\x00\x00\x0C\x74\x65\x73\x74\x5F\x70\x72\x69\x76\x61\x74\x65\x00\x00\x00\x00\x42\x00\x54\x05\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x18\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x55\x73\x61\x67\x65\x20\x4D\x61\x73\x6B\x42\x00\x0B\x02\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x6E\x01\x00\x00\x00\x80\x42\x00\x08\x01\x00\x00\x00\x40\x42\x00\x0A\x07\x00\x00\x00\x04\x4E\x61\x6D\x65\x00\x00\x00\x00\x42\x00\x0B\x01\x00\x00\x00\x28\x42\x00\x55\x07\x00\x00\x00\x0B\x74\x65\x73\x74\x5F\x70\x75\x62\x6C\x69\x63\x00\x00\x00\x00\x00\x42\x00\x54\x05\x00\x00\x00\x04\x00\x00\x00\x01\x00\x00\x00\x00\x42\x00\x08\x01\x00\x00\x00\x30\x42\x00\x0A\x07\x00\x00\x00\x18\x43\x72\x79\x70\x74\x6F\x67\x72\x61\x70\x68\x69\x63\x20\x55\x73\x61\x67\x65\x20\x4D\x61\x73\x6B\x42\x00\x0B\x02\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00")
)
