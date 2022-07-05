package template

import "html/template"

var ItemsList = template.Must(template.New("itemsList").Parse(`
<h1>Items</h1>
<table>
<tr style='text-align: left'>
  <td>Item: price</td>
</tr>
{{ range $key, $value := . }}
<tr style='text-align: left'>
 	  <td><strong>{{ $key }}</strong>: {{ $value }}</td>
</tr>
{{ end }}
</table>
`))

// /list to print its output as an HTML table
