// Code generated by goyacc -o wkt_generated.go -p wkt wkt.y. DO NOT EDIT.

//line wkt.y:12

package wkt

import __yyfmt__ "fmt"

//line wkt.y:13

import "github.com/twpayne/go-geom"

func isValidLineString(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 2*stride {
		wktlex.(*wktLex).Error("syntax error: non-empty linestring with only one point")
		return false
	}
	return true
}

func isValidPolygonRing(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 4*stride {
		wktlex.(*wktLex).Error("syntax error: polygon ring doesn't have enough points")
		return false
	}
	for i := 0; i < stride; i++ {
		if flatCoords[i] != flatCoords[len(flatCoords)-stride+i] {
			wktlex.(*wktLex).Error("syntax error: polygon ring not closed")
			return false
		}
	}
	return true
}

type geomFlatCoordsRepr struct {
	flatCoords []float64
	ends       []int
}

func makeGeomFlatCoordsRepr(flatCoords []float64) geomFlatCoordsRepr {
	return geomFlatCoordsRepr{flatCoords: flatCoords, ends: []int{len(flatCoords)}}
}

func appendGeomFlatCoordsReprs(p1 geomFlatCoordsRepr, p2 geomFlatCoordsRepr) geomFlatCoordsRepr {
	if len(p1.ends) > 0 {
		p1LastEnd := p1.ends[len(p1.ends)-1]
		for i, _ := range p2.ends {
			p2.ends[i] += p1LastEnd
		}
	}
	return geomFlatCoordsRepr{flatCoords: append(p1.flatCoords, p2.flatCoords...), ends: append(p1.ends, p2.ends...)}
}

type multiPolygonFlatCoordsRepr struct {
	flatCoords []float64
	endss      [][]int
}

func makeMultiPolygonFlatCoordsRepr(p geomFlatCoordsRepr) multiPolygonFlatCoordsRepr {
	if p.flatCoords == nil {
		return multiPolygonFlatCoordsRepr{flatCoords: nil, endss: [][]int{nil}}
	}
	return multiPolygonFlatCoordsRepr{flatCoords: p.flatCoords, endss: [][]int{p.ends}}
}

func appendMultiPolygonFlatCoordsRepr(
	p1 multiPolygonFlatCoordsRepr, p2 multiPolygonFlatCoordsRepr,
) multiPolygonFlatCoordsRepr {
	p1LastEndsLastEnd := 0
	for i := len(p1.endss) - 1; i >= 0; i-- {
		if len(p1.endss[i]) > 0 {
			p1LastEndsLastEnd = p1.endss[i][len(p1.endss[i])-1]
			break
		}
	}
	if p1LastEndsLastEnd > 0 {
		for i, _ := range p2.endss {
			for j, _ := range p2.endss[i] {
				p2.endss[i][j] += p1LastEndsLastEnd
			}
		}
	}
	return multiPolygonFlatCoordsRepr{
		flatCoords: append(p1.flatCoords, p2.flatCoords...), endss: append(p1.endss, p2.endss...),
	}
}

//line wkt.y:94
type wktSymType struct {
	yys               int
	str               string
	geom              geom.T
	coord             float64
	coordList         []float64
	flatRepr          geomFlatCoordsRepr
	multiPolyFlatRepr multiPolygonFlatCoordsRepr
}

const POINT = 57346
const POINTM = 57347
const POINTZ = 57348
const POINTZM = 57349
const LINESTRING = 57350
const LINESTRINGM = 57351
const LINESTRINGZ = 57352
const LINESTRINGZM = 57353
const POLYGON = 57354
const POLYGONM = 57355
const POLYGONZ = 57356
const POLYGONZM = 57357
const MULTIPOINT = 57358
const MULTIPOINTM = 57359
const MULTIPOINTZ = 57360
const MULTIPOINTZM = 57361
const MULTILINESTRING = 57362
const MULTILINESTRINGM = 57363
const MULTILINESTRINGZ = 57364
const MULTILINESTRINGZM = 57365
const MULTIPOLYGON = 57366
const MULTIPOLYGONM = 57367
const MULTIPOLYGONZ = 57368
const MULTIPOLYGONZM = 57369
const EMPTY = 57370
const NUM = 57371

var wktToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"POINT",
	"POINTM",
	"POINTZ",
	"POINTZM",
	"LINESTRING",
	"LINESTRINGM",
	"LINESTRINGZ",
	"LINESTRINGZM",
	"POLYGON",
	"POLYGONM",
	"POLYGONZ",
	"POLYGONZM",
	"MULTIPOINT",
	"MULTIPOINTM",
	"MULTIPOINTZ",
	"MULTIPOINTZM",
	"MULTILINESTRING",
	"MULTILINESTRINGM",
	"MULTILINESTRINGZ",
	"MULTILINESTRINGZM",
	"MULTIPOLYGON",
	"MULTIPOLYGONM",
	"MULTIPOLYGONZ",
	"MULTIPOLYGONZM",
	"EMPTY",
	"NUM",
	"'('",
	"')'",
	"','",
}

var wktStatenames = [...]string{}

const wktEofCode = 1
const wktErrCode = 2
const wktInitialStackSize = 16

//line yacctab:1
var wktExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const wktPrivate = 57344

const wktLast = 323

var wktAct = [...]int{
	50, 99, 181, 184, 179, 185, 180, 162, 157, 159,
	138, 163, 147, 168, 142, 136, 123, 116, 51, 141,
	137, 115, 128, 158, 178, 156, 236, 237, 135, 235,
	234, 233, 234, 231, 232, 229, 230, 227, 228, 52,
	225, 226, 224, 223, 114, 222, 223, 220, 221, 218,
	219, 216, 217, 214, 215, 108, 213, 212, 211, 212,
	100, 209, 210, 100, 130, 207, 208, 117, 205, 206,
	203, 204, 201, 202, 199, 200, 109, 132, 189, 109,
	197, 198, 195, 196, 188, 118, 127, 177, 118, 74,
	125, 187, 143, 151, 172, 149, 124, 170, 169, 164,
	186, 246, 101, 193, 194, 74, 119, 148, 139, 177,
	101, 69, 69, 160, 119, 134, 104, 45, 110, 61,
	182, 134, 103, 40, 104, 45, 110, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 56, 134, 242, 250, 134, 102, 37, 177, 155,
	261, 61, 155, 155, 56, 246, 103, 40, 177, 155,
	66, 53, 98, 96, 97, 95, 94, 92, 93, 91,
	90, 88, 89, 87, 86, 84, 85, 83, 82, 80,
	81, 79, 78, 76, 77, 75, 241, 73, 71, 74,
	69, 117, 68, 65, 69, 66, 53, 60, 132, 61,
	58, 55, 56, 56, 49, 1, 53, 127, 44, 243,
	45, 118, 248, 249, 247, 252, 143, 151, 253, 251,
	255, 172, 258, 257, 259, 254, 164, 262, 263, 264,
	186, 265, 260, 256, 119, 245, 242, 42, 108, 40,
	39, 36, 40, 37, 267, 266, 104, 103, 240, 239,
	238, 192, 117, 244, 191, 190, 102, 174, 175, 154,
	153, 173, 133, 183, 165, 167, 166, 171, 152, 176,
	161, 144, 33, 46, 48, 64, 62, 59, 72, 63,
	67, 70, 47, 54, 57, 131, 146, 145, 150, 129,
	140, 120, 122, 121, 126, 35, 113, 112, 43, 34,
	38, 41, 111, 107, 106, 105, 8, 7, 6, 5,
	4, 3, 2,
}

var wktPact = [...]int{
	123, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 223,
	222, 219, 190, 186, 183, 182, 179, 175, 174, 170,
	169, 165, 164, 161, 160, 157, 156, 153, 152, 149,
	148, 145, 144, -1000, -1000, -1000, -1000, 237, -1000, -1000,
	228, -1000, -1000, -1000, -1000, 227, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 237, -1000, -1000, 228, -1000, -1000, -1000,
	-1000, 227, -1000, -1000, -1000, -1000, 176, -1000, -1000, 121,
	-1000, -1000, -1000, -1000, 89, 127, -1000, 93, -1000, 93,
	-1000, 87, -1000, 141, -1000, 134, -1000, 134, -1000, 131,
	-1000, 140, -1000, 81, -1000, 81, -1000, 59, -1000, 60,
	53, 47, 236, 235, 232, 72, 51, 49, -1000, -1000,
	-1000, 43, 41, 39, -1000, -1000, -1000, -1000, -1000, -1000,
	37, 34, 30, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 27, -1000, -1000, -1000, 25,
	22, -1000, -1000, -1000, 20, 18, 16, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 14, -1000, -1000, -1000,
	11, 9, -1000, -1000, -1000, 6, 4, 2, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 0, -1000,
	-1000, -1000, -2, -5, -1000, -1000, -1000, -1000, -1000, -1000,
	231, 230, 229, -1000, 217, -1000, 228, -1000, 227, -1000,
	71, -1000, 121, -1000, 89, -1000, 124, -1000, 137, -1000,
	95, -1000, 93, -1000, -1000, 87, -1000, 135, -1000, 121,
	-1000, 89, -1000, 134, -1000, -1000, 131, -1000, 130, -1000,
	82, -1000, 75, -1000, 81, -1000, -1000, 59, 226, -1000,
	226, -1000, 225, -1000, -1000, -1000, 217, -1000, -1000, -1000,
	217, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 71, -1000, -1000, -1000, -1000, -1000, -1000,
}

