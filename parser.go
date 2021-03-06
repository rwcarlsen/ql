// Copyright (c) 2013 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Inital yacc source generated by ebnf2y[1]
// at 2013-10-04 23:10:47.861401015 +0200 CEST
//
//  $ ebnf2y -o ql.y -oe ql.ebnf -start StatementList -pkg ql -p _
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
//
//   [1]: http://github.com/cznic/ebnf2y

package ql

import __yyfmt__ "fmt"

import (
	"fmt"
)

type yySymType struct {
	yys  int
	line int
	col  int
	item interface{}
	list []interface{}
}

const add = 57346
const alter = 57347
const and = 57348
const andand = 57349
const andnot = 57350
const as = 57351
const asc = 57352
const begin = 57353
const between = 57354
const boolType = 57355
const by = 57356
const byteType = 57357
const column = 57358
const commit = 57359
const complex128Type = 57360
const complex64Type = 57361
const create = 57362
const deleteKwd = 57363
const desc = 57364
const distinct = 57365
const drop = 57366
const eq = 57367
const falseKwd = 57368
const float = 57369
const float32Type = 57370
const float64Type = 57371
const floatLit = 57372
const from = 57373
const ge = 57374
const group = 57375
const identifier = 57376
const imaginaryLit = 57377
const in = 57378
const insert = 57379
const intType = 57380
const int16Type = 57381
const int32Type = 57382
const int64Type = 57383
const int8Type = 57384
const is = 57385
const into = 57386
const intLit = 57387
const le = 57388
const lsh = 57389
const neq = 57390
const not = 57391
const null = 57392
const order = 57393
const oror = 57394
const qlParam = 57395
const rollback = 57396
const rsh = 57397
const runeType = 57398
const selectKwd = 57399
const stringType = 57400
const stringLit = 57401
const tableKwd = 57402
const transaction = 57403
const trueKwd = 57404
const truncate = 57405
const uintType = 57406
const uint16Type = 57407
const uint32Type = 57408
const uint64Type = 57409
const uint8Type = 57410
const update = 57411
const values = 57412
const where = 57413

var yyToknames = []string{
	"add",
	"alter",
	"and",
	"andand",
	"andnot",
	"as",
	"asc",
	"begin",
	"between",
	"boolType",
	"by",
	"byteType",
	"column",
	"commit",
	"complex128Type",
	"complex64Type",
	"create",
	"deleteKwd",
	"desc",
	"distinct",
	"drop",
	"eq",
	"falseKwd",
	"float",
	"float32Type",
	"float64Type",
	"floatLit",
	"from",
	"ge",
	"group",
	"identifier",
	"imaginaryLit",
	"in",
	"insert",
	"intType",
	"int16Type",
	"int32Type",
	"int64Type",
	"int8Type",
	"is",
	"into",
	"intLit",
	"le",
	"lsh",
	"neq",
	"not",
	"null",
	"order",
	"oror",
	"qlParam",
	"rollback",
	"rsh",
	"runeType",
	"selectKwd",
	"stringType",
	"stringLit",
	"tableKwd",
	"transaction",
	"trueKwd",
	"truncate",
	"uintType",
	"uint16Type",
	"uint32Type",
	"uint64Type",
	"uint8Type",
	"update",
	"values",
	"where",
}
var yyStatenames = []string{}

const yyEOFCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 175
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 626

