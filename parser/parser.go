package parser

import "fmt"

const (
    GET = "GET"
    POST = "POST"
    PUT = "PUT"
    PATCH = "PATCH"
    DELETE = "DELETE"
)


type Parser struct {
	message []byte

    ch byte
    ptr int
    next int
}

type RequestInfo struct {
	Method          string
	Route           string
	ProtocolVersion string
}

type RequestHeaderInfo struct {
	Headers []Header 
}

type ParseResult struct {
	info       *RequestInfo
	headerInfo *RequestHeaderInfo
}

type Header struct {
    header string
    value string
}

func New(msg []byte) *Parser {
    p := &Parser{ message: msg, next: 0, ptr: -1 }
    p.advPtr()
	return p
}

func (p * ParseResult) GetObject() * ParseResult {
    return p 
}

func (p * ParseResult) GetMethod() string {
    return p.info.Method
}

func (p *ParseResult) PrintParseResult() {
    fmt.Printf("info\n\tMethod: %s\n\tRoute: %s\n\tProtocolVersion: %s\nheader-info\n", p.info.Method, p.info.Route, p.info.ProtocolVersion)
    
    for _, header := range p.headerInfo.Headers {
        fmt.Printf("\t%s: %s\n", header.header, header.value)
    }
}

func (p *Parser) ParseMessage() *ParseResult {
	result := &ParseResult{}

    //msg := p.message

    result.info = p.parseInfo()
    result.headerInfo = p.parseHeaders()

    return result
}

func (p *Parser) nextToken() string {
    var token string = ""

    p.skipWs()
    for p.ch != 0 {
        if p.message[p.ptr] == ' ' || p.message[p.ptr] == '\n' {
            p.advPtr()
            break
        } else {
            token = token + string(p.message[p.ptr])
            p.advPtr()
        }
    }

    return token 
}

func (p *Parser) nextTokenOnNewLine() string {
    var token string = ""

    p.skipWs()
    for p.ch != 0 {
        if p.message[p.ptr] == '\n' {
            p.advPtr()
            break
        } else {
            token = token + string(p.message[p.ptr])
            p.advPtr()
        }
    }

    return token 
}

func (p * Parser) getNextHeader() Header {
    header := p.nextToken()
    var value string = ""

    if len(header) > 0 {
        if header[len(header) - 1] == ':' {
            header = header[:len(header) - 1]
        }
        value = p.nextTokenOnNewLine()
    }

    return Header{header: header, value: value}
}

func (p * Parser) advPtr() {
    if p.next >= len(p.message) {
        p.ch = 0
    } else {
        p.ch = p.message[p.next]
    }

    p.ptr = p.next
    p.next += 1
}

func (p * Parser) skipWs() {
	for p.message[p.ptr] == ' ' || p.message[p.ptr] == '\t' || p.message[p.ptr] == '\n' || p.message[p.ptr] == '\r' {
        p.advPtr()
	}
}

func (p * Parser) parseInfo() *RequestInfo {
    res := &RequestInfo{}
    
    method := p.nextToken()
    switch method {
    case "GET":
        res.Method = GET
    case "PATCH":
        res.Method = PATCH
    case "PUT":
        res.Method = PUT
    case "DELETE":
        res.Method = DELETE
    case "POST":
        res.Method = POST
    }

    res.Route = p.nextToken()[1:]
    res.ProtocolVersion = p.nextToken()

    return res
}

func (p * Parser) parseHeaders() *RequestHeaderInfo {
    var headers []Header

    for p.ch != 0 {
        header := p.getNextHeader()
        if header.header == "" && header.value == "" {
            break
        }
        headers = append(headers, header)
    }

    res := &RequestHeaderInfo{ Headers: headers } 
    return res
}

