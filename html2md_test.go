package html2md

import (
	"fmt"
	"testing"
)

var (
	html = []string{`<html><head></head><body>
	<p>hello world</p>
	<h1>testbefore br</h1>
	<br/>
	<h2>testafter br</h2>
	<hr>
	<b>strong</b>
	<strong>sss</strong>
	<img src="a.png" title="test" alt="ss"/>
	<img src="b.png" title="test" alt="ssb" />

	<img src='http://www.lunny.info/wp-content/uploads/2007/04/a710is.jpg' alt='A710IS' />

	<pre>

	stess
	sdfsf

	</pre>

	<blockquote>
	fsafdsaff
	fdsafdsafda
	</blockquote>

	<ul>
	<li>test1</li>
	<li>test2</li>
	</ul>

	<a href="http://github.com/lunny/html2md">go - html2md</a>

	</body></html>
	`,
		` 
	<h2>Title</h2><p><strong>Strong</strong></p>
<!--break-->
<p>Normal – text <em>Em</em>.</p>
<p>Normal<br>Break Line</p>
<ul>
<li>Foo</li>
<li>Bar</li>
</ul>
<ol>
<li>One</li>
<li>Two</li>
</ol>
<code>
var i = 0;
alert(i);
</code>
<blockquote>
<p>It always seems impossible until it is done.</p>

<i>Nelson Mandela</i>
</blockquote>
<p><a href="http://scito.ch">Scito</a></p>
<p><a href="http://scito.ch" title="Visit me">Scito</a></p>
<hr />
<ul>
<li>Foo</li>
<li>Bar
<ol>
<li>One</li>
<li>Two</li>
</ol>
</li>
</ul>
<code>
<strong>Strong in Code</strong>
</code>
<p>Some inline <code><strong></code> in text.</p>

<blockquote>
<p>It always seems impossible until it is done.</p>

<i>Nelson Mandela</i>
<blockquote>
<p>It always seems impossible until it is done.</p>

<i>Nelson Mandela</i>
</blockquote>
</blockquote>
<span><em>Some EM</em></span>
<h2 id="some-title">An id title Test</h2>
<h3>

Header
Containing

Newlines

</h3>
<ul>
	<li>List items </li>
	<li>Ending with </li>
	<li>A space </li>
</ul>
`}
)

func TestHtml2md(t *testing.T) {
	for _, hm := range html {
		mk := Convert(hm)
		fmt.Println(mk)
	}
}