var wktPgo = [...]int{
	0, 322, 321, 320, 319, 318, 317, 316, 1, 22,
	64, 272, 299, 295, 315, 314, 313, 0, 18, 39,
	269, 278, 270, 44, 21, 17, 312, 307, 306, 304,
	20, 14, 303, 302, 10, 9, 16, 15, 19, 301,
	28, 300, 298, 23, 11, 12, 8, 7, 297, 296,
	281, 25, 280, 268, 271, 267, 279, 277, 6, 5,
	2, 276, 275, 13, 4, 3, 274, 24, 273, 215,
}

var wktR1 = [...]int{
	0, 69, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 61, 61,
	62, 62, 66, 66, 67, 67, 68, 68, 63, 63,
	64, 64, 65, 65, 57, 58, 59, 60, 53, 54,
	55, 56, 26, 26, 27, 27, 28, 28, 23, 24,
	25, 48, 48, 49, 49, 50, 50, 51, 51, 52,
	52, 45, 45, 46, 46, 47, 47, 42, 43, 44,
	20, 21, 22, 17, 18, 19, 35, 14, 14, 15,
	15, 16, 16, 32, 32, 33, 33, 39, 39, 40,
	40, 41, 41, 36, 36, 37, 37, 38, 38, 29,
	29, 30, 30, 31, 31, 34, 11, 12, 13, 8,
	9, 10,
}

var wktR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 4, 4,
	4, 4, 4, 4, 2, 2, 2, 2, 4, 4,
	4, 4, 4, 4, 2, 2, 2, 2, 4, 4,
	4, 4, 4, 4, 2, 2, 2, 2, 3, 1,
	3, 1, 3, 1, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 3,
	3, 1, 3, 1, 3, 1, 3, 1, 1, 1,
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 1, 3, 1, 3,
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 3, 3, 2,
	3, 4,
}

var wktChk = [...]int{
	-1000, -69, -1, -2, -3, -4, -5, -6, -7, 4,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, -11, -12, -13, 28, 30, -12, 28,
	30, -12, 28, -13, 28, 30, -20, -21, -22, 28,
	-17, -18, -19, 30, -21, 28, 30, -21, 28, -22,
	28, 30, -53, -54, -55, 28, 30, -54, 28, 30,
	-54, 28, -55, 28, 30, 30, 28, 30, 28, 30,
	28, 30, 28, 30, 28, 30, 28, 30, 28, 30,
	28, 30, 28, 30, 28, 30, 28, 30, 28, -8,
	-9, -10, 29, 29, 29, -14, -15, -16, -8, -9,
	-10, -26, -27, -28, -23, -24, -25, -17, -18, -19,
	-39, -32, -33, -36, -30, -31, -29, -34, -9, -12,
	-10, -13, -8, -11, 28, -40, -37, -30, -34, -40,
	-41, -38, -31, -34, -50, -48, -49, -45, -43, -44,
	-42, -35, -21, -22, -20, 28, -51, -46, -43, -35,
	-51, -52, -47, -44, -35, -66, -61, -62, -63, -58,
	-59, -57, -60, -54, -55, -53, -56, 28, -67, -64,
	-58, -60, -67, -68, -65, -59, -60, 31, 31, 31,
	29, 29, 29, 31, 32, 31, 32, 31, 32, 31,
	32, 31, 32, 31, 32, 31, 32, 31, 32, 31,
	32, 31, 32, 31, 31, 32, 31, 32, 31, 32,
	31, 32, 31, 32, 31, 31, 32, 31, 32, 31,
	32, 31, 32, 31, 32, 31, 31, 32, 29, 29,
	29, -8, 29, -9, -10, -23, 30, -24, -25, -36,
	30, -30, -31, -37, -38, -45, -43, -44, -46, -47,
	-63, 30, -58, -59, -64, -65, 29, 29,
}

var wktDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 9, 10, 14, 0, 11, 15,
	0, 12, 16, 13, 17, 0, 18, 19, 20, 24,
	120, 121, 122, 0, 21, 25, 0, 22, 26, 23,
	27, 0, 28, 29, 30, 34, 0, 31, 35, 0,
	32, 36, 33, 37, 0, 0, 44, 0, 45, 0,
	46, 0, 47, 0, 54, 0, 55, 0, 56, 0,
	57, 0, 64, 0, 65, 0, 66, 0, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 128, 130,
	132, 0, 0, 0, 93, 95, 97, 98, 99, 100,
	0, 0, 0, 138, 134, 136, 143, 144, 151, 152,
	153, 154, 149, 150, 155, 0, 140, 145, 146, 0,
	0, 142, 147, 148, 0, 0, 0, 106, 102, 104,
	111, 112, 118, 119, 117, 126, 0, 108, 113, 114,
	0, 0, 110, 115, 116, 0, 0, 0, 73, 69,
	71, 78, 79, 85, 86, 84, 87, 91, 0, 75,
	80, 81, 0, 0, 77, 82, 83, 156, 157, 158,
	159, 0, 0, 123, 0, 124, 0, 125, 0, 88,
	0, 89, 0, 90, 0, 38, 0, 39, 0, 40,
	0, 41, 0, 42, 43, 0, 48, 0, 49, 0,
	50, 0, 51, 0, 52, 53, 0, 58, 0, 59,
	0, 60, 0, 61, 0, 62, 63, 0, 160, 160,
	0, 127, 0, 129, 131, 92, 0, 94, 96, 137,
	0, 133, 135, 139, 141, 105, 101, 103, 107, 109,
	72, 0, 68, 70, 74, 76, 161, 159,
}

var wktTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 31, 3, 3, 32,
}

var wktTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
}

var wktTok3 = [...]int{
	0,
}

var wktErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	wktDebug        = 0
	wktErrorVerbose = true
)

type wktLexer interface {
	Lex(lval *wktSymType) int
	Error(s string)
}

type wktParser interface {
	Parse(wktLexer) int
	Lookahead() int
}

type wktParserImpl struct {
	lval  wktSymType
	stack [wktInitialStackSize]wktSymType
	char  int
}

func (p *wktParserImpl) Lookahead() int {
	return p.char
}

func wktNewParser() wktParser {
	return &wktParserImpl{}
}

const wktFlag = -1000