var yyAct = []int{

	193, 150, 194, 156, 253, 233, 52, 149, 208, 159,
	96, 12, 66, 49, 108, 209, 238, 53, 75, 50,
	76, 47, 26, 77, 78, 144, 108, 54, 274, 258,
	257, 67, 79, 80, 81, 70, 48, 108, 97, 74,
	71, 218, 108, 82, 83, 84, 85, 86, 250, 161,
	72, 108, 219, 260, 248, 68, 126, 102, 63, 239,
	224, 87, 137, 88, 73, 195, 136, 69, 142, 89,
	90, 91, 92, 93, 122, 123, 124, 125, 245, 65,
	161, 222, 205, 269, 243, 56, 103, 58, 59, 162,
	241, 214, 103, 204, 207, 128, 57, 167, 143, 105,
	101, 267, 247, 129, 216, 230, 153, 157, 227, 152,
	48, 146, 199, 107, 148, 103, 155, 23, 28, 35,
	162, 170, 164, 173, 174, 175, 176, 177, 178, 163,
	154, 166, 127, 130, 131, 132, 31, 29, 27, 190,
	179, 180, 181, 182, 23, 110, 196, 108, 255, 55,
	217, 200, 32, 202, 183, 184, 185, 186, 187, 188,
	189, 169, 114, 201, 122, 123, 124, 125, 213, 172,
	171, 98, 236, 197, 212, 121, 215, 165, 122, 123,
	124, 125, 116, 106, 235, 168, 112, 37, 108, 30,
	271, 2, 99, 115, 221, 34, 118, 151, 120, 113,
	264, 97, 272, 256, 211, 228, 138, 139, 140, 141,
	225, 36, 100, 111, 237, 240, 232, 231, 38, 266,
	246, 226, 244, 242, 229, 198, 117, 119, 158, 145,
	14, 13, 251, 1, 134, 249, 44, 33, 252, 39,
	11, 40, 41, 42, 43, 210, 160, 94, 259, 64,
	261, 270, 254, 60, 62, 262, 104, 263, 157, 10,
	265, 75, 133, 76, 234, 268, 77, 78, 46, 109,
	273, 51, 223, 3, 67, 79, 80, 81, 70, 9,
	8, 203, 74, 71, 7, 61, 82, 83, 84, 85,
	86, 6, 206, 72, 192, 135, 5, 147, 68, 95,
	4, 63, 0, 0, 87, 0, 88, 73, 0, 0,
	69, 0, 89, 90, 91, 92, 93, 0, 0, 75,
	0, 76, 65, 0, 77, 78, 0, 220, 56, 0,
	58, 59, 67, 79, 80, 81, 70, 0, 0, 57,
	74, 71, 0, 0, 82, 83, 84, 85, 86, 0,
	0, 72, 0, 0, 0, 0, 68, 0, 0, 63,
	0, 0, 87, 0, 88, 73, 0, 0, 69, 0,
	89, 90, 91, 92, 93, 0, 0, 75, 0, 76,
	65, 0, 77, 78, 0, 0, 56, 0, 58, 59,
	67, 79, 80, 81, 70, 0, 191, 57, 74, 71,
	0, 0, 82, 83, 84, 85, 86, 0, 0, 72,
	0, 0, 0, 0, 68, 0, 0, 63, 0, 0,
	87, 0, 88, 73, 0, 0, 69, 0, 89, 90,
	91, 92, 93, 0, 0, 0, 0, 0, 65, 0,
	0, 0, 0, 75, 56, 76, 58, 59, 77, 78,
	0, 45, 0, 0, 0, 57, 67, 79, 80, 81,
	70, 0, 0, 0, 74, 71, 0, 0, 82, 83,
	84, 85, 86, 0, 0, 72, 0, 0, 0, 0,
	68, 0, 0, 63, 0, 0, 87, 0, 88, 73,
	0, 0, 69, 0, 89, 90, 91, 92, 93, 0,
	0, 75, 0, 76, 65, 0, 77, 78, 0, 0,
	56, 0, 58, 59, 67, 79, 80, 81, 70, 0,
	0, 57, 74, 71, 0, 0, 82, 83, 84, 85,
	86, 0, 0, 72, 0, 0, 0, 0, 68, 0,
	0, 63, 0, 75, 87, 76, 88, 73, 77, 78,
	69, 0, 89, 90, 91, 92, 93, 79, 80, 81,
	0, 15, 65, 0, 0, 0, 0, 16, 82, 83,
	84, 85, 86, 17, 0, 0, 18, 19, 0, 0,
	20, 0, 0, 0, 0, 0, 87, 0, 88, 0,
	0, 0, 0, 21, 89, 90, 91, 92, 93, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	22, 0, 0, 23, 0, 0, 0, 0, 0, 24,
	0, 0, 0, 0, 0, 25,
}
var yyPact = []int{

	556, -67, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 78, 57, -1000, 77, 158,
	76, 108, -1000, 172, 59, 153, 556, 153, -1000, 153,
	153, 153, 153, 364, -1000, 153, 137, -1000, -1000, 188,
	26, 44, -1000, 25, 152, -1000, 40, -1000, 136, 206,
	-1000, 150, -6, 48, -1000, -12, 488, 488, 488, 488,
	-1000, -1000, -1000, -1000, -1000, 430, 24, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -63, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 44, -1000, 42, -1000, 137,
	181, 137, -1000, 430, 60, 137, 46, 430, 430, -1000,
	143, 430, 23, 149, 430, 120, 430, 430, 430, 430,
	430, 430, 430, 430, 430, 430, 430, 430, 430, 430,
	430, 430, 430, -1000, -1000, -1000, 306, 430, -12, -12,
	-12, -12, -10, 430, 139, -1000, -1000, 39, 430, -1000,
	530, 137, -1000, 95, 19, -1000, 7, -1000, 21, -1000,
	195, -1000, 87, -1000, 206, -1000, -1000, 430, 17, 430,
	98, -1000, 100, -6, -6, -6, -6, -6, -6, 48,
	48, 48, 48, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-38, 248, 6, -1000, 95, -1000, -15, -1000, -1000, 137,
	95, -1000, -1000, 35, 430, -1000, 32, 15, 151, -1000,
	-1000, 138, -73, -16, 430, 84, 430, -1000, -1000, 5,
	-1000, -1, -1000, 29, -1000, -1000, -21, 137, -27, -1000,
	137, -1000, 151, 97, -1000, 189, -1000, -45, -1000, -1000,
	-46, 430, -6, -1000, -26, -1000, -1000, 430, -1000, -1000,
	-1000, -1000, 97, -1000, -1000, 186, 137, -1000, -1000, -6,
	-1000, 95, 28, -1000, 430, -1000, -1000, 9, 180, 430,
	-1000, -1000, -1000, -47, -1000,
}
var yyPgo = []int{

	0, 300, 10, 299, 297, 296, 295, 294, 7, 1,
	3, 292, 291, 285, 284, 281, 280, 279, 273, 2,
	0, 272, 19, 271, 21, 269, 268, 264, 262, 259,
	256, 255, 254, 253, 252, 251, 249, 149, 6, 17,
	9, 246, 245, 240, 11, 237, 236, 8, 5, 4,
	234, 191, 233, 211, 13, 231, 12, 27, 230, 229,
	15, 228, 225, 224, 221, 220, 219, 214,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 4, 4, 62, 62, 5,
	6, 7, 7, 8, 9, 10, 11, 11, 63, 63,
	12, 13, 14, 15, 15, 64, 64, 16, 16, 17,
	18, 19, 19, 20, 21, 21, 65, 65, 22, 22,
	22, 22, 22, 22, 22, 23, 23, 23, 23, 23,
	23, 23, 24, 25, 25, 26, 26, 27, 28, 29,
	29, 30, 30, 31, 31, 66, 66, 32, 32, 32,
	32, 32, 32, 32, 33, 33, 33, 33, 34, 35,
	35, 35, 37, 37, 37, 37, 37, 38, 38, 38,
	38, 38, 39, 39, 39, 39, 39, 39, 39, 39,
	36, 36, 40, 41, 41, 67, 67, 42, 42, 61,
	61, 43, 44, 44, 45, 45, 46, 46, 46, 47,
	47, 48, 48, 49, 49, 50, 50, 50, 50, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 52, 52, 53, 54, 54, 55, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 58, 59, 59, 57,
	57, 57, 57, 57, 60,
}
var yyR2 = []int{

	0, 5, 6, 3, 3, 0, 3, 0, 1, 2,
	3, 0, 1, 2, 1, 3, 0, 3, 0, 1,
	1, 4, 8, 0, 3, 0, 1, 3, 4, 3,
	0, 1, 3, 3, 0, 3, 0, 1, 1, 5,
	6, 5, 6, 3, 4, 1, 3, 3, 3, 3,
	3, 3, 2, 0, 2, 1, 3, 3, 3, 10,
	5, 0, 3, 0, 5, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 0,
	1, 1, 1, 1, 2, 2, 2, 1, 3, 3,
	3, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	1, 3, 2, 1, 4, 0, 1, 0, 2, 1,
	3, 1, 8, 9, 0, 1, 1, 1, 2, 0,
	1, 0, 1, 0, 1, 3, 4, 4, 5, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 0, 1, 1,
	2, 2, 2, 2, 2,
}
var yyChk = []int{

	-1000, -52, -51, -18, -1, -5, -12, -14, -16, -17,
	-29, -43, -44, -55, -58, 5, 11, 17, 20, 21,
	24, 37, 54, 57, 63, 69, 89, 60, 61, 60,
	31, 60, 44, -45, 23, 60, -53, 34, -51, -53,
	-53, -53, -53, -53, -46, 87, -26, -24, -19, -54,
	-22, -23, -38, -39, -57, -37, 80, 91, 82, 83,
	-33, -13, -32, 53, -36, 74, -56, 26, 50, 62,
	30, 35, 45, 59, 34, 13, 15, 18, 19, 27,
	28, 29, 38, 39, 40, 41, 42, 56, 58, 64,
	65, 66, 67, 68, -53, -3, -2, -9, 34, 4,
	24, 74, -60, 71, -30, 74, 31, 73, 52, -25,
	9, 7, 36, 49, 12, 43, 32, 76, 46, 77,
	48, 25, 80, 81, 82, 83, 8, 84, 47, 55,
	85, 86, 87, -28, -50, -6, 78, 74, -37, -37,
	-37, -37, -19, 74, 88, -59, -60, -4, 72, -8,
	-9, 16, -8, -19, 70, -44, -10, -9, -61, -40,
	-41, 34, 74, -24, -54, 34, -22, 74, 36, 12,
	-38, 50, 49, -38, -38, -38, -38, -38, -38, -39,
	-39, -39, -39, -57, -57, -57, -57, -57, -57, -57,
	-19, 90, -7, -20, -19, 75, -19, 34, -62, 73,
	-19, -56, -9, -15, 74, 75, -11, 73, -47, -60,
	-42, 9, -44, -20, 74, -38, 6, 50, 79, 90,
	79, -19, 75, -21, 75, -2, -64, 73, -20, -63,
	73, -40, -47, -48, -27, 33, 34, -67, 89, 75,
	-20, 6, -38, 79, -19, 79, -65, 73, 75, -8,
	75, -9, -48, -49, -34, 51, 14, 75, 75, -38,
	79, -19, -31, -49, 14, -10, -66, 73, -20, 74,
	-35, 10, 22, -20, 75,
}
var yyDef = []int{

	30, -2, 141, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 0, 0, 20, 0, 0,
	0, 0, 111, 114, 0, 0, 30, 0, 9, 0,
	0, 0, 0, 0, 115, 0, 0, 143, 142, 0,
	0, 27, 29, 61, 0, 116, 117, 55, 53, 31,
	144, 38, 45, 87, 92, 169, 0, 0, 0, 0,
	82, 83, 74, 75, 76, 0, 0, 67, 68, 69,
	70, 71, 72, 73, 100, 147, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 165, 146, 167, 5, 0, 14, 0,
	0, 0, 28, 0, 0, 0, 0, 118, 0, 52,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 84, 85, 86, 0, 11, 170, 171,
	172, 173, 0, 0, 0, 166, 168, 7, 0, 1,
	0, 0, 23, 174, 0, 60, 0, 16, 119, 109,
	107, 103, 0, 56, 32, 54, 145, 0, 0, 0,
	0, 43, 0, 46, 47, 48, 49, 50, 51, 88,
	89, 90, 91, 93, 94, 95, 96, 97, 98, 99,
	0, 0, 0, 12, 34, 77, 0, 101, 4, 8,
	3, 13, 2, 25, 0, 62, 18, 119, 121, 120,
	102, 0, 105, 0, 0, 0, 0, 44, 58, 0,
	125, 0, 10, 36, 21, 6, 0, 26, 0, 15,
	19, 110, 121, 123, 122, 0, 108, 0, 106, 39,
	0, 0, 41, 127, 0, 126, 33, 37, 22, 24,
	63, 17, 123, 112, 124, 0, 0, 104, 40, 42,
	128, 35, 65, 113, 0, 57, 59, 66, 79, 0,
	78, 80, 81, 0, 64,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 91, 3, 3, 3, 85, 84, 3,
	74, 75, 87, 83, 73, 82, 88, 86, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 90, 89,
	77, 72, 76, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 78, 3, 79, 80, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 81,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
}
var yyTok3 = []int{
	0,
}

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEOFCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:

		{
			yyVAL.item = &alterTableAddStmt{tableName: yyS[yypt-2].item.(string), c: yyS[yypt-0].item.(*col)}
		}
	case 2:

		{
			yyVAL.item = &alterTableDropColumnStmt{tableName: yyS[yypt-3].item.(string), colName: yyS[yypt-0].item.(string)}
		}
	case 3:

		{
			yyVAL.item = assignment{colName: yyS[yypt-2].item.(string), expr: yyS[yypt-0].item.(expression)}
		}
	case 4:

		{
			yyVAL.item = append([]assignment{yyS[yypt-2].item.(assignment)}, yyS[yypt-1].item.([]assignment)...)
		}
	case 5:

		{
			yyVAL.item = []assignment{}
		}
	case 6:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]assignment), yyS[yypt-0].item.(assignment))
		}
	case 9:

		{
			yyVAL.item = beginTransactionStmt{}
		}
	case 10:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 11:

		{
			yyVAL.item = []expression{}
		}
	case 12:
		yyVAL.item = yyS[yypt-0].item
	case 13:

		{
			yyVAL.item = &col{name: yyS[yypt-1].item.(string), typ: yyS[yypt-0].item.(int)}
		}
	case 14:
		yyVAL.item = yyS[yypt-0].item
	case 15:

		{
			yyVAL.item = append([]string{yyS[yypt-2].item.(string)}, yyS[yypt-1].item.([]string)...)
		}
	case 16:

		{
			yyVAL.item = []string{}
		}
	case 17:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]string), yyS[yypt-0].item.(string))
		}
	case 20:

		{
			yyVAL.item = commitStmt{}
		}
	case 21:

		{
			yyVAL.item = &conversion{typ: yyS[yypt-3].item.(int), val: yyS[yypt-1].item.(expression)}
		}
	case 22:

		{
			yyVAL.item = &createTableStmt{tableName: yyS[yypt-5].item.(string), cols: append([]*col{yyS[yypt-3].item.(*col)}, yyS[yypt-2].item.([]*col)...)}
		}
	case 23:

		{
			yyVAL.item = []*col{}
		}
	case 24:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]*col), yyS[yypt-0].item.(*col))
		}
	case 27:

		{
			yyVAL.item = &truncateTableStmt{yyS[yypt-0].item.(string)}
		}
	case 28:

		{
			yyVAL.item = &deleteStmt{tableName: yyS[yypt-1].item.(string), where: yyS[yypt-0].item.(*whereRset).expr}
		}
	case 29:

		{
			yyVAL.item = &dropTableStmt{tableName: yyS[yypt-0].item.(string)}
		}
	case 30:

		{
			yyVAL.item = nil
		}
	case 31:
		yyVAL.item = yyS[yypt-0].item
	case 32:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(oror, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 33:

		{
			yyVAL.item = append([]expression{yyS[yypt-2].item.(expression)}, yyS[yypt-1].item.([]expression)...)
		}
	case 34:

		{
			yyVAL.item = []expression(nil)
		}
	case 35:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]expression), yyS[yypt-0].item.(expression))
		}
	case 38:
		yyVAL.item = yyS[yypt-0].item
	case 39:

		{
			yyVAL.item = &pIn{expr: yyS[yypt-4].item.(expression), list: yyS[yypt-1].item.([]expression)}
		}
	case 40:

		{
			yyVAL.item = &pIn{expr: yyS[yypt-5].item.(expression), not: true, list: yyS[yypt-1].item.([]expression)}
		}
	case 41:

		{
			yyVAL.item = &pBetween{expr: yyS[yypt-4].item.(expression), l: yyS[yypt-2].item.(expression), h: yyS[yypt-0].item.(expression)}
		}
	case 42:

		{
			yyVAL.item = &pBetween{expr: yyS[yypt-5].item.(expression), not: true, l: yyS[yypt-2].item.(expression), h: yyS[yypt-0].item.(expression)}
		}
	case 43:

		{
			yyVAL.item = &isNull{expr: yyS[yypt-2].item.(expression)}
		}
	case 44:

		{
			yyVAL.item = &isNull{expr: yyS[yypt-3].item.(expression), not: true}
		}
	case 45:
		yyVAL.item = yyS[yypt-0].item
	case 46:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(ge, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 47:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('>', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 48:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(le, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 49:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('<', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 50:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(neq, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 51:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(eq, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 52:

		{
			expr, name := yyS[yypt-1].item.(expression), yyS[yypt-0].item.(string)
			if name == "" {
				s, ok := expr.(*ident)
				if ok {
					name = s.s
				}
			}
			yyVAL.item = &fld{expr: expr, name: name}
		}
	case 53:

		{
			yyVAL.item = ""
		}
	case 54:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 55:

		{
			yyVAL.item = []*fld{yyS[yypt-0].item.(*fld)}
		}
	case 56:

		{
			l, f := yyS[yypt-2].item.([]*fld), yyS[yypt-0].item.(*fld)
			if f.name != "" {
				if f := findFld(l, f.name); f != nil {
					yylex.(*lexer).err("duplicate field name %q", f.name)
					goto ret1
				}
			}

			yyVAL.item = append(yyS[yypt-2].item.([]*fld), yyS[yypt-0].item.(*fld))
		}
	case 57:

		{
			yyVAL.item = &groupByRset{colNames: yyS[yypt-0].item.([]string)}
		}
	case 58:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 59:

		{
			yyVAL.item = &insertIntoStmt{tableName: yyS[yypt-7].item.(string), colNames: yyS[yypt-6].item.([]string), lists: append([][]expression{yyS[yypt-3].item.([]expression)}, yyS[yypt-1].item.([][]expression)...)}
		}
	case 60:

		{
			yyVAL.item = &insertIntoStmt{tableName: yyS[yypt-2].item.(string), colNames: yyS[yypt-1].item.([]string), sel: yyS[yypt-0].item.(*selectStmt)}
		}
	case 61:

		{
			yyVAL.item = []string{}
		}
	case 62:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 63:

		{
			yyVAL.item = [][]expression{}
		}
	case 64:

		{
			yyVAL.item = append(yyS[yypt-4].item.([][]expression), yyS[yypt-1].item.([]expression))
		}
	case 67:
		yyVAL.item = yyS[yypt-0].item
	case 68:
		yyVAL.item = yyS[yypt-0].item
	case 69:
		yyVAL.item = yyS[yypt-0].item
	case 70:
		yyVAL.item = yyS[yypt-0].item
	case 71:
		yyVAL.item = yyS[yypt-0].item
	case 72:
		yyVAL.item = yyS[yypt-0].item
	case 73:
		yyVAL.item = yyS[yypt-0].item
	case 74:

		{
			yyVAL.item = value{yyS[yypt-0].item}
		}
	case 75:

		{
			yyVAL.item = parameter{yyS[yypt-0].item.(int)}
		}
	case 76:

		{
			yyVAL.item = &ident{yyS[yypt-0].item.(string)}
		}
	case 77:

		{
			yyVAL.item = &pexpr{expr: yyS[yypt-1].item.(expression)}
		}
	case 78:

		{
			yyVAL.item = &orderByRset{by: yyS[yypt-1].item.([]expression), asc: yyS[yypt-0].item.(bool)}
		}
	case 79:

		{
			yyVAL.item = true // ASC by default
		}
	case 80:

		{
			yyVAL.item = true
		}
	case 81:

		{
			yyVAL.item = false
		}
	case 82:
		yyVAL.item = yyS[yypt-0].item
	case 83:
		yyVAL.item = yyS[yypt-0].item
	case 84:

		{
			var err error
			if yyVAL.item, err = newIndex(yyS[yypt-1].item.(expression), yyS[yypt-0].item.(expression)); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 85:

		{
			var err error
			s := yyS[yypt-0].item.([2]*expression)
			if yyVAL.item, err = newSlice(yyS[yypt-1].item.(expression), s[0], s[1]); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 86:

		{
			x := yylex.(*lexer)
			f, ok := yyS[yypt-1].item.(*ident)
			if !ok {
				x.err("expected identifier or qualified identifier")
				goto ret1
			}

			var err error
			var agg bool
			if yyVAL.item, agg, err = newCall(f.s, yyS[yypt-0].item.([]expression)); err != nil {
				x.err("%v", err)
				goto ret1
			}
			if n := len(x.agg); n > 0 {
				x.agg[n-1] = x.agg[n-1] || agg
			}
		}
	case 87:
		yyVAL.item = yyS[yypt-0].item
	case 88:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('^', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 89:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('|', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 90:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('-', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 91:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('+', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 92:
		yyVAL.item = yyS[yypt-0].item
	case 93:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(andnot, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 94:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('&', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 95:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(lsh, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 96:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(rsh, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 97:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('%', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 98:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('/', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 99:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('*', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 100:
		yyVAL.item = yyS[yypt-0].item
	case 101:

		{
			yyVAL.item = fmt.Sprintf("%s.%s", yyS[yypt-2].item.(string), yyS[yypt-0].item.(string))
		}
	case 102:

		{
			yyVAL.item = []interface{}{yyS[yypt-1].item, yyS[yypt-0].item}
		}
	case 103:
		yyVAL.item = yyS[yypt-0].item
	case 104:

		{
			yyVAL.item = yyS[yypt-2].item
		}
	case 107:

		{
			yyVAL.item = ""
		}
	case 108:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 109:

		{
			yyVAL.list = []interface{}{yyS[yypt-0].item}
		}
	case 110:

		{
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].item)
		}
	case 111:

		{
			yyVAL.item = rollbackStmt{}
		}
	case 112:

		{
			x := yylex.(*lexer)
			n := len(x.agg)
			yyVAL.item = &selectStmt{
				distinct:      yyS[yypt-6].item.(bool),
				flds:          yyS[yypt-5].item.([]*fld),
				from:          &crossJoinRset{sources: yyS[yypt-3].list},
				hasAggregates: x.agg[n-1],
				where:         yyS[yypt-2].item.(*whereRset),
				group:         yyS[yypt-1].item.(*groupByRset),
				order:         yyS[yypt-0].item.(*orderByRset),
			}
			x.agg = x.agg[:n-1]
		}
	case 113:

		{
			x := yylex.(*lexer)
			n := len(x.agg)
			yyVAL.item = &selectStmt{
				distinct:      yyS[yypt-7].item.(bool),
				flds:          yyS[yypt-6].item.([]*fld),
				from:          &crossJoinRset{sources: yyS[yypt-4].list},
				hasAggregates: x.agg[n-1],
				where:         yyS[yypt-2].item.(*whereRset),
				group:         yyS[yypt-1].item.(*groupByRset),
				order:         yyS[yypt-0].item.(*orderByRset),
			}
			x.agg = x.agg[:n-1]
		}
	case 114:

		{
			yyVAL.item = false
		}
	case 115:

		{
			yyVAL.item = true
		}
	case 116:

		{
			yyVAL.item = []*fld{}
		}
	case 117:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 118:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 119:

		{
			yyVAL.item = (*whereRset)(nil)
		}
	case 120:
		yyVAL.item = yyS[yypt-0].item
	case 121:

		{
			yyVAL.item = (*groupByRset)(nil)
		}
	case 122:
		yyVAL.item = yyS[yypt-0].item
	case 123:

		{
			yyVAL.item = (*orderByRset)(nil)
		}
	case 124:
		yyVAL.item = yyS[yypt-0].item
	case 125:

		{
			yyVAL.item = [2]*expression{nil, nil}
		}
	case 126:

		{
			hi := yyS[yypt-1].item.(expression)
			yyVAL.item = [2]*expression{nil, &hi}
		}
	case 127:

		{
			lo := yyS[yypt-2].item.(expression)
			yyVAL.item = [2]*expression{&lo, nil}
		}
	case 128:

		{
			lo := yyS[yypt-3].item.(expression)
			hi := yyS[yypt-1].item.(expression)
			yyVAL.item = [2]*expression{&lo, &hi}
		}
	case 129:
		yyVAL.item = yyS[yypt-0].item
	case 130:
		yyVAL.item = yyS[yypt-0].item
	case 131:
		yyVAL.item = yyS[yypt-0].item
	case 132:
		yyVAL.item = yyS[yypt-0].item
	case 133:
		yyVAL.item = yyS[yypt-0].item
	case 134:
		yyVAL.item = yyS[yypt-0].item
	case 135:
		yyVAL.item = yyS[yypt-0].item
	case 136:
		yyVAL.item = yyS[yypt-0].item
	case 137:
		yyVAL.item = yyS[yypt-0].item
	case 138:
		yyVAL.item = yyS[yypt-0].item
	case 139:
		yyVAL.item = yyS[yypt-0].item
	case 140:
		yyVAL.item = yyS[yypt-0].item
	case 141:

		{
			if yyS[yypt-0].item != nil {
				yylex.(*lexer).list = []stmt{yyS[yypt-0].item.(stmt)}
			}
		}
	case 142:

		{
			if yyS[yypt-0].item != nil {
				yylex.(*lexer).list = append(yylex.(*lexer).list, yyS[yypt-0].item.(stmt))
			}
		}
	case 143:
		yyVAL.item = yyS[yypt-0].item
	case 144:
		yyVAL.item = yyS[yypt-0].item
	case 145:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(andand, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 146:

		{
			yyVAL.item = &truncateTableStmt{tableName: yyS[yypt-0].item.(string)}
		}
	case 147:
		yyVAL.item = yyS[yypt-0].item
	case 148:
		yyVAL.item = yyS[yypt-0].item
	case 149:
		yyVAL.item = yyS[yypt-0].item
	case 150:
		yyVAL.item = yyS[yypt-0].item
	case 151:
		yyVAL.item = yyS[yypt-0].item
	case 152:
		yyVAL.item = yyS[yypt-0].item
	case 153:
		yyVAL.item = yyS[yypt-0].item
	case 154:
		yyVAL.item = yyS[yypt-0].item
	case 155:
		yyVAL.item = yyS[yypt-0].item
	case 156:
		yyVAL.item = yyS[yypt-0].item
	case 157:
		yyVAL.item = yyS[yypt-0].item
	case 158:
		yyVAL.item = yyS[yypt-0].item
	case 159:
		yyVAL.item = yyS[yypt-0].item
	case 160:
		yyVAL.item = yyS[yypt-0].item
	case 161:
		yyVAL.item = yyS[yypt-0].item
	case 162:
		yyVAL.item = yyS[yypt-0].item
	case 163:
		yyVAL.item = yyS[yypt-0].item
	case 164:
		yyVAL.item = yyS[yypt-0].item
	case 165:
		yyVAL.item = yyS[yypt-0].item
	case 166:

		{
			yyVAL.item = &updateStmt{tableName: yyS[yypt-2].item.(string), list: yyS[yypt-1].item.([]assignment), where: yyS[yypt-0].item.(*whereRset).expr}
		}
	case 167:

		{
			yyVAL.item = nowhere
		}
	case 168:
		yyVAL.item = yyS[yypt-0].item
	case 169:
		yyVAL.item = yyS[yypt-0].item
	case 170:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('^', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 171:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('!', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 172:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('-', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 173:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('+', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 174:

		{
			yyVAL.item = &whereRset{expr: yyS[yypt-0].item.(expression)}
		}
	}
	goto yystack /* stack new state and value */
}
