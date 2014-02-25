require! <[fs cheerio request async]>
{lists-to-obj, flatten} = require \prelude-ls

columns = <[town village name]>
options =
  * headers:
      'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.73.11 (KHTML, like Gecko) Version/7.0.1 Safari/537.73.11'
      'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8'
      'Referer': 'http://axe-level-4.herokuapp.com/lv4/?page=1'
request = request.defaults options

tasks = for let i in [1 to 24]
  (cb) ->
    d = []
    err, res, body <- request.get "http://axe-level-4.herokuapp.com/lv4/?page=#{i}"
    $ = cheerio.load body
    row = $ 'table.table tr' .first!.next!
    while row.text!
      d.push lists-to-obj columns, (row.children!.map (,e) -> $(e).text!).to-array!
      row .= next!
    cb null, d

err, data <- async.series tasks
data |> flatten |> JSON.stringify |> fs.write-file-sync 'output.json', _