func wktTokname(c int) string {
	if c >= 1 && c-1 < len(wktToknames) {
		if wktToknames[c-1] != "" {
			return wktToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func wktStatname(s int) string {
	if s >= 0 && s < len(wktStatenames) {
		if wktStatenames[s] != "" {
			return wktStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func wktErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !wktErrorVerbose {
		return "syntax error"
	}

	for _, e := range wktErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + wktTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := wktPact[state]
	for tok := TOKSTART; tok-1 < len(wktToknames); tok++ {
		if n := base + tok; n >= 0 && n < wktLast && wktChk[wktAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if wktDef[state] == -2 {
		i := 0
		for wktExca[i] != -1 || wktExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; wktExca[i] >= 0; i += 2 {
			tok := wktExca[i]
			if tok < TOKSTART || wktExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if wktExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += wktTokname(tok)
	}
	return res
}

func wktlex1(lex wktLexer, lval *wktSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = wktTok1[0]
		goto out
	}
	if char < len(wktTok1) {
		token = wktTok1[char]
		goto out
	}
	if char >= wktPrivate {
		if char < wktPrivate+len(wktTok2) {
			token = wktTok2[char-wktPrivate]
			goto out
		}
	}
	for i := 0; i < len(wktTok3); i += 2 {
		token = wktTok3[i+0]
		if token == char {
			token = wktTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = wktTok2[1] /* unknown char */
	}
	if wktDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", wktTokname(token), uint(char))
	}
	return char, token
}

func wktParse(wktlex wktLexer) int {
	return wktNewParser().Parse(wktlex)
}

func (wktrcvr *wktParserImpl) Parse(wktlex wktLexer) int {
	var wktn int
	var wktVAL wktSymType
	var wktDollar []wktSymType
	_ = wktDollar // silence set and not used
	wktS := wktrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	wktstate := 0
	wktrcvr.char = -1
	wkttoken := -1 // wktrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		wktstate = -1
		wktrcvr.char = -1
		wkttoken = -1
	}()
	wktp := -1
	goto wktstack

ret0:
	return 0

ret1:
	return 1

wktstack:
	/* put a state and value onto the stack */
	if wktDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", wktTokname(wkttoken), wktStatname(wktstate))
	}

	wktp++
	if wktp >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktS[wktp] = wktVAL
	wktS[wktp].yys = wktstate

wktnewstate:
	wktn = wktPact[wktstate]
	if wktn <= wktFlag {
		goto wktdefault /* simple state */
	}
	if wktrcvr.char < 0 {
		wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
	}
	wktn += wkttoken
	if wktn < 0 || wktn >= wktLast {
		goto wktdefault
	}
	wktn = wktAct[wktn]
	if wktChk[wktn] == wkttoken { /* valid shift */
		wktrcvr.char = -1
		wkttoken = -1
		wktVAL = wktrcvr.lval
		wktstate = wktn
		if Errflag > 0 {
			Errflag--
		}
		goto wktstack
	}

wktdefault:
	/* default state action */
	wktn = wktDef[wktstate]
	if wktn == -2 {
		if wktrcvr.char < 0 {
			wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if wktExca[xi+0] == -1 && wktExca[xi+1] == wktstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			wktn = wktExca[xi+0]
			if wktn < 0 || wktn == wkttoken {
				break
			}
		}
		wktn = wktExca[xi+1]
		if wktn < 0 {
			goto ret0
		}
	}
	if wktn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			wktlex.Error(wktErrorMessage(wktstate, wkttoken))
			Nerrs++
			if wktDebug >= 1 {
				__yyfmt__.Printf("%s", wktStatname(wktstate))
				__yyfmt__.Printf(" saw %s\n", wktTokname(wkttoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for wktp >= 0 {
				wktn = wktPact[wktS[wktp].yys] + wktErrCode
				if wktn >= 0 && wktn < wktLast {
					wktstate = wktAct[wktn] /* simulate a shift of "error" */
					if wktChk[wktstate] == wktErrCode {
						goto wktstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if wktDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", wktS[wktp].yys)
				}
				wktp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if wktDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", wktTokname(wkttoken))
			}
			if wkttoken == wktEofCode {
				goto ret1
			}
			wktrcvr.char = -1
			wkttoken = -1
			goto wktnewstate /* try again in the same state */
		}
	}

	/* reduction by production wktn */
	if wktDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", wktn, wktStatname(wktstate))
	}

	wktnt := wktn
	wktpt := wktp
	_ = wktpt // guard against "declared and not used"

	wktp -= wktR2[wktn]
	// wktp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if wktp+1 >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktVAL = wktS[wktp+1]

	/* consult goto table to find next state */
	wktn = wktR1[wktn]
	wktg := wktPgo[wktn]
	wktj := wktg + wktS[wktp].yys + 1

	if wktj >= wktLast {
		wktstate = wktAct[wktg]
	} else {
		wktstate = wktAct[wktj]
		if wktChk[wktstate] != -wktn {
			wktstate = wktAct[wktg]
		}
	}
	// dummy call; replaced with literal code
	switch wktnt {

	case 1:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:159
		{
			wktlex.(*wktLex).ret = wktDollar[1].geom
		}
	case 8:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:173
		{
			wktVAL.geom = geom.NewPointFlat(geom.XY, wktDollar[2].coordList)
		}
	case 9:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:177
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 10:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:181
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 11:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:185
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYM, wktDollar[2].coordList)
		}
	case 12:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:189
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 13:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:193
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 14:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:197
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XY)
		}
	case 15:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:201
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYM)
		}
	case 16:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:205
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYZ)
		}
	case 17:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:209
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYZM)
		}
	case 18:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:215
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XY, wktDollar[2].coordList)
		}
	case 19:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:219
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 20:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:223
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 21:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:227
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYM, wktDollar[2].coordList)
		}
	case 22:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:231
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 23:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:235
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 24:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:239
		{
			wktVAL.geom = geom.NewLineString(geom.XY)
		}
	case 25:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:243
		{
			wktVAL.geom = geom.NewLineString(geom.XYM)
		}
	case 26:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:247
		{
			wktVAL.geom = geom.NewLineString(geom.XYZ)
		}
	case 27:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:251
		{
			wktVAL.geom = geom.NewLineString(geom.XYZM)
		}
	case 28:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:257
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XY, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 29:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:261
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZ, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 30:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:265
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZM, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 31:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:269
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYM, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 32:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:273
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZ, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 33:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:277
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZM, wktDollar[2].flatRepr.flatCoords, wktDollar[2].flatRepr.ends)
		}
	case 34:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:281
		{
			wktVAL.geom = geom.NewPolygon(geom.XY)
		}
	case 35:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:285
		{
			wktVAL.geom = geom.NewPolygon(geom.XYM)
		}
	case 36:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:289
		{
			wktVAL.geom = geom.NewPolygon(geom.XYZ)
		}
	case 37:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:293
		{
			wktVAL.geom = geom.NewPolygon(geom.XYZM)
		}
	case 38:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:299
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XY, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 39:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:303
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZ, wktDollar[3].coordList)
		}
	case 40:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:307
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZM, wktDollar[3].coordList)
		}
	case 41:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:311
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYM, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 42:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:315
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 43:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:319
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 44:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:323
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XY)
		}
	case 45:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:327
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYM)
		}
	case 46:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:331
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYZ)
		}
	case 47:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:335
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYZM)
		}
	case 48:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:341
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XY, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 49:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:345
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 50:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:349
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 51:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:353
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 52:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:357
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 53:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:361
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 54:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:365
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XY)
		}
	case 55:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:369
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYM)
		}
	case 56:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:373
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYZ)
		}
	case 57:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:377
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYZM)
		}
	case 58:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:383
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XY, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 59:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:387
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XYZ, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 60:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:391
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XYZM, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 61:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:395
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XYM, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 62:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:399
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XYZ, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 63:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:403
		{
			wktVAL.geom = geom.NewMultiPolygonFlat(geom.XYZM, wktDollar[3].multiPolyFlatRepr.flatCoords, wktDollar[3].multiPolyFlatRepr.endss)
		}
	case 64:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:407
		{
			wktVAL.geom = geom.NewMultiPolygon(geom.XY)
		}
	case 65:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:411
		{
			wktVAL.geom = geom.NewMultiPolygon(geom.XYM)
		}
	case 66:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:415
		{
			wktVAL.geom = geom.NewMultiPolygon(geom.XYZ)
		}
	case 67:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:419
		{
			wktVAL.geom = geom.NewMultiPolygon(geom.XYZM)
		}
	case 68:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:425
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 70:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:432
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 72:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:439
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 74:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:446
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 76:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:453
		{
			wktVAL.multiPolyFlatRepr = appendMultiPolygonFlatCoordsRepr(wktDollar[1].multiPolyFlatRepr, wktDollar[3].multiPolyFlatRepr)
		}
	case 84:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:472
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 85:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:478
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 86:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:484
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 87:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:490
		{
			wktVAL.multiPolyFlatRepr = makeMultiPolygonFlatCoordsRepr(wktDollar[1].flatRepr)
		}
	case 88:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:496
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 89:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:502
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 90:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:508
		{
			wktVAL.flatRepr = wktDollar[2].flatRepr
		}
	case 91:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:514
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(nil)
		}
	case 92:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:520
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 94:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:527
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 96:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:534
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 98:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:541
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 2) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 99:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:550
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 3) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 100:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:559
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 4) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 101:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:570
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 103:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:577
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 105:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:584
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 107:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:591
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 109:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:598
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 117:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:617
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 118:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:623
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 119:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:629
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 120:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:635
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 2) {
				return 1
			}
		}
	case 121:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:643
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 3) {
				return 1
			}
		}
	case 122:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:651
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 4) {
				return 1
			}
		}
	case 123:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:659
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 124:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:665
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 125:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:671
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 126:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:677
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(nil)
		}
	case 127:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:683
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 129:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:690
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 131:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:697
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 133:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:706
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 135:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:713
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 137:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:720
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 139:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:727
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 141:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:734
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 143:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:741
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 145:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:748
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 147:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:755
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 155:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:774
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(nil)
		}
	case 156:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:780
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 157:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:786
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 158:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:792
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 159:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:798
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord}
		}
	case 160:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:804
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord, wktDollar[3].coord}
		}
	case 161:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:810
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord, wktDollar[3].coord, wktDollar[4].coord}
		}
	}
	goto wktstack /* stack new state and value */
}
