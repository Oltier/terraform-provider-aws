// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by internal/generate/acctestconsts/main.go; DO NOT EDIT.

package acctest

import (
	"fmt"
)

const (
	Ct0                     = "0"
	Ct1                     = "1"
	Ct10                    = "10"
	Ct2                     = "2"
	Ct3                     = "3"
	Ct4                     = "4"
	CtBasic                 = "basic"
	CtCertificatePEM        = "certificate_pem"
	CtDisappears            = "disappears"
	CtFalse                 = "false"
	CtFalseCaps             = "FALSE"
	CtKey1                  = "key1"
	CtKey2                  = "key2"
	CtName                  = "name"
	CtOverlapKey1           = "overlapkey1"
	CtOverlapKey2           = "overlapkey2"
	CtPrivateKeyPEM         = "private_key_pem"
	CtProviderKey1          = "providerkey1"
	CtProviderTags          = "provider_tags"
	CtProviderValue1        = "providervalue1"
	CtRName                 = "rName"
	CtResourceKey1          = "resourcekey1"
	CtResourceKey2          = "resourcekey2"
	CtResourceOwner         = "resource_owner"
	CtResourceTags          = "resource_tags"
	CtResourceValue1        = "resourcevalue1"
	CtResourceValue1Updated = "resourcevalue1updated"
	CtResourceValue2        = "resourcevalue2"
	CtRulePound             = "rule.#"
	CtTagsAllPercent        = "tags_all.%"
	CtTagsKey1              = "tags.key1"
	CtTagsKey2              = "tags.key2"
	CtTagsPercent           = "tags.%"
	CtTrue                  = "true"
	CtTrueCaps              = "TRUE"
	CtValue1                = "value1"
	CtValue1Updated         = "value1updated"
	CtValue2                = "value2"
)

// ConstOrQuote returns the constant name for the given attribute if it exists.
// Otherwise, it returns the attribute quoted. This is intended for use in
// generated code and templates.
func ConstOrQuote(constant string) string {
	allConstants := map[string]string{
		"0":                     "Ct0",
		"1":                     "Ct1",
		"10":                    "Ct10",
		"2":                     "Ct2",
		"3":                     "Ct3",
		"4":                     "Ct4",
		"basic":                 "CtBasic",
		"certificate_pem":       "CtCertificatePEM",
		"disappears":            "CtDisappears",
		"false":                 "CtFalse",
		"FALSE":                 "CtFalseCaps",
		"key1":                  "CtKey1",
		"key2":                  "CtKey2",
		"name":                  "CtName",
		"overlapkey1":           "CtOverlapKey1",
		"overlapkey2":           "CtOverlapKey2",
		"private_key_pem":       "CtPrivateKeyPEM",
		"providerkey1":          "CtProviderKey1",
		"provider_tags":         "CtProviderTags",
		"providervalue1":        "CtProviderValue1",
		"rName":                 "CtRName",
		"resourcekey1":          "CtResourceKey1",
		"resourcekey2":          "CtResourceKey2",
		"resource_owner":        "CtResourceOwner",
		"resource_tags":         "CtResourceTags",
		"resourcevalue1":        "CtResourceValue1",
		"resourcevalue1updated": "CtResourceValue1Updated",
		"resourcevalue2":        "CtResourceValue2",
		"rule.#":                "CtRulePound",
		"tags_all.%":            "CtTagsAllPercent",
		"tags.key1":             "CtTagsKey1",
		"tags.key2":             "CtTagsKey2",
		"tags.%":                "CtTagsPercent",
		"true":                  "CtTrue",
		"TRUE":                  "CtTrueCaps",
		"value1":                "CtValue1",
		"value1updated":         "CtValue1Updated",
		"value2":                "CtValue2",
	}

	if v, ok := allConstants[constant]; ok {
		return fmt.Sprintf("acctest.%s", v)
	}
	return fmt.Sprintf("%q", constant)
}
