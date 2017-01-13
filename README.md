## Function: TestRandString

#### parameter: 
t (&{%!s(token.Pos=2410) %!s(*ast.SelectorExpr=&{0xc420096be0 0xc420096c00})}) 





## Function: BenchmarkRandString

#### parameter: 
b (&{%!s(token.Pos=2588) %!s(*ast.SelectorExpr=&{0xc420096fc0 0xc420096fe0})}) 





## Function: UrlEncode

#### Comment: 

urlencode

#### parameter: 
url1 (string) 





## Function: UrlDecode

#### parameter: 
url1 (string) 





## Function: FormatUrl

#### Comment: 

FormatUrl is a function for spiders to format urls when the url comes like "/path/to/html" to "http://www.xxx.com/path/to/html"

#### parameter: 
url1 site (string) 

#### return: 
(string) (error) 



## Function: IsCurrentSite

#### Comment: 

IsCurrentSite is a function to test the current url belongs to the current site.

#### parameter: 
url1 (string) site (string) protocol (string) 

#### return: 
(bool) 



## Function: IsDockerHost

#### parameter: 


#### return: 
(bool) 



## Function: IsDockerContainer

#### parameter: 


#### return: 
(bool) 



## Function: EncodingTest

#### Comment: 

EncodingTest is a third party library to get the encoding of certain type `[]byte`, which usually is UTF-8.

#### parameter: 
content (&{%!s(token.Pos=565) %!s(*ast.ArrayType=&{566 <nil> 0xc42000ea00})}) 

#### return: 
encoding(string) err(error) 



## Function: Get

#### Comment: 

Get is a common http client for httpGet operations and hide UA as GoogleBot.

#### parameter: 
urls (string) 

#### return: 
(&{%!s(token.Pos=1017) <nil> byte}) (error) 



## Function: P

#### Comment: 

P is a quick function to print struct.

#### parameter: 
title (string) c (&{%!s(token.Pos=1822) %!s(*ast.FieldList=&{1831 [] 1832}) %!s(bool=false)}) 





## Function: RandString

#### parameter: 
size (int) include (string) 

#### return: 
(string) 



