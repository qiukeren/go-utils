## Function: P

#### Comment: 

P is a quick function to print struct.

#### parameter: 
title (string) c (&{%!s(token.Pos=1464) %!s(*ast.FieldList=&{1473 [] 1474}) %!s(bool=false)}) 





## Function: RandString

#### parameter: 
size (int) include (string) 

#### return: 
(string) 



## Function: TestRandString

#### parameter: 
t (&{%!s(token.Pos=2028) %!s(*ast.SelectorExpr=&{0xc4200985e0 0xc420098600})}) 





## Function: BenchmarkRandString

#### parameter: 
b (&{%!s(token.Pos=2206) %!s(*ast.SelectorExpr=&{0xc4200989c0 0xc4200989e0})}) 





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



## Function: EncodingTest

#### Comment: 

EncodingTest is a third party library to get the encoding of certain type `[]byte`, which usually is UTF-8.

#### parameter: 
content (&{%!s(token.Pos=207) %!s(*ast.ArrayType=&{208 <nil> 0xc42000e3e0})}) 

#### return: 
encoding(string) err(error) 



## Function: Get

#### Comment: 

Get is a common http client for httpGet operations and hide UA as GoogleBot.

#### parameter: 
urls (string) 

#### return: 
(&{%!s(token.Pos=659) <nil> byte}) (error) 



